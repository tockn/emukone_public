package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
)

type bookerUser struct {
	db     *gorm.DB
	hashID repository.HashID
}

func NewBookerUser(db *gorm.DB, h repository.HashID) repository.BookerUser {
	return &bookerUser{
		db,
		h,
	}
}

func (repo *bookerUser) FindByID(idStr string) (*entity.BookerUser, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}
	var u mysql.User
	err = repo.db.
		Preload("BookerUser").
		Preload("UserTags").
		Preload("Locations").
		Where("id = ? AND identifier = ?", id, mysql.BOOKERIDF).
		First(&u).
		Error
	if err != nil {
		return nil, err
	}
	entityID := repo.hashID.Encode(u.ID)
	return u.ToDomainBooker(entityID)
}

func (repo *bookerUser) Store(du *entity.BookerUser) (*entity.BookerUser, error) {
	var id int64
	if du.ID != "" {
		idd, err := repo.hashID.Decode(du.ID)
		if err != nil {
			return nil, err
		}
		id = idd
	}

	bu := mysql.NewUserFromDomainBookerAndID(du, id)
	if err := bu.FindOrCreate(repo.db); err != nil {
		return nil, err
	}

	newEntityID := repo.hashID.Encode(bu.ID)
	nbu, err := bu.ToDomainBooker(newEntityID)
	if err != nil {
		return nil, err
	}
	return nbu, nil
}

func (repo *bookerUser) Update(du *entity.BookerUser) (*entity.BookerUser, error) {
	id, err := repo.hashID.Decode(du.ID)
	if err != nil {
		return nil, err
	}
	bu := mysql.NewUserFromDomainBookerAndID(du, id)
	if err != nil {
		return nil, err
	}
	cols := []string{
		"name",
		"icon_url",
		"header_image_url",
		"meta_description",
		"why_description",
		"how_description",
		"free_description",
	}
	if err := bu.UpdateBookerUser(repo.db, cols); err != nil {
		return nil, err
	}
	return du, nil
}

func (repo *bookerUser) Delete(du *entity.BookerUser) error {
	id, err := repo.hashID.Decode(du.ID)
	if err != nil {
		return err
	}
	bu := mysql.NewUserFromDomainBookerAndID(du, id)
	if err != nil {
		return err
	}
	if err := repo.db.Model(bu).Update("deleted", 1).Error; err != nil {
		return err
	}
	return nil
}
