package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alexshelto/weather-art-pi/internal/art"
	"github.com/alexshelto/weather-art-pi/internal/client"
	"github.com/alexshelto/weather-art-pi/internal/config"
)

func main() {
	// Load config from .env file
	weatherApiConfig, err := config.LoadWeatherApiConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
    fmt.Println(weatherApiConfig)

    dc := art.InitializeDrawingContext() // Initialize the drawing context


    // Create a ticker that ticks every 5 minutes
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

    // Loop to fetch weather and generate art every 5 minutes
	for {
		select {
		case <-ticker.C:

            weatherData, err := client.GetWeather(weatherApiConfig)
            if err != nil {
                fmt.Println(err.Error())
                panic(1)
            }

            fmt.Println(weatherData)
            art.DrawArt(dc, weatherData)
            // Save or update display logic here
            dc.SavePNG("current_art.png")
			//generateArt(weatherData) // Generate new art
		}
	}



	fmt.Println(weatherApiConfig)
}
