package goweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/exybore/goweather/types"
)

// API is the main piece of the library : it centralizes requests made to OpenWeatherMap and parameters.
type API struct {
	// The HTTP client to use to make requests
	client http.Client

	// The OpenWeatherMap API key
	key string

	// Lang of the returned data
	lang string

	// Units of the returned data
	unit string
}

// NewAPI creates a new weather API structure
func NewAPI(key string, lang string, unit string) (api API, err error) {
	if !contains(types.Locales, lang) {
		return api, fmt.Errorf("cannot use %s as lang", lang)
	}
	if !contains(types.Units, unit) {
		return api, fmt.Errorf("cannot use %s as unit", unit)
	}

	client := http.Client{
		Timeout: time.Second * 4,
	}

	api = API{
		client: client,
		key:    key,
		lang:   lang,
		unit:   unit,
	}

	return
}

// Current gets current weather information for a specific location
func (api API) Current(location string) (current types.Current, err error) {
	result, err := api.getData(types.Endpoints["current"], location)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &current)
	if err != nil {
		return
	}

	return
}

// Forecast get forecasts for a specific location
func (api API) Forecast(location string) (forecast types.Forecast, err error) {
	result, err := api.getData(types.Endpoints["forecast"], location)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &forecast)
	if err != nil {
		return
	}

	return
}

// Gets data from the OpenWeatherMap API
func (api API) getData(requestURL string, location string) (data []byte, err error) {
	request, err := api.encodeRequest(requestURL, location)
	if err != nil {
		return
	}

	response, err := api.makeRequest(request)
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return data, nil
}

// Encodes a ready-to-use request structure from an URL and a location
func (api API) encodeRequest(requestURL string, location string) (request *http.Request, err error) {
	params := url.Values{}
	params.Add("APPID", api.key)
	params.Add("q", location)
	params.Add("units", api.unit)
	params.Add("lang", api.lang)

	request, err = http.NewRequest(http.MethodGet, fmt.Sprint(requestURL, "?", params.Encode()), nil)
	if err != nil {
		return
	}
	request.Header.Add("User-Agent", "goweather")

	return
}

// Makes a request to the OpenWeatherMap API from a request structure
func (api API) makeRequest(request *http.Request) (response *http.Response, err error) {
	response, err = api.client.Do(request)
	if err != nil {
		return
	}

	switch response.StatusCode {
	case 401:
		return response, errors.New("unauthorized API key")
	case 404:
		return response, errors.New("unknown location")
	}

	return
}
