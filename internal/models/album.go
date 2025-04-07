package models

type Album struct {
	AlbumID      int     	`json:"album_id" db:"album_id"`
	UserID       uint     	`json:"user_id" db:"user_id"`
	Name         string    	`json:"name" db:"name"`
	Description  string    	`json:"description,omitempty" db:"description"`
	CreationDate string 	`json:"start_date,omitempty" db:"creation_date"`
}
