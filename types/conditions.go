package types

import (
	"fmt"
	"strings"
)

// Structure representing weather conditions at a specific moment
type WeatherConditions struct {
	// Timestamp at which the conditions occur
	Timestamp          float64
	// Group of the conditions (Rain, snow, clouds...)
	Name               string
	// Full description of the conditions
	Description        string
	// URL to the PNG icon
	IconUrl            string
	// Emoji associated to the condition
	Emoji              rune

	// Temperature at the location
	Temperature        float64
	// Minimum temperature that can be felt
	MinimumTemperature float64
	// Maximum temperature that can be felt
	MaximumTemperature float64
	// Pressure at the location
	Pressure           float64
	// Humidity at the location (percentage)
	Humidity           float64

	// Level of clouds (percentage)
	Clouds             float64
	// Speed of the wind
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