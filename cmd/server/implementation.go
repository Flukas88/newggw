package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	srv "github.com/Flukas88/newggw/pkg/models/server"
	json "github.com/json-iterator/go"
)

func GetRespJSON(city string) (srv.OpenWeather, error) {
	token := os.Getenv("GGW")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, token)
	resp, getErr := http.Get(url)
	if getErr != nil {
		return srv.OpenWeather{}, errors.New("error in trying to get to openweathermap")
	}
	// converts as needed for json.Unmarshal
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return srv.OpenWeather{}, errors.New("error in trying to get parse the body of the api req")
	}
	var weather srv.OpenWeather

	jsonError := json.Unmarshal(body, &weather)

	if jsonError != nil {
		return srv.OpenWeather{}, errors.New("error in unmarshalling the json response")
	}

	if weather.Cod != 200 {
		retCode := weather.Cod
		errMsg := fmt.Sprintf("API returned code %d. Look for explanation at http://openweathermap.org/faq#error%d", retCode, retCode)
		return srv.OpenWeather{}, errors.New(errMsg)
	}
	return weather, nil
}
