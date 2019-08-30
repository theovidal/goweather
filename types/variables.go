package types

// APIURL contains the address of the OpenWeatherMap API
const APIURL = "https://api.openweathermap.org"

// Endpoints contains available endpoints through the API
var Endpoints = map[string]string{
	"current":  APIURL + "/data/2.5/weather",
	"forecast": APIURL + "/data/2.5/forecast",
}

// Units contains accepted units in the API
var Units = []string{"default", "metric", "imperial"}

// Locales contains accepted locales in the API
var Locales = []string{"ar", "bg", "ca", "cz", "de", "el", "fa", "fi", "fr", "gl", "hr", "hu",
	"it", "ja", "kr", "la", "lt", "mk", "nl", "pl", "pt", "ro", "ru", "se", "sk", "sl", "es", "tr",
	"ua", "vi", "zh_cn", "zh_tw", "en"}

// ConditionEmojis contains OpenWeatherMap condition codes translated in emojis
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
