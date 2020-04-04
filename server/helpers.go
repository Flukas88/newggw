package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"
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

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler(s *grpc.Server) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nClosing...")
		s.GracefulStop()
		os.Exit(0)
	}()
}
