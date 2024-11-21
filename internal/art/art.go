package art

import (
	"github.com/alexshelto/weather-art-pi/internal/client"


)

const W, H = 800, 400

// InitializeDrawingContext creates and returns a new drawing context.
func InitializeDrawingContext() *gg.Context {
	return gg.NewContext(W, H)
}

// DrawArt draws on the given context based on the weather data.
func DrawArt(dc *gg.Context, wd client.WeatherResponse) {
	dc.Clear() // Clear the canvas for the new drawing

	// Example drawing commands
	dc.DrawCircle(400, 300, 100)
	dc.SetRGB(0, 0, 1) // Blue color
	dc.Fill()

}
