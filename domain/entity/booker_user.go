package entity

type BookerUser struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	IconURL         string        `json:"icon_url"`
	Identifier      string        `json:"identifier"`
	HeaderImageURL  string        `json:"header_image_url"`
	MetaDescription string        `json:"meta_description"`
	WhyDescription  string        `json:"why_description"`
	HowDescription  string        `json:"how_description"`
	FreeDescription string        `json:"free_description"`
	Tags            []*UserTag    `json:"tags"`
	Locations       []*Location   `json:"locations"`
	WebsiteURLs     []*WebsiteURL `json:"website_urls"`
	UserImages      []*UserImage  `json:"images"`
}

var BOOKERIDF = "booker"
