package client

import (
	"encoding/json"
	"fmt"
	"github.com/alexshelto/weather-art-pi/internal/config"
	"net/http"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Rain struct {
		OneHour float64 `json:"1h"`
	} `json:"rain,omitempty"`
}

func (wr WeatherResponse) GetTemp() float64 {
	return wr.Main.Temp
}

func (wr WeatherResponse) GetWindspeed() float64 {
	return wr.Wind.Speed
}

func (wr WeatherResponse) GetRain() float64 {
	return wr.Rain.OneHour
}

func (wr WeatherResponse) String() string {
    return fmt.Sprintf("Temperature: %.2f°F\nWind Speed: %.2f mph\nRain (last hour): %.2f mm",
    wr.Main.Temp, wr.Wind.Speed, wr.Rain.OneHour)
}

// GetWeather makes an API call to fetch the essential weather data
func GetWeather(weatherConfig config.WeatherApiConfig) (WeatherResponse, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=%s",
		weatherConfig.Lat,
		weatherConfig.Long,
		weatherConfig.ApiKey,
        weatherConfig.Units,
	)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("failed to fetch weather data: %s", resp.Status)
	}

	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return WeatherResponse{}, err
	}
	return weatherResponse, nil
}
