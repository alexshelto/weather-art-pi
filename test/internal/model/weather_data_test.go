package model

import (
	"testing"

	"github.com/alexshelto/weather-art-pi/internal/client"
	"github.com/alexshelto/weather-art-pi/internal/model"
)


func TestCanParseFromWeatherResponse(t *testing.T) {
    eTemp := float64(10.1)
    eWind := float64(11.2)
    eRain := float64(12.3)

    response := client.WeatherResponse{
        Main: struct {
            Temp float64 `json:"temp"`
        }{Temp: eTemp},
        Wind: struct {
            Speed float64 `json:"speed"`
        }{Speed: eWind},
        Rain: struct {
            OneHour float64 `json:"1h"`
        }{OneHour: eRain},
    }

    wd := model.FromWeatherResponse(response)

    if wd.Temp != eTemp {
        t.Errorf("Temp was incorrect, got: %f, want: %f.", wd.Temp, eTemp) 
    }
    if wd.Wind != eWind {
        t.Errorf("Wind was incorrect, got: %f, want: %f.", wd.Wind, eWind)
    }
    if wd.Temp != eRain {
        t.Errorf("Rain was incorrect, got: %f, want: %f.", wd.Rain, eRain)
    }

}

