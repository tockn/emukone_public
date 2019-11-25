package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
)

type artistUser struct {
	db     *gorm.DB
	hashID repository.HashID
}

func NewArtistUser(db *gorm.DB, h repository.HashID) repository.ArtistUser {
	return &artistUser{
		db,
		h,
	}
}

func (repo *artistUser) FindByID(idStr string) (*entity.ArtistUser, error) {
	id, err := repo.hashID.Decode(idStr)
	if err != nil {
		return nil, err
	}
	var u mysql.User
	err = repo.db.
		Preload("ArtistUser").
		Preload("UserTags").
		Preload("Locations").
		Preload("WebsiteURLs").
		Preload("UserImages").
		Preload("Ichioshis").
		Preload("ArtistMembers").
		First(&u, "id = ? AND identifier = ?", id, mysql.ARTISTIDF).Error
	if err != nil {
		return nil, err
	}
	entityID := repo.hashID.Encode(u.ID)
	return u.ToDomainArtist(entityID)
}

func (repo *artistUser) Store(du *entity.ArtistUser) (*entity.ArtistUser, error) {
	var id int64
	if du.ID != "" {
		idd, err := repo.hashID.Decode(du.ID)
		if err != nil {
			return nil, err
		}
		id = idd
	}

	bu := mysql.NewUserFromDomainArtistAndID(du, id)
	if err := bu.FindOrCreate(repo.db); err != nil {
		return nil, err
	}

	newEntityID := repo.hashID.Encode(bu.ID)
	nbu, err := bu.ToDomainArtist(newEntityID)
	if err != nil {
		return nil, err
	}
	return nbu, nil
}

func (repo *artistUser) Update(du *entity.ArtistUser) (*entity.ArtistUser, error) {
	id, err := repo.hashID.Decode(du.ID)
	if err != nil {
		return nil, err
	}
	bu := mysql.NewUserFromDomainArtistAndID(du, id)
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
	if err := bu.UpdateArtistUser(repo.db, cols); err != nil {
		return nil, err
	}
	return du, nil
}

func (repo *artistUser) Delete(du *entity.ArtistUser) error {
	id, err := repo.hashID.Decode(du.ID)
	if err != nil {
		return err
	}
	bu := mysql.NewUserFromDomainArtistAndID(du, id)
	if err != nil {
		return err
	}
	if err := repo.db.Model(bu).Update("deleted", 1).Error; err != nil {
		return err
	}
	return nil
}
