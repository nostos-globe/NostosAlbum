package db

import (
	"main/internal/models"

	"gorm.io/gorm"
)

type AlbumRepository struct {
	DB *gorm.DB
}

func (repo *AlbumRepository) CreateAlbum(album models.Album) (any, error) {
	result := repo.DB.Table("albums.albums").Create(&album)

	if result.Error != nil {
		return nil, result.Error
	}

	return album, nil
}

func (repo *AlbumRepository) GetAlbumsByUserID(userID uint) ([]models.Album, error) {
	var albums []models.Album
	result := repo.DB.Table("albums.albums").Where("user_id = ?", userID).Find(&albums)	

	if result.Error!= nil {
		return nil, result.Error
	}

	return albums, nil	
}

func (repo *AlbumRepository) GetPublicAlbums() ([]models.Album, error) {
	var albums []models.Album
	result := repo.DB.Table("albums.albums").Where("visibility = ?", "PUBLIC").Find(&albums)	

	if result.Error!= nil {
		return nil, result.Error
	}

	return albums, nil	
}

func (repo *AlbumRepository) GetTripsByAlbumID(albumID string) ([]uint, error) {
	var trips []uint
	result := repo.DB.Table("albums.album_trips").Select("trip_id").Where("album_id =?", albumID).Find(&trips)
	
	if result.Error!= nil {
		return nil, result.Error	
	}

	return trips, nil
}

func (repo *AlbumRepository) GetAlbumByID(albumID string) (models.Album, error) {
	var album models.Album
	result := repo.DB.Table("albums.albums").Where("album_id =?", albumID).Find(&album)
	
	if result.Error!= nil {
		return models.Album{}, result.Error	
	}

	return album, nil
}

func (repo *AlbumRepository) UpdateAlbum(album models.Album) (models.Album, error) {
	result := repo.DB.Table("albums.albums").Where("album_id = ?", album.AlbumID).Updates(&album)

	if result.Error!= nil {
		return models.Album{}, result.Error
	}	
	return album, nil
}

func (repo *AlbumRepository) DeleteAlbum(albumID string) error {
	result := repo.DB.Table("albums.albums").Where("album_id =?", albumID).Delete(&models.Album{})

	if result.Error!= nil {
		return result.Error	
	}	

	return nil
}