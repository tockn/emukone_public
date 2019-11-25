package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type User struct {
	// For All User
	ID              int64 `gorm:"primary_key"`
	Name            string
	Identifier      string
	IconURL         string
	HeaderImageURL  string
	MetaDescription string
	Deleted         int64
	UserTags        []*UserTag  `gorm:"many2many:user_tags_related_users"`
	Locations       []*Location `gorm:"many2many:locations_related_users"`
	WebsiteURLs     []*WebsiteURL
	UserImages      []*UserImage
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// For BookerUser
	BookerUser *BookerUser

	// For ArtistUser
	ArtistUser    *ArtistUser
	Ichioshis     []*Ichioshi     `gorm:"foreignkey:ArtistUserID"`
	ArtistMembers []*ArtistMember `gorm:"foreignkey:ArtistUserID"`

	ExternalServiceUIDs []*ExternalServiceUID
}

func NewUserFromDomainArtistAndID(du *entity.ArtistUser, id int64) (u *User) {
	bu := &ArtistUser{
		UserID:          id,
		Identifier:      ARTISTIDF,
		WhyDescription:  du.WhyDescription,
		HowDescription:  du.HowDescription,
		FreeDescription: du.FreeDescription,
	}
	u = &User{
		ID:              id,
		Name:            du.Name,
		IconURL:         du.IconURL,
		Identifier:      ARTISTIDF,
		HeaderImageURL:  du.HeaderImageURL,
		MetaDescription: du.MetaDescription,
		UserTags:        NewUserTagsFromDomains(du.Tags),
		Locations:       NewLocationsFromDomains(du.Locations),
		WebsiteURLs:     NewWebsiteURLsFromDomains(du.WebsiteURLs, id),
		UserImages:      NewUserImagesFromDomains(du.UserImages, id),

		ArtistUser:    bu,
		Ichioshis:     NewIchioshisFromDomainAndID(du.Ichioshis, id),
		ArtistMembers: NewArtistMembersFromDomainsAndID(du.ArtistMembers, id),
	}
	return
}

func NewUserFromDomain(du *entity.User, id int64) *User {
	return &User{
		ID:              id,
		Name:            du.Name,
		IconURL:         du.IconURL,
		Identifier:      du.Identifier,
		HeaderImageURL:  du.HeaderImageURL,
		MetaDescription: du.MetaDescription,
		UserTags:        NewUserTagsFromDomains(du.Tags),
		Locations:       NewLocationsFromDomains(du.Locations),
		WebsiteURLs:     NewWebsiteURLsFromDomains(du.WebsiteURLs, id),
		UserImages:      NewUserImagesFromDomains(du.UserImages, id),
	}
}

func NewUserFromDomainBookerAndID(du *entity.BookerUser, id int64) (u *User) {
	bu := &BookerUser{
		UserID:          id,
		Identifier:      BOOKERIDF,
		WhyDescription:  du.WhyDescription,
		HowDescription:  du.HowDescription,
		FreeDescription: du.FreeDescription,
	}
	u = &User{
		ID:              id,
		Name:            du.Name,
		IconURL:         du.IconURL,
		Identifier:      BOOKERIDF,
		HeaderImageURL:  du.HeaderImageURL,
		MetaDescription: du.MetaDescription,
		UserTags:        NewUserTagsFromDomains(du.Tags),
		Locations:       NewLocationsFromDomains(du.Locations),
		WebsiteURLs:     NewWebsiteURLsFromDomains(du.WebsiteURLs, id),
		UserImages:      NewUserImagesFromDomains(du.UserImages, id),

		BookerUser: bu,
	}
	return
}

func (u *User) ToDomainMeta(entityID string) *entity.UserMeta {
	if u.Deleted == 1 {
		return &entity.UserMeta{}
	}
	return &entity.UserMeta{
		ID:              entityID,
		Name:            u.Name,
		Identifier:      u.Identifier,
		IconURL:         u.IconURL,
		HeaderImageURL:  u.HeaderImageURL,
		MetaDescription: u.MetaDescription,
		UserTags:        UserTagsToDomains(u.UserTags),
		Locations:       LocationsToDomains(u.Locations),
	}
}

func UsersToDomainMetas(us []*User, entityIDs []string) []*entity.UserMeta {
	eus := make([]*entity.UserMeta, len(us))
	for i, u := range us {
		eus[i] = u.ToDomainMeta(entityIDs[i])
	}
	return eus
}

func (u *User) ToDomain(entityID string) *entity.User {
	if u.Deleted == 1 {
		return &entity.User{}
	}
	return &entity.User{
		ID:              entityID,
		Name:            u.Name,
		Identifier:      u.Identifier,
		IconURL:         u.IconURL,
		HeaderImageURL:  u.HeaderImageURL,
		MetaDescription: u.MetaDescription,
	}
}

