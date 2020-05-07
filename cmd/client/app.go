package main

import (
	"io/ioutil"
	"log"
	"os"

	cl "github.com/Flukas88/newggw/pkg/models/client"
	json "github.com/json-iterator/go"
)

func NewApp(certFile string) *cl.App {
	var config cl.Config
	outLogger := log.New(os.Stdout, "ClientApp - ", log.LstdFlags)
	errLogger := log.New(os.Stderr, "ClientApp - ", log.LstdFlags)
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

	return &cl.App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
		CertFile:  certFile,
	}
}
