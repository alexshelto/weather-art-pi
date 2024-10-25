package client

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/alexshelto/weather-art-pi/internal/client"
)

var mockJSON_noRain string = `{
    "coord": {
        "lon": -84.2675,
        "lat": 39.2719
    },
    "weather": [
        {
            "id": 802,
            "main": "Clouds",
            "description": "scattered clouds",
            "icon": "03d"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 60,
        "feels_like": 63.21,
        "temp_min": 61.72,
        "temp_max": 66.96,
        "pressure": 1019,
        "humidity": 57,
        "sea_level": 1019,
        "grnd_level": 991
    },
    "visibility": 10000,
    "wind": {
        "speed": 3,
        "deg": 225,
        "gust": 8.01
    },
    "clouds": {
        "all": 49
    },
    "dt": 1729871722,
    "sys": {
        "type": 2,
        "id": 2003504,
        "country": "US",
        "sunrise": 1729857460,
        "sunset": 1729896254
    },
    "timezone": -14400,
    "id": 4517140,
    "name": "Loveland",
    "cod": 200
}`

var mockJSON string = `{
    "coord": {
        "lon": -84.2675,
        "lat": 39.2719
    },
    "weather": [
        {
            "id": 802,
            "main": "Clouds",
            "description": "scattered clouds",
            "icon": "03d"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 60,
        "feels_like": 63.21,
        "temp_min": 61.72,
        "temp_max": 66.96,
        "pressure": 1019,
        "humidity": 57,
        "sea_level": 1019,
        "grnd_level": 991
    },
    "visibility": 10000,
    "wind": {
        "speed": 3,
        "deg": 225,
        "gust": 8.01
    },
    "rain": {
      "1h": 2.79
   },
    "clouds": {
        "all": 49
    },
    "dt": 1729871722,
    "sys": {
        "type": 2,
        "id": 2003504,
        "country": "US",
        "sunrise": 1729857460,
        "sunset": 1729896254
    },
    "timezone": -14400,
    "id": 4517140,
    "name": "Loveland",
    "cod": 200
}`

func TestWeatherParse_NoRain(t *testing.T) {
	mockResponse := io.NopCloser(strings.NewReader(mockJSON_noRain))

	// Create a WeatherResponse struct to decode into
	var weatherResponse client.WeatherResponse

	// Decode the mock response into the struct
	if err := json.NewDecoder(mockResponse).Decode(&weatherResponse); err != nil {
		t.Fatalf("failed to decode mock response: %v", err)
	}

	// Validate the response
	if weatherResponse.GetTemp() != 60 {
		t.Errorf("Expected temperature to be 60, got %f", weatherResponse.GetTemp())
	}
	if weatherResponse.GetWindspeed() != 3 {
		t.Errorf("Expected wind speed be 3, got %f", weatherResponse.GetWindspeed())
	}
	if weatherResponse.GetRain() != 0 {
		t.Errorf("Expected rain to be 0, got %f", weatherResponse.GetRain())
	}
}

func TestWeatherParse_withRain(t *testing.T) {
	mockResponse := io.NopCloser(strings.NewReader(mockJSON))

	// Create a WeatherResponse struct to decode into
	var weatherResponse client.WeatherResponse

	// Decode the mock response into the struct
	if err := json.NewDecoder(mockResponse).Decode(&weatherResponse); err != nil {
		t.Fatalf("failed to decode mock response: %v", err)
	}

	// Validate the response
	if weatherResponse.GetTemp() != 60 {
		t.Errorf("Expected temperature to be 60, got %f", weatherResponse.GetTemp())
	}
	if weatherResponse.GetWindspeed() != 3 {
		t.Errorf("Expected wind speed be 3, got %f", weatherResponse.GetWindspeed())
	}
	if weatherResponse.GetRain() != 2.79 {
		t.Errorf("Expected rain to be 2.79, got %f", weatherResponse.GetRain())
	}
}
