package models

type Album struct {
	AlbumID      int         `gorm:"primaryKey;autoIncrement"`
	UserID       uint     	`json:"user_id" db:"user_id"`
	Name         string    	`json:"name" db:"name"`
	Description  string    	`json:"description,omitempty" db:"description"`
	Visibility   string 	`json:"visibility" db:"visibility" default:"PRIVATE"`
	CreationDate string 	`json:"creation_date,omitempty" db:"creation_date"`
}
