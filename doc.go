// Library to easily fetch weather data through the OpenWeatherMap API.
//
// Before using the library, get an API key at https://openweathermap.org.
//
// Create an API structure :
//   api, _ := goweather.NewAPI("token", "en", "metric")
//
// Get the current weather
//   data, _ := api.Current("Lyon, FR")
//
// Get the forecast
//   data, _ := api.Forecast("Lyon, FR")
package goweather
