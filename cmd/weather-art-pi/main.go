package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/alexshelto/weather-art-pi/internal/config"
)

func main() {
	// Command-line flag to specify mock temperature
	mockTemp := flag.Float64("temp", -999, "Mock temperature (°F) to use for testing (bypasses API call)")
	flag.Parse()

	// Load config from .env file
	weatherApiConfig, err := config.LoadWeatherApiConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Check if mock temperature is provided
	if *mockTemp != -999 {
		fmt.Printf("Using mock temperature: %.2f°F\n", *mockTemp)
	} else {
		fmt.Println("Not using mock temperature")
	}

	fmt.Println(weatherApiConfig)
}
