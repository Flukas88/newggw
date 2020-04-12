package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"os"

	json "github.com/json-iterator/go"
)

func NewApp(certFile string) *App {
	var config ClientConfig
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

	return &App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
		CertFile:  certFile,
	}
}

func (a App) setCreds() (*tls.Config, error) {
	b, _ := ioutil.ReadFile(a.CertFile)
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, errors.New("credentials: failed to append certificates")
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}
	return config, nil
}
