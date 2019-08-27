// This package contains all the types associated to the library : weather conditions, cities...
package types

// Structure representing weather forecast
type Forecast struct {
	// The city where the weather applies
	City       City

	// A list of all the weather conditions that will apply in the near future
	Conditions []WeatherConditions
}
