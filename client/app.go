package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func NewApp() *App {
	var config ClientConfig

	// Reading config
	configFile, err := ioutil.ReadFile("client.json")
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
		Config: config,
	}
}