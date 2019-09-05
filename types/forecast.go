package types

import "encoding/json"

// Forecast represents weather forecast
type Forecast struct {
	// The city where the weather applies
	City City
	// A list of all the weather conditions that will apply in the near future
	Conditions []WeatherConditions
}

// UnmarshalJSON parses data returned from the API into the structure
func (forecast *Forecast) UnmarshalJSON(raw []byte) (err error) {
	var data map[string]interface{}
	if err = json.Unmarshal(raw, &data); err != nil {
		return
	}

	city := data["city"].(map[string]interface{})
	coords := city["coord"].(map[string]interface{})

	forecast.City = City{
		Name:    city["name"].(string),
		Country: city["country"].(string),
		Sunrise: city["sunrise"].(float64),
		Sunset:  city["sunset"].(float64),
		Coordinates: Coordinates{
			Longitude: coords["lon"].(float64),
			Latitude:  coords["lat"].(float64),
		},
	}

	conditionsList := data["list"].([]interface{})
	for _, rawCondition := range conditionsList {
		conditionData := rawCondition.(map[string]interface{})
		forecast.Conditions = append(forecast.Conditions, ParseWeatherConditions(conditionData))
	}

	return
}
