package main

import (
	"testing"

	json "github.com/json-iterator/go"
)

func Test_getCityInfo(t *testing.T) {
	var milanC CityTemp
	var milanF CityTemp
	sampleJSON := `{"coord":{"lon":-0.13,"lat":51.51},"weather":[{"id":300,"main":"Drizzle","description":"light intensity drizzle","icon":"09d"}],"base":"stations","main":{"temp":303.15,"pressure":1012,"humidity":81,"temp_min":279.15,"temp_max":281.15},"visibility":10000,"wind":{"speed":4.1,"deg":80},"clouds":{"all":90},"dt":1485789600,"sys":{"type":1,"id":5091,"message":0.0103,"country":"GB","sunrise":1485762037,"sunset":1485794875},"id":2643743,"name":"Milan","cod":200}`
	var data OpenWeather
	_ = json.Unmarshal([]byte(sampleJSON), &data)
	_ = milanC.getCityInfo("Milan", "C", data)
	_ = milanF.getCityInfo("Milan", "F", data)

	tt := []struct {
		name     string
		expected CityTemp
		got      CityTemp
	}{
		{"Milan, C", CityTemp{City: "Milan", Temp: "30.0", Degrees: "C"}, milanC},
		{"Milan, F", CityTemp{City: "Milan", Temp: "62.1", Degrees: "F"}, milanF},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected.Temp != tc.got.Temp {
				t.Fatalf("Case %v should be %v; got %v", tc.name, tc.expected.Temp, tc.got.Temp)
			} else if tc.expected.City != tc.got.City {
				t.Fatalf("Case %v should be %v; got %v", tc.name, tc.expected.City, tc.got.City)
			} else if tc.expected.Degrees != tc.got.Degrees {
				t.Fatalf("Case %v should be %v; got %v", tc.name, tc.expected.Degrees, tc.got.Degrees)
			}
		})
	}
}
