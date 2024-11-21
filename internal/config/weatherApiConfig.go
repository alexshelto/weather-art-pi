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
    Units  string
}

func LoadWeatherApiConfig() (WeatherApiConfig, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return WeatherApiConfig{}, fmt.Errorf("error loading .env file: %v", err)
	}

	weatherApi := WeatherApiConfig{
		ApiKey: os.Getenv("API_KEY"),
		Lat:    os.Getenv("LATITUDE"),
		Long:   os.Getenv("LONGITUDE"),
		Units:  os.Getenv("UNITS"),
	}

	if weatherApi.ApiKey == "" || weatherApi.Lat == "" || weatherApi.Long == "" || weatherApi.Units == "" {
		return WeatherApiConfig{}, fmt.Errorf("missing required environment variables")
	}

	return weatherApi, nil
}
