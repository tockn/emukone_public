package entity

type ArtistUser struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	IconURL         string          `json:"icon_url"`
	Identifier      string          `json:"identifier"`
	HeaderImageURL  string          `json:"header_image_url"`
	MetaDescription string          `json:"meta_description"`
	WhyDescription  string          `json:"why_description"`
	HowDescription  string          `json:"how_description"`
	FreeDescription string          `json:"free_description"`
	Tags            []*UserTag      `json:"tags"`
	Locations       []*Location     `json:"locations"`
	WebsiteURLs     []*WebsiteURL   `json:"website_urls"`
	UserImages      []*UserImage    `json:"images"`
	Ichioshis       []*Ichioshi     `json:"ichioshis"`
	ArtistMembers   []*ArtistMember `json:"members"`
}

type Ichioshi struct {
	Title   string `json:"title"`
	Service string `json:"service"`
	URL     string `json:"url"`
	Embed   string `json:"embed"`
}

type ArtistMember struct {
	Name string `json:"name"`
	Part string `json:"part"`
}

var ARTISTIDF = "artist"
