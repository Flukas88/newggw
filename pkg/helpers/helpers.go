package helpers

import (
	"strconv"
)

// Conversion function type
type ConvertFn func(float64) string

// Kelvin2Celsius converts from K -> C
func Kelvin2Celsius(k float64) string {
	return strconv.FormatFloat(k-273.15, 'f', 1, 64)
}

// Kelvin2Fahrenheit converts from K -> F
func Kelvin2Fahrenheit(k float64) string {
	return strconv.FormatFloat(9/5*(k-273)+32, 'f', 1, 64)
}
