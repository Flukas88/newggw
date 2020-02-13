package main

import "testing"

func Test_kelvin2Celsius(t *testing.T) {
	expected := "30.0"
	got := kelvin2Celsius(303.15)

	if expected != got {
		t.Errorf("kelvin2Celsius is wrong. Expecting %v, got %v", expected, got)
	}
}

func Test_kelvin2Fahrenheit(t *testing.T) {
	expected := "62.1"
	got := kelvin2Fahrenheit(303.15)
	if expected != got {
		t.Errorf("kelvin2Celsius is wrong. Expecting %v, got %v", expected, got)
	}
}
