package service

import (
	"main/internal/db"
	"main/internal/models"
)

type AlbumService struct {
    AlbumRepo *db.AlbumRepository
}

func (s *AlbumService) CreateAlbum(album models.Album) (any, error) {
    result, err := s.AlbumRepo.CreateAlbum(album)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (s *AlbumService) GetAlbumsByUserID(userID uint) ([]models.Album, error) {
    albums, err := s.AlbumRepo.GetAlbumsByUserID(userID)
    if err != nil {
        return nil, err
    }

    return albums, nil
}

func (s *AlbumService) GetPublicAlbums() ([]models.Album, error) {
    albums, err := s.AlbumRepo.GetPublicAlbums()
    if err!= nil {
        return nil, err
    }

    return albums, nil 
}

func (s *AlbumService) GetAlbumByID(albumID string) (models.Album, error) {
    album, err := s.AlbumRepo.GetAlbumByID(albumID)
    if err!= nil {
        return models.Album{}, err
    } 

    return album, nil
}

func (s *AlbumService) GetTripsByAlbumID(albumID string) ([]uint, error) {
    trips, err := s.AlbumRepo.GetTripsByAlbumID(albumID)
    if err!= nil {
        return nil, err
    } 

    return trips, nil
}

func (s *AlbumService) UpdateAlbum(album models.Album) (models.Album, error) {
    updatedAlbum, err := s.AlbumRepo.UpdateAlbum(album)
    if err!= nil {
        return models.Album{}, err
    }

    return updatedAlbum, nil
}

func (s *AlbumService) DeleteAlbum(albumID string) error {
    err := s.AlbumRepo.DeleteAlbum(albumID)
    if err!= nil {
        return err
    } 

    return nil
}