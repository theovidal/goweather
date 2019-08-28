package types

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
	return WeatherConditions{
		Timestamp: data["dt"].(float64),
	}
}
