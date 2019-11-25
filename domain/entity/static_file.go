package entity

type Image struct {
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
	Base64      string `json:"base64"`
}
