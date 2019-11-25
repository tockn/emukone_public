package entity

import "time"

type Event struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	MetaDescription string        `json:"meta_description"`
	WhyDescription  string        `json:"why_description"`
	FreeDescription string        `json:"free_description"`
	Place           string        `json:"place"`
	TicketPrice     string        `json:"ticket"`
	HeaderImageURL  string        `json:"header_image_url"`
	OpenTime        time.Time     `json:"open_time"`
	StartTime       time.Time     `json:"start_time"`
	EventImages     []*EventImage `json:"images"`
}

type EventImage struct {
	ImageURL string `json:"url"`
}
