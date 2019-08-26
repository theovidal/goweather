package goweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// The main API
type API struct {
	client http.Client
	key string
	lang string
	unit string
}

// Creates a new weather API structure
func NewAPI(key string, lang string, unit string) (API, error) {
	if !contains(locales, lang) {
		return API{}, errors.New(fmt.Sprintf("cannot use %s as lang", lang))
	}
	if !contains(units, unit) {
		return API{}, errors.New(fmt.Sprintf("cannot use %s as unit", unit))
	}

	client := http.Client{
		Timeout: time.Second * 4,
	}

	return API{
		client: client,
		key: key,
		lang: lang,
		unit: unit,
	}, nil
}

// Gets the current weather for a specific location
func (api API) Current(location string) (CurrentWeather, error) {
	result, err := api.getData(endpoints["current"], location)
	if err != nil {
		return CurrentWeather{}, err
	}

  current := CurrentWeather{}
  err = json.Unmarshal(result, &current)
  if err != nil {
  	return CurrentWeather{}, err
	}

  return current, nil
}

// Gets data from the OpenWeatherMap API
func (api API) getData(requestUrl string, location string) ([]byte, error) {
	request, err := api.encodeRequest(requestUrl, location)
	if err != nil {
		return []byte{}, err
	}

	response, err := api.makeRequest(request)
	if err != nil {
		return []byte{}, err
	}

  body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

  return body, nil
}

// Encodes a ready-to-use request structure from an URL and a location
func (api API) encodeRequest(requestUrl string, location string) (*http.Request, error) {
	params := url.Values{}
	params.Add("APPID", api.key)
	params.Add("q", location)
	params.Add("units", api.unit)
	params.Add("lang", api.lang)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(requestUrl, "?", params.Encode()), nil)
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Add("User-Agent", "goweather")

	return req, nil
}

// Makes a request to the OpenWeatherMap API from a request structure
func (api API) makeRequest(request *http.Request) (*http.Response, error) {
	response, err := api.client.Do(request)
  if err != nil {
  	return &http.Response{}, err
	}

	switch response.StatusCode {
	case 401:
		return &http.Response{}, errors.New("unauthorized API key")
	case 404:
		return &http.Response{}, errors.New("unknown location")
	}

	return response, nil
}