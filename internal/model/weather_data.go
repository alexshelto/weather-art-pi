package model

import "github.com/alexshelto/weather-art-pi/internal/client"

type WeatherData struct {
    Temp float64
    Wind float64
    Rain float64
}

func FromWeatherResponse(wr client.WeatherResponse) WeatherData {
    return WeatherData{
        wr.GetTemp(),
        wr.GetWindspeed(),
        wr.GetRain(),
    }
}


