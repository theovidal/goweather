package types

type City struct {
	Name        string
	Country     string
	Sunrise     float64
	Sunset      float64
	Coordinates Coordinates
}

type Coordinates struct {
  Longitude float64
  Latitude  float64
}
