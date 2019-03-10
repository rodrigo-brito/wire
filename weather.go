package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

const baseURL = "http://api.openweathermap.org/data/2.5/weather"

type WeatherService interface {
	ByCity(cityName string) (string, error)
}

type WeatherServiceOption func(*weatherService)

type weatherService struct {
	key string
}

func (w *weatherService) ByCity(cityName string) (string, error) {
	client := new(http.Client)
	response, err := client.Get(fmt.Sprintf("%s?q=%s&APPID=%s", baseURL, cityName, w.key))
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.New("error on fetch city weather")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	tempMin := gjson.GetBytes(body, "main.temp_min")
	tempMax := gjson.GetBytes(body, "main.temp_max")
	forecast := gjson.GetBytes(body, "weather.0.main")

	return fmt.Sprintf("%s - Min %.2fF, Max %.2f", forecast.Str, tempMin.Num, tempMax.Num), nil
}

func NewWeatherService(options ...WeatherServiceOption) WeatherService {
	service := new(weatherService)

	for _, option := range options {
		option(service)
	}

	return service
}

func WithAPIKey(key string) WeatherServiceOption {
	return func(service *weatherService) {
		service.key = key
	}
}
