package main

import (
	"strconv"
)

// Conversion function type
type convertFn func(float64) string

// kelvin2Celsius converts from K -> C
func kelvin2Celsius(k float64) string {
	return strconv.FormatFloat(k-273.15, 'f', 1, 64)
}

// kelvin2Fahrenheit converts from K -> F
func kelvin2Fahrenheit(k float64) string {
	return strconv.FormatFloat(9/5*(k-273)+32, 'f', 1, 64)
}
