package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type WeatherApiConfig struct {
	ApiKey string
	Lat    string
	Long   string
}

func LoadWeatherApiConfig() (*WeatherApiConfig, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	weatherApi := &WeatherApiConfig{
		ApiKey: os.Getenv("API_KEY"),
		Lat:    os.Getenv("LATITUDE"),
		Long:   os.Getenv("LONGITUDE"),
	}

	if weatherApi.ApiKey == "" || weatherApi.Lat == "" || weatherApi.Long == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return weatherApi, nil
}
