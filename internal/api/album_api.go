package controller

import (
	"fmt"
	"main/internal/models"
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumService *service.AlbumService
	AuthClient *service.AuthClient
	ProfileClient *service.ProfileClient
}

func (c *AlbumController) CreateAlbum(ctx *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Visibility  string `json:"visibility"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
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

	createdAlbum, err := c.AlbumService.CreateAlbum(req, tokenResponse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create trip"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAlbum)
}