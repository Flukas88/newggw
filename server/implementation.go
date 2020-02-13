package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getRespJson(city string) (OpenWeather, error) {
	token := os.Getenv("GGW")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, token)
	resp, getErr := http.Get(url)
	if getErr != nil {
		return OpenWeather{}, errors.New("error in trying to get to openweathermap")
	}
	// converts as needed for json.Unmarshal
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return OpenWeather{}, errors.New("error in trying to get parse the body of the api req")
	}
	var weather OpenWeather

	jsonError := json.Unmarshal(body, &weather)

	if jsonError != nil {
		return OpenWeather{}, errors.New("error in unmarshalling the json response")
	}

	if weather.Cod != 200 {
		retCode := weather.Cod
		errMsg := fmt.Sprintf("API returned code %d. Look for explanation at http://openweathermap.org/faq#error%d", retCode, retCode)
		return OpenWeather{}, errors.New(errMsg)
	}
	return weather, nil
}

// GetCityInfo gets "city" weather in "degrees"
func (c *CityTemp) getCityInfo(city, degrees string, data OpenWeather) error {
	Conversions := map[string]convertFn{
		"C": kelvin2Celsius,
		"F": kelvin2Fahrenheit,
	}

	c.City = city
	c.Temp = Conversions[degrees](data.Main.Temp)
	c.Degrees = degrees

	return nil
}
