package entity

type User struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	IconURL         string        `json:"icon_url"`
	Identifier      string        `json:"identifier"`
	HeaderImageURL  string        `json:"header_image_url"`
	MetaDescription string        `json:"meta_description"`
	Tags            []*UserTag    `json:"tags"`
	Locations       []*Location   `json:"locations"`
	WebsiteURLs     []*WebsiteURL `json:"website_urls"`
	UserImages      []*UserImage  `json:"user_images"`
}

type UserMeta struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Identifier      string      `json:"identifier"`
	IconURL         string      `json:"icon_url"`
	HeaderImageURL  string      `json:"header_image_url"`
	MetaDescription string      `json:"meta_description"`
	UserTags        []*UserTag  `json:"user_tags"`
	Locations       []*Location `json:"locations"`
}

type UserTag struct {
	Name string `json:"name"`
}

type Location struct {
	Name string `json:"name"`
}

type WebsiteURL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type UserImage struct {
	ImageURL string `json:"url"`
}
