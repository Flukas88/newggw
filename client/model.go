package main

import (
	"log"
)

// ClientConfig is the client config
type ClientConfig struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

// App is the app
type App struct {
	Config    ClientConfig
	OutLogger *log.Logger
	ErrLogger *log.Logger
	CertFile  string
}
