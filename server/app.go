package main

import (
	"io/ioutil"
	"log"
	"os"

	json "github.com/json-iterator/go"
)

func NewApp() *App {
	var config ServerConfig
	outLogger := log.New(os.Stdout, "ServerApp - ", log.LstdFlags)
	errLogger := log.New(os.Stderr, "ServerApp - ", log.LstdFlags)
	// Reading config
	configFile, err := ioutil.ReadFile("server.json")
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Error in un-marshalling config JSON")
		return nil
	}

	return &App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
	}
}
