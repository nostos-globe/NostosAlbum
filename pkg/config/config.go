package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBPort            string
	JWTSecret         string
	AuthServiceUrl    string
	ProfileServiceUrl string
	TripsServiceUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		DBHost:            os.Getenv("DB_HOST"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		DBPort:            os.Getenv("DB_PORT"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		AuthServiceUrl:    os.Getenv("AUTH_SERVICE_URL"),
		ProfileServiceUrl: os.Getenv("PROFILE_SERVICE_URL"),
		TripsServiceUrl:   os.Getenv("TRIPS_SERVICE_URL"),
	}
}
