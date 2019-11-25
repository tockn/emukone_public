package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
)

type invite struct {
	db     *gorm.DB
	hashID repository.HashID
}

func NewInvite(db *gorm.DB, h repository.HashID) repository.Invite {
	return &invite{
		db,
		h,
	}
}

func (repo *invite) FindByID(idStr string) (*entity.Invite, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}

	var mi mysql.Invite
	if err := preloadedInvite(repo.db).First(&mi, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return mysqlInviteToDomainAndIDs(&mi, repo.hashID), nil
}

func (repo *invite) FindByArtistID(idStr string) ([]*entity.Invite, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}

	var mis []*mysql.Invite
	if err := preloadedInvite(repo.db).Find(&mis, "artist_user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return mysqlInvitesToDomainsAndIDs(mis, repo.hashID), nil
}

func (repo *invite) FindByEventID(idStr string) ([]*entity.Invite, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}

	var mis []*mysql.Invite
	if err := preloadedInvite(repo.db).Find(&mis, "event_id = ?", id).Error; err != nil {
		return nil, err
	}
	return mysqlInvitesToDomainsAndIDs(mis, repo.hashID), nil
}

func (repo *invite) FindByIDAndArtistID(idStr, artistIDStr string) (*entity.Invite, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}
	artistID, err := repo.hashID.Decode(artistIDStr)
	if err != nil {
		return nil, err
	}

	var mi mysql.Invite
	if err := preloadedInvite(repo.db).First(&mi, "id = ? AND artist_user_id = ?", id, artistID).Error; err != nil {
		return nil, err
	}
	return mysqlInviteToDomainAndIDs(&mi, repo.hashID), nil
}

func (repo *invite) FindByIDAndInviterID(idStr, inviterIDStr string) (*entity.Invite, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}
	inviterID, err := repo.hashID.Decode(inviterIDStr)
	if err != nil {
		return nil, err
	}

	var mi mysql.Invite
	if err := preloadedInvite(repo.db).First(&mi, "id = ? AND inviter_user_id = ?", id, inviterID).Error; err != nil {
		return nil, err
	}
	return mysqlInviteToDomainAndIDs(&mi, repo.hashID), nil
}

func (repo *invite) FindByEventIDAndArtistID(eventIDStr, artistIDStr string) (*entity.Invite, error) {
	eventID, err := repo.hashID.Decode(eventIDStr)
	if err != nil {
		return nil, err
	}
	artistID, err := repo.hashID.Decode(artistIDStr)
	if err != nil {
		return nil, err
	}

	var mi mysql.Invite
	if err := preloadedInvite(repo.db).First(&mi, "event_id = ? AND artist_user_id = ?", eventID, artistID).Error; err != nil {
		return nil, err
	}
	return mysqlInviteToDomainAndIDs(&mi, repo.hashID), nil
}

func (repo *invite) Store(ei *entity.Invite) (*entity.Invite, error) {

	mi, err := newMySQLInviteFromDomainAndIDs(ei, repo.hashID)
	if err != nil {
		return nil, err
	}
	if err := repo.db.Create(&mi).Error; err != nil {
		return nil, err
	}
	ei.ID = repo.hashID.Encode(mi.ID)
	return ei, nil
}

func (repo invite) Update(di *entity.Invite) (*entity.Invite, error) {
	bi, err := newMySQLInviteFromDomainAndIDs(di, repo.hashID)
	if err != nil {
		return nil, err
	}
	cols := []string{
		"description",
	}
	if err := repo.db.Model(&bi).Select(cols).Update(bi).Error; err != nil {
		return nil, err
	}
	return di, nil
}

func (repo invite) UpdateStatus(di *entity.Invite) (*entity.Invite, error) {
	id, err := repo.hashID.Decode(di.ID)
	if err != nil {
		return nil, err
	}
	artistUserID, err := repo.hashID.Decode(di.ArtistID)
	if err != nil {
		return nil, err
	}
	bi := mysql.NewInviteFromDomainAndIDs(di, id, artistUserID, 0, 0)
	cols := []string{
		"status",
	}
	if err := repo.db.Model(&bi).Select(cols).Update(bi).Error; err != nil {
		return nil, err
	}
	return di, nil
}

// TODO Canceledカラム作ってるけど、Deleteは論理削除でなくdeleteする。ここら辺どうするか議論したい
func (repo invite) Delete(de *entity.Invite) error {
	bi, err := newMySQLInviteFromDomainAndIDs(de, repo.hashID)
	if err != nil {
		return err
	}
	if err := repo.db.Delete(&bi).Error; err != nil {
		return err
	}
	return nil
}

func newMySQLInviteFromDomainAndIDs(ei *entity.Invite, h repository.HashID) (*mysql.Invite, error) {
	var id int64
	if ei.ID == "" {
		id = 0
	} else {
		idd, err := h.Decode(ei.ID)
		if err != nil {
			return nil, err
		}
		id = idd
	}

	artistUserID, err := h.Decode(ei.ArtistID)
	if err != nil {
		return nil, err
	}

	eventID, err := h.Decode(ei.EventID)
	if err != nil {
		return nil, err
	}

	inviterID, err := h.Decode(ei.InviterID)
	if err != nil {
		return nil, err
	}
	return mysql.NewInviteFromDomainAndIDs(ei, id, artistUserID, eventID, inviterID), nil
}

func mysqlInviteToDomainAndIDs(i *mysql.Invite, h repository.HashID) *entity.Invite {
	id := h.Encode(i.ID)
	artistUserID := h.Encode(i.ArtistUserID)
	eventID := h.Encode(i.EventID)
	inviterID := h.Encode(i.InviterUserID)
	return i.ToDomain(id, artistUserID, eventID, inviterID)
}

func mysqlInvitesToDomainsAndIDs(is []*mysql.Invite, h repository.HashID) []*entity.Invite {
	dis := make([]*entity.Invite, len(is))
	for i, v := range is {
		dis[i] = mysqlInviteToDomainAndIDs(v, h)
	}
	return dis
}

func preloadedInvite(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Event").
		Preload("Inviter").
		Preload("Artist")
}
