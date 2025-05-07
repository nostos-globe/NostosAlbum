package models

type Media struct {
	MediaID   int64   `json:"mediaId"`
	URL       string  `json:"url"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
