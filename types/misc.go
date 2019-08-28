package types

// Structure representing a real city and its data
type City struct {
	// The name of the city, e.g. "Lyon"
	Name        string

	// Country code of the city, e.g. "FR"
	Country     string

	// The timestamp at which the sun rises
	Sunrise     float64

	// The timestamp at which the sun sets
	Sunset      float64

	// Coordinates of the city
	Coordinates Coordinates
}

// Structure representing coordinates of a specific location
type Coordinates struct {
  Longitude float64
  Latitude  float64
}