func (u *User) ToDomainBooker(entityID string) (*entity.BookerUser, error) {
	if u.Deleted == 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.BookerUser{
		ID:              entityID,
		Name:            u.Name,
		IconURL:         u.IconURL,
		Identifier:      BOOKERIDF,
		HeaderImageURL:  u.HeaderImageURL,
		MetaDescription: u.MetaDescription,
		WhyDescription:  u.BookerUser.WhyDescription,
		HowDescription:  u.BookerUser.HowDescription,
		FreeDescription: u.BookerUser.FreeDescription,
		Tags:            UserTagsToDomains(u.UserTags),
		Locations:       LocationsToDomains(u.Locations),
		WebsiteURLs:     WebsiteURLsToDomains(u.WebsiteURLs),
		UserImages:      UserImagesToDomains(u.UserImages),
	}, nil
}

func (u *User) ToDomainArtist(entityID string) (*entity.ArtistUser, error) {
	if u.Deleted == 1 || u.ArtistUser == nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.ArtistUser{
		ID:              entityID,
		Name:            u.Name,
		IconURL:         u.IconURL,
		Identifier:      ARTISTIDF,
		HeaderImageURL:  u.HeaderImageURL,
		MetaDescription: u.MetaDescription,
		WhyDescription:  u.ArtistUser.WhyDescription,
		HowDescription:  u.ArtistUser.HowDescription,
		FreeDescription: u.ArtistUser.FreeDescription,
		Tags:            UserTagsToDomains(u.UserTags),
		Locations:       LocationsToDomains(u.Locations),
		WebsiteURLs:     WebsiteURLsToDomains(u.WebsiteURLs),
		UserImages:      UserImagesToDomains(u.UserImages),
		Ichioshis:       IchioshisToDomains(u.Ichioshis),
		ArtistMembers:   ArtistMembersToDomains(u.ArtistMembers),
	}, nil
}

func (u *User) FindOrCreate(db *gorm.DB) error {
	return db.FirstOrCreate(u, "id = ?", u.ID).Error
}

func (u *User) UpdateArtistUser(db *gorm.DB, cols []string) error {
	tx := db.Begin()
	if err := u.MakeRelationOfUserTags(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.MakeRelationOfLocations(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.UpdateArtistMembers(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.UpdateIchioshis(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.UpdateUserImages(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.UpdateWebsiteURLs(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(u.ArtistUser).Select(cols).Save(u.ArtistUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(u).Select(cols).Save(u).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *User) UpdateBookerUser(db *gorm.DB, cols []string) error {
	tx := db.Begin()
	if err := u.MakeRelationOfUserTags(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.MakeRelationOfLocations(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.UpdateUserImages(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := u.UpdateWebsiteURLs(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(u).Select(cols).Save(u).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(u.BookerUser).Select(cols).Save(u.BookerUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *User) MakeRelationOfUserTags(db *gorm.DB) error {
	urus := make([]*UserTagsRelatedUsers, len(u.UserTags))
	for i := range u.UserTags {
		err := db.First(u.UserTags[i], "name = ?", u.UserTags[i].Name).Error
		if err == gorm.ErrRecordNotFound {
			if err1 := db.Create(u.UserTags[i]).Error; err1 != nil {
				return err1
			}
		} else if err != nil {
			return err
		}
		urus[i] = &UserTagsRelatedUsers{
			UserID:    u.ID,
			UserTagID: u.UserTags[i].ID,
		}
	}
	if err := db.Delete(urus, "user_id = ?", u.ID).Error; err != nil {
		return err
	}
	for _, uru := range urus {
		if err := db.Create(uru).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *User) MakeRelationOfLocations(db *gorm.DB) error {
	lrus := make([]*LocationsRelatedUsers, len(u.Locations))
	for i := range u.Locations {
		err := db.First(u.Locations[i], "name = ?", u.Locations[i].Name).Error
		if err == gorm.ErrRecordNotFound {
			if err1 := db.Create(u.Locations[i]).Error; err1 != nil {
				return err1
			}
		} else if err != nil {
			return err
		}
		lrus[i] = &LocationsRelatedUsers{
			UserID:     u.ID,
			LocationID: u.Locations[i].ID,
		}
	}
	if err := db.Delete(lrus, "user_id = ?", u.ID).Error; err != nil {
		return err
	}
	for _, lru := range lrus {
		if err := db.Create(lru).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *User) UpdateWebsiteURLs(db *gorm.DB) error {
	if err := db.Delete(u.WebsiteURLs, "user_id = ?", u.ID).Error; err != nil {
		return nil
	}
	for _, url := range u.WebsiteURLs {
		if err := db.Create(url).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *User) UpdateUserImages(db *gorm.DB) error {
	if err := db.Delete(u.UserImages, "user_id = ?", u.ID).Error; err != nil {
		return nil
	}
	for _, img := range u.UserImages {
		if err := db.Create(img).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *User) UpdateIchioshis(db *gorm.DB) error {
	if err := db.Delete(u.Ichioshis, "artist_user_id = ?", u.ID).Error; err != nil {
		return nil
	}
	for _, ich := range u.Ichioshis {
		if err := db.Create(ich).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *User) UpdateArtistMembers(db *gorm.DB) error {
	if err := db.Delete(u.ArtistMembers, "artist_user_id = ?", u.ID).Error; err != nil {
		return nil
	}
	for _, mem := range u.ArtistMembers {
		if err := db.Create(mem).Error; err != nil {
			return err
		}
	}
	return nil
}
