package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Flukas88/newggw/pkg/helpers"
	"google.golang.org/grpc"
)

// OpenWeather is the JSON structure of a api call
type OpenWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

// CityTemp is the JSON response
type CityTemp struct {
	City    string `json:"city"`
	Temp    string `json:"temp"`
	Degrees string `json:"degrees"`
}

// Config is the JSON config for the server
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// App is the app
type App struct {
	Config    Config
	OutLogger *log.Logger
	ErrLogger *log.Logger
	CertFile  string
	KeyFile   string
	Server    *grpc.Server
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func (a *App) SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		a.OutLogger.Println("\nClosing...")
		a.Server.GracefulStop()
		os.Exit(0)
	}()
}

// GetCityInfo gets "city" weather in "degrees"
func (c *CityTemp) GetCityInfo(city, degrees string, data OpenWeather) error {
	Conversions := map[string]helpers.ConvertFn{
		"C": helpers.Kelvin2Celsius,
		"F": helpers.Kelvin2Fahrenheit,
	}

	c.City = city
	c.Temp = Conversions[degrees](data.Main.Temp)
	c.Degrees = degrees

	return nil
}
