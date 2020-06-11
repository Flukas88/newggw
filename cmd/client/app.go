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
	configFile, loadErr := ioutil.ReadFile("client.json")
	if loadErr != nil {
		log.Fatalf("error in reading client config file: %v", loadErr)
		return nil
	}
	jsonErr := json.Unmarshal(configFile, &config)
	if jsonErr != nil {
		log.Fatalf("error in un-marshalling config JSON: %v", jsonErr)
		return nil
	}

	return &cl.App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
		CertFile:  certFile,
	}
}
