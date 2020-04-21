package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/Flukas88/newggw/pkg/models/models/client"
	json "github.com/json-iterator/go"
)

func NewApp(certFile string) *client.App {
	var config client.ClientConfig
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

	return &client.App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
		CertFile:  certFile,
	}
}

func (a client.App) setCreds() (*tls.Config, error) {
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
