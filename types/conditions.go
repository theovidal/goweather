// This package contains all the types associated to the library : weather conditions, cities...
package types

// Structure representing weather conditions at a specific moment
type WeatherConditions struct {
	Timestamp float64
}

// Returns weather conditions parsed in a structure
func ParseWeatherConditions(data map[string]interface{}) WeatherConditions {
	return WeatherConditions{
		Timestamp: data["dt"].(float64),
	}
}
