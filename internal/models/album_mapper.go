package models

import "time"

type AlbumMapper struct{}

func (m *AlbumMapper) ToAlbum(req struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
}, tokenResponse interface{}) Album {
    return Album{
        Name:        req.Name,
        UserID:      tokenResponse.(uint),
        Description: req.Description,
		Visibility:  req.Visibility,
		CreationDate: time.Now().Format("2006-01-02"),
    }
}