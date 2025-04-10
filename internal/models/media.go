package models

type Media struct {
	MediaID   int64  `json:"mediaId"`
	URL       string `json:"url"`
	Latitude  uint   `json:"latitude"`
	Longitude uint   `json:"longitude"`
}
