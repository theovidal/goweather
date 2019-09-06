package goweather

import (
	"os"
	"testing"
)

// Explicitly pass an unknown key to the API
func TestUnknownKey(t *testing.T) {
	api, err := NewAPI("ilovefrance", "fr", "metric")
	if err != nil {
		t.Error("error:", err)
		return
	}
	_, err = api.Current("Lyon,FR")
	if err == nil {
		t.Error("no error was thrown")
		return
	}
}

// Explicitly pass an unknown locale to the API
func TestUnknownLocale(t *testing.T) {
	_, err := NewAPI(os.Getenv("WEATHER_API_KEY"), "pnl", "metric")
	if err == nil {
		t.Error("no error was thrown")
	} else {
		t.Log(err)
	}
}

// Explicitly pass an unknown unit system to the API
func TestUnknownUnit(t *testing.T) {
	_, err := NewAPI(os.Getenv("WEATHER_API_KEY"), "fr", "estonia")
	if err == nil {
		t.Error("no error was thrown")
	} else {
		t.Log(err)
	}
}

// Explicitly pass an unknown location to the API
func TestUnknownLocation(t *testing.T) {
	api, err := NewAPI(os.Getenv("WEATHER_API_KEY"), "fr", "metric")
	if err != nil {
		t.Error("error:", err)
		return
	}
	_, err = api.Current("ExyboreLand,FR")
	if err == nil {
		t.Error("no error was thrown")
		return
	}
}
