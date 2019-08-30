// Easily fetch weather information using Go and the OpenWeatherMap API
//
// Before playing around with the library, get an API key at https://openweathermap.org.
//
// Create an API structure :
//   api, _ := goweather.NewAPI("token", "en", "metric")
//
// Get the current weather using this API :
//   data, _ := api.Current("Lyon, FR")
//
// Get the forecast using this API :
//   data, _ := api.Forecast("Lyon, FR")
package goweather
