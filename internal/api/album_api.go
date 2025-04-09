package controller

import (
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumService  *service.AlbumService
	AuthClient    *service.AuthClient
	ProfileClient *service.ProfileClient
	TripClient    *service.TripClient
}

func (c *AlbumController) CreateAlbum(ctx *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Visibility  string `json:"visibility"`
	}

	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumMapper := &models.AlbumMapper{}
	album := albumMapper.ToAlbum(req, tokenResponse)

	createdAlbum, err := c.AlbumService.CreateAlbum(album)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create trip"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAlbum)
}

func (c *AlbumController) UpdateAlbum(ctx *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Visibility  string `json:"visibility"`
	}

	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumID := ctx.Param("id")

	album, err := c.AlbumService.GetAlbumByID(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get album"})
		return
	}

	if album.UserID != tokenResponse {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you are not the owner of this album"})
		return
	}

	albumMapper := &models.AlbumMapper{}
	album = albumMapper.ToAlbum(req, tokenResponse)
	id, err := strconv.ParseUint(albumID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid album ID"})
		return
	}
	album.AlbumID = int(id)

	updatedAlbum, err := c.AlbumService.UpdateAlbum(album)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update album"})
		return
	}

	ctx.JSON(http.StatusOK, updatedAlbum)
}

func (c *AlbumController) DeleteAlbum(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albumID := ctx.Param("id")

	album, err := c.AlbumService.GetAlbumByID(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get album"})
		return
	}

	if album.UserID != tokenResponse {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you are not the owner of this album"})
		return
	}

	err = c.AlbumService.DeleteAlbum(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete album"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}

func (c *AlbumController) GetMyAlbums(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albums, err := c.AlbumService.GetAlbumsByUserID(tokenResponse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get albums"})
		return
	}

	ctx.JSON(http.StatusOK, albums)
}

func (c *AlbumController) GetPublicAlbums(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albums, err := c.AlbumService.GetPublicAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get albums"})
		return
	}

	ctx.JSON(http.StatusOK, albums)
}

func (c *AlbumController) GetAlbumByID(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albumID := ctx.Param("id")

	album, err := c.AlbumService.GetAlbumByID(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get album"})
		return
	}

	ctx.JSON(http.StatusOK, album)
}

func (c *AlbumController) GetAlbumsByUserID(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	albums, err := c.AlbumService.GetAlbumsByUserID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get albums"})
		return
	}

	ctx.JSON(http.StatusOK, albums)
}

// Albums with trips
func (c *AlbumController) GetMyAlbumsWithTrips(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albums, err := c.AlbumService.GetAlbumsByUserID(tokenResponse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get albums"})
		return
	}

	var responses []gin.H
	for _, album := range albums {
		tripList, err := c.AlbumService.GetTripsByAlbumID(strconv.Itoa(album.AlbumID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get trips for album"})
			return
		}

		var trips []models.TripMedia
		for _, tripID := range tripList {
			trip, err := c.TripClient.GetTripByID(tokenCookie, tripID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get trip details"})
				return
			}
			trips = append(trips, trip)
		}

		response := gin.H{
			"album_id":     album.AlbumID,
			"name":         album.Name,
			"description":  album.Description,
			"visibility":   album.Visibility,
			"trips_count":  len(trips),
			"trips":        trips,
		}
		responses = append(responses, response)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_albums": len(responses),
		"albums":       responses,
	})
}

func (c *AlbumController) GetPublicAlbumsWithTrips(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albums, err := c.AlbumService.GetPublicAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get albums"})
		return
	}

	var responses []gin.H
	for _, album := range albums {
		tripList, err := c.AlbumService.GetTripsByAlbumID(strconv.Itoa(album.AlbumID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get trips for album"})
			return
		}

		var trips []models.TripMedia
		for _, tripID := range tripList {
			trip, err := c.TripClient.GetTripByID(tokenCookie, tripID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get trip details"})
				return
			}
			trips = append(trips, trip)
		}

		response := gin.H{
			"album_id":     album.AlbumID,
			"name":         album.Name,
			"description":  album.Description,
			"visibility":   album.Visibility,
			"trips_count":  len(trips),
			"trips":        trips,
		}
		responses = append(responses, response)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_albums": len(responses),
		"albums":       responses,
	})
}

func (c *AlbumController) GetAlbumByIDWithTrips(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	albumID := ctx.Param("id")

	album, err := c.AlbumService.GetAlbumByID(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get album"})
		return
	}

	tripList, err := c.AlbumService.GetTripsByAlbumID(albumID)
	if err != nil {
		tripList = []uint{}
	}

	var trips []models.TripMedia
	for _, tripID := range tripList {
		trip, err := c.TripClient.GetTripByID(tokenCookie, tripID)
		if err != nil {
			continue
		}
		trips = append(trips, trip)
	}

	response := gin.H{
		"album": album,
		"trips": trips,
	}

	ctx.JSON(http.StatusOK, response)
}