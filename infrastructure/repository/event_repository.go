package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
)

type event struct {
	db     *gorm.DB
	hashID repository.HashID
}

func NewEvent(db *gorm.DB, h repository.HashID) repository.Event {
	return &event{
		db,
		h,
	}
}

func (repo *event) FindByID(idStr string) (*entity.Event, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}

	var e mysql.Event
	if err := repo.db.Find(&e, id).Error; err != nil {
		return nil, err
	}
	entityID := repo.hashID.Encode(e.ID)
	return e.ToDomain(entityID), nil
}

func (repo *event) FindByUserID(idStr string) ([]*entity.Event, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}

	var es []*mysql.Event
	err = repo.db.Table("events").
		Joins("inner join events_related_users as eru on events.id = eru.event_id").
		Joins("inner join users as u on u.id = eru.user_id").
		Where("u.id = ?", id).
		Scan(&es).Error
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(es))
	for i, e := range es {
		id := repo.hashID.Encode(e.ID)
		ids[i] = id
	}
	ees, err := mysql.EventToDomains(es, ids)
	if err != nil {
		return nil, err
	}
	return ees, nil
}

func (repo *event) FindByIDAndUserID(idStr, userIDStr string) (*entity.Event, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}
	userID, err := repo.hashID.Decode(userIDStr)
	if err != nil {
		return nil, err
	}

	var e mysql.Event
	err = repo.db.Table("events").
		Joins("inner join events_related_users as eru on events.id = eru.event_id").
		Joins("inner join users as u on u.id = eru.user_id").
		Where("events.id = ? AND u.id = ?", id, userID).
		Scan(&e).Error
	if err != nil {
		return nil, err
	}
	return e.ToDomain(repo.hashID.Encode(e.ID)), nil
}

func (repo *event) Store(du *entity.Event, authorIDStr string) (*entity.Event, error) {
	authorID, err := repo.hashID.Decode(authorIDStr)
	if err != nil {
		return nil, err
	}

	var id int64
	if du.ID != "" {
		idd, err := repo.hashID.Decode(du.ID)
		if err != nil {
			return nil, err
		}
		id = idd
	}

	tx := repo.db.Begin()

	bu := mysql.NewEventFromDomainAndID(du, id)
	if err := tx.FirstOrCreate(&bu, "id = ?", bu.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//TODO 関係者を増やすAPIとか作る
	eru := mysql.NewEventsRelatedUsers(bu.ID, authorID)
	if err := tx.Create(&eru).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	newEntityID := repo.hashID.Encode(bu.ID)
	return bu.ToDomain(newEntityID), nil
}

func (repo *event) Update(de *entity.Event) (*entity.Event, error) {
	id, err := repo.hashID.Decode(de.ID)
	if err != nil {
		return nil, err
	}
	be := mysql.NewEventFromDomainAndID(de, id)
	if err != nil {
		return nil, err
	}
	cols := []string{
		"name",
		"meta_description",
		"why_description",
		"free_description",
		"open_time",
		"start_time",
		"end_time",
		"canceled",
		"place",
		"ticket_price",
		"header_image_url",
	}
	if err := repo.db.Model(&be).Select(cols).Update(be).Error; err != nil {
		return nil, err
	}
	return de, nil
}

// TODO Canceledカラム作ってるけど、Deleteは論理削除でなくdeleteする。ここら辺どうするか議論したい
func (repo *event) Delete(de *entity.Event) error {
	id, err := repo.hashID.Decode(de.ID)
	if err != nil {
		return err
	}
	bu := mysql.NewEventFromDomainAndID(de, id)
	if err != nil {
		return err
	}
	if err := repo.db.Delete(&bu).Error; err != nil {
		return err
	}
	return nil
}
