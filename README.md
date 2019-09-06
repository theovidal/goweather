# goweather

**🌦 Easily fetch weather information using Go and the OpenWeatherMap API.**

[![codecov badge](https://codecov.io/gh/exybore/goweather/branch/master/graph/badge.svg)](https://codecov.io/gh/exybore/goweather)
[![CircleCI Badge](https://circleci.com/gh/exybore/goweather.svg?style=svg)](https://circleci.com/gh/exybore/goweather)
[![Travis Badge](https://travis-ci.org/exybore/goweather.svg?branch=master)](https://travis-ci.org/exybore/goweather)
[![License Badge](https://img.shields.io/github/license/exybore/goweather)](./LICENSE)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/4be91f3c0fae4c02ade3d0e96c5d6149)](https://app.codacy.com/app/exybore/goweather?utm_source=github.com&utm_medium=referral&utm_content=exybore/goweather&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Badge](https://goreportcard.com/badge/github.com/exybore/goweather)](https://goreportcard.com/report/github.com/exybore/goweather)
[![Godoc Badge](https://godoc.org/github.com/exybore/goweather?status.svg)](https://godoc.org/github.com/exybore/goweather)
[![Latest version Badge](https://img.shields.io/github/v/tag/exybore/goweather?sort=semver)](https://github.com/exybore/goweather/releases)
[![Code size Badge](https://img.shields.io/github/languages/code-size/exybore/goweather)](https://github.com/exybore/goweather)
[![Open issues Badge](https://img.shields.io/github/issues/exybore/goweather)](https://github.com/exybore/goweather/issues)

- [🔧 Installation](#-installation)
- [⌨ Usage](#-usage)
  - [Initializing the API](#initializing-the-api)
  - [Get current weather](#get-current-weather)
  - [Get forecast](#get-forecast)
  - [More examples](#more-examples)
- [💻 Developing](#-developing)
- [📜 Credits](#-credits)
- [🔐 License](#-license)

## 🔧 Installation

First of all, make sure you have Go 1.12 or earlier installed on your system.

Install the library using `go get`...

```bash
go get github.com/exybore/goweather
```

... or use whatever dependency manager you like! Then, you can start writing your code !

## ⌨ Usage

### Initializing the API

Before continuing, head over to [openweathermap.org](https://openweathermap.org) and grab an API key. It's free up to 60 requests per minute, which is pretty comfortable.

After this step, you're ready to initialize the API using the `goweather.NewAPI` function :

```go
package main

import "github.com/exybore/goweather"

func main() {
  api, err := goweather.NewAPI("ez5e4fz44e84fg89ze8", "en", "metric")
  if err != nil {
    panic(err)
  }
}
```

This function takes three parameters :

- the API token that you generated using your account on openweathermap.org
- the language to use. It can be one of these Arabic - `ar`, Bulgarian - `bg`, Catalan - `ca`, Czech - `cz`, German - `de`, Greek - `el`, English - `en`, Persian (Farsi) - `fa`, Finnish - `fi`, French - `fr`, Galician - `gl`, Croatian - `hr`, Hungarian - `hu`, Italian - `it`, Japanese - `ja`, Korean - `kr`, Latvian - `la`, Lithuanian - `lt`, Macedonian - `mk`, Dutch - `nl`, Polish - `pl`, Portuguese - `pt`, Romanian - `ro`, Russian - `ru`, Swedish - `se`, Slovak - `sk`, Slovenian - `sl`, Spanish - `es`, Turkish - `tr`, Ukrainian - `ua`, Vietnamese - `vi`, Chinese Simplified - `zh_cn`, Chinese Traditional - `zh_tw`
- the unit system. It can be one of these :
  - `default` (temperatures in Kelvin)
  - `metric` (temperatures in Celsius)
  - `imperial` (temperatures in Fahrenheit)

### Get current weather

To get the current weather, use the `Current` method of the API :

```go
  weather, err := api.Current("Lyon,FR")
  if err != nil {
    panic(err)
  }
```

You can use any existing location in the parameter, but it's recommended to use the `City,Country code` format so the result is 100% accurate.

The method returns a `types.Current` structure that you can explore [on GoDoc](https://godoc.org/github.com/exybore/goweather/types#Current)

### Get forecast

To get the forecast, use the `Forecast` method of the API :

```go
  weather, err := api.Forecast("Lyon,FR")
  if err != nil {
    panic(err)
  }
```

You can use any existing location in the parameter, but it's recommended to use the `City,Country code` format so the result is 100% accurate.

The method returns a `types.Forecast` structure that you can explore [on GoDoc](https://godoc.org/github.com/exybore/goweather/types#Forecast)

### More examples

Check the [examples file](examples_test.go) for more examples of scripts using the goweather library.

## 💻 Developing

Thank you for being interested in goweather's development ! Please follow these steps :

- [Fork the repository](https://github.com/exybore/goweather/fork) on your profile
- Clone it on your computer and install the required dependencies using the `go get` command
- Make your changes
- Write tests and execute them using the `go test` command. Your API key must be stored in a `WEATHER_API_KEY` environment variable
- Commit and push your work
- Open a pull request and explain why you made these changes

## 📜 Credits

- Service : [OpenWeatherMap](https://openweathermap.org)
- Maintainer : [Exybore](https://github.com/exybore)

## 🔐 License

MIT License

Copyright (c) 2019 Exybore

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
