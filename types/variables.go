// goweather - A library to easily fetch weather data through the OpenWeatherMap API.
package types

// Address of the OpenWeatherMap API
const ApiUrl = "https://api.openweathermap.org"

// Available endpoints through the API
var Endpoints = map[string]string{
	"current":  ApiUrl + "/data/2.5/weather",
	"forecast": ApiUrl + "/data/2.5/forecast",
}

// Accepted units
var Units = []string{"default", "metric", "imperial"}

// Accepted locales
var Locales = []string{"ar", "bg", "ca", "cz", "de", "el", "fa", "fi", "fr", "gl", "hr", "hu",
	"it", "ja", "kr", "la", "lt", "mk", "nl", "pl", "pt", "ro", "ru", "se", "sk", "sl", "es", "tr",
	"ua", "vi", "zh_cn", "zh_tw", "en"}

// OpenWeatherMap condition codes translated in emojis
var ConditionEmojis = map[string]rune{
	"01d": '☀',
	"02d": '⛅',
	"03d": '☁',
	"04d": '☁',
	"09d": '🌧',
	"10d": '🌦',
	"11d": '🌩',
	"13d": '🌨',
	"50d": '🌫',
}
