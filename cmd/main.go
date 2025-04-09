package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	controller "main/internal/api"
	dbRepo "main/internal/db"
	"main/internal/service"
	"main/pkg/config"
	"main/pkg/db"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	minioManager := config.InitMinIO()
	if minioManager == nil {
		log.Println("Faliing to init MinIO")
	}

	secretsManager := config.GetSecretsManager()
	if secretsManager != nil {
		secrets := secretsManager.LoadSecrets()
		for key, value := range secrets {
			os.Setenv(key, value)
		}
	} else {
		log.Println("Falling back to environment variables")
	}
}

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize repositories
	albumRepo := &dbRepo.AlbumRepository{DB: database}

	// Initialize clients for external apis
	authClient := &service.AuthClient{BaseURL: cfg.AuthServiceUrl}
	profileClient := &service.ProfileClient{BaseURL: cfg.ProfileServiceUrl}
	tripClient := &service.TripClient{BaseURL: cfg.TripsServiceUrl}

	// Initialize services
	albumService := &service.AlbumService{AlbumRepo: albumRepo}

	// Initialize controllers
	albumHandler := &controller.AlbumController{
		AlbumService:     	albumService,
		AuthClient:       	authClient,
		ProfileClient: 		profileClient,
		TripClient:			tripClient,
	}

	// Initialize Gin
	r := gin.Default()

	// Trip routes
	api := r.Group("/api/albums")
	{
		api.POST("/", albumHandler.CreateAlbum)
		api.GET("/", albumHandler.GetMyAlbums)
		api.PUT("/:id", albumHandler.UpdateAlbum)
		api.DELETE("/:id", albumHandler.DeleteAlbum)
		api.GET("/:id", albumHandler.GetAlbumByID)
		api.GET("/public", albumHandler.GetPublicAlbums)
	}

	// Media routes in separate group
	tripsApi := r.Group("/api/albums/trips")
	{
		tripsApi.GET("/", albumHandler.GetMyAlbumsWithTrips)
		tripsApi.GET("/public", albumHandler.GetPublicAlbumsWithTrips)
		tripsApi.GET("/:id", albumHandler.GetAlbumByIDWithTrips)
	}

	// Start server
	log.Println("Server running on http://localhost:8085")
	if err := r.Run(":8085"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
