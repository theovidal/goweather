package tests

import (
	"github.com/exybore/goweather"
	"testing"
)

// Gets the current weather and displays some information about it
func TestCurrent(t *testing.T) {
		api, err := goweather.NewAPI("", "fr", "metric")
		if err != nil {
			t.Error("error:", err)
			return
		}
    result, err := api.Current("Lyon,FR")
    if err != nil {
    	t.Error("error:", err)
    	return
		}
    t.Log(result)
    t.Log(result.City.Name)
    t.Log(result.Conditions.Timestamp)
}