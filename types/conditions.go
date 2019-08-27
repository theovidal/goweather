package types

type WeatherConditions struct {
	Timestamp float64
}

func ParseWeatherConditions(data map[string]interface{}) WeatherConditions {
	return WeatherConditions{
		Timestamp: data["dt"].(float64),
	}
}
