// This package contains all the types associated to the library : weather conditions, cities...
package types

import (
	"fmt"
	"strings"
)

// Structure representing weather conditions at a specific moment
type WeatherConditions struct {
	Timestamp          float64
	Name               string
	Description        string
	IconUrl            string
	Emoji              rune

	Temperature        float64
	MinimumTemperature float64
	MaximumTemperature float64
	Pressure           float64
	Humidity           float64

	Clouds             float64
	WindSpeed          float64
}

// Returns weather conditions parsed in a structure
func ParseWeatherConditions(data map[string]interface{}) WeatherConditions {
	clouds := data["clouds"].(map[string]interface{})
	main := data["main"].(map[string]interface{})
	wind := data["wind"].(map[string]interface{})
	misc := data["weather"].([]interface{})[0].(map[string]interface{})

	iconUrl := fmt.Sprintf("https://openweathermap.org/img/w/%s.png", misc["icon"].(string))
	emoji := ConditionEmojis[strings.ReplaceAll(misc["icon"].(string), "n", "d")]

	return WeatherConditions{
		Timestamp:          data["dt"].(float64),
		Name:               misc["main"].(string),
		Description:        misc["description"].(string),
		IconUrl:            iconUrl,
		Emoji:              emoji,

		Temperature:        main["temp"].(float64),
		MinimumTemperature: main["temp_min"].(float64),
		MaximumTemperature: main["temp_max"].(float64),
		Pressure:           main["pressure"].(float64),
		Humidity:           main["humidity"].(float64),

		Clouds:             clouds["all"].(float64),
		WindSpeed:          wind["speed"].(float64),
	}
}
