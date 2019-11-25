package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
)

type user struct {
	db     *gorm.DB
	hashID repository.HashID
}

func NewUser(db *gorm.DB, h repository.HashID) repository.User {
	return &user{
		db,
		h,
	}
}

func (repo *user) FindMetaByID(idStr string) (*entity.UserMeta, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}

	var u mysql.User
	if err := repo.db.Preload("UserTags").Preload("Locations").Find(&u, id).Error; err != nil {
		return nil, err
	}
	entityID := repo.hashID.Encode(id)
	du := u.ToDomainMeta(entityID)
	return du, nil
}

func (repo *user) FindIDByExternalService(s *entity.ExternalService) (string, bool, error) {
	var es mysql.ExternalServiceUID
	res := repo.db.Where("service = ? AND uid = ?", s.Service, s.UID).
		Find(&es)
	if res.RecordNotFound() {
		return "", false, nil
	}
	if res.Error != nil {
		return "", false, res.Error
	}
	return repo.hashID.Encode(es.UserID), true, nil
}

func (repo *user) StoreExternalService(de *entity.ExternalService, userIDStr string) (err error) {
	userID, err := repo.hashID.Decode(userIDStr)
	if err != nil {
		return err
	}
	es := mysql.NewExternalServiceUIDFromDomainAndUserID(de, userID)
	err = repo.db.Create(es).Error
	return
}

func (repo *user) FindMetaByEventID(idStr string) ([]*entity.UserMeta, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}
	var us []*mysql.User
	err = repo.db.
		Table("users").
		Joins("inner join events_related_users as eru on eru.user_id = users.id").
		Joins("inner join events on eru.event_id = events.id").
		Where("events.id = ?", id).
		Scan(&us).Error
	if err != nil {
		return nil, err
	}

	for i, u := range us {
		err = repo.db.Model(u).Related(&u.UserTags, "UserTags").Error
		if err != nil {
			return nil, err
		}
		err = repo.db.Model(u).Related(&u.Locations, "Locations").Error
		if err != nil {
			return nil, err
		}
		us[i] = u
	}
	uIDs := make([]string, len(us))
	for i, u := range us {
		uIDs[i] = repo.hashID.Encode(u.ID)
	}
	return mysql.UsersToDomainMetas(us, uIDs), nil
}

func (repo *user) FindTag() (tags []*entity.UserTag, err error) {
	err = repo.db.Find(&tags).Error
	return
}
