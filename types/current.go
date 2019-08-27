package types

import "encoding/json"

type Current struct {
	City       City
	Conditions WeatherConditions
}

func (current *Current) UnmarshalJSON(raw []byte) error {
 	var data map[string]interface{}
 	if err := json.Unmarshal(raw, &data); err != nil {
 		return err
 	}

 	coord := data["coord"].(map[string]interface{})
 	sys := data["sys"].(map[string]interface{})

 	current.City = City{
 		Name:    data["name"].(string),
 		Country: sys["country"].(string),
 		Sunrise: sys["sunrise"].(float64),
 		Sunset:  sys["sunset"].(float64),
 		Coordinates: Coordinates{
 			Longitude: coord["lon"].(float64),
 			Latitude:  coord["lat"].(float64),
		},
 	}
  current.Conditions = ParseWeatherConditions(data)

 	return nil
}
