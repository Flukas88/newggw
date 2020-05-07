package main

import (
	"io/ioutil"
	"log"
	"os"

	srv "github.com/Flukas88/newggw/pkg/models/server"
	json "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewApp(keyFile, certFile string) *srv.App {
	var config srv.ServerConfig
	outLogger := log.New(os.Stdout, "ServerApp - ", log.LstdFlags)
	errLogger := log.New(os.Stderr, "ServerApp - ", log.LstdFlags)
	// Reading config
	configFile, err := ioutil.ReadFile("server.json")
	if err != nil {
		errLogger.Fatal(err.Error())
		return nil
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		errLogger.Fatal("Error in un-marshalling config JSON")
		return nil
	}

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		errLogger.Println("failed to setup TLS")
		return nil
	}

	return &srv.App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
		CertFile:  certFile,
		KeyFile:   keyFile,
		Server:    grpc.NewServer(grpc.Creds(creds)),
	}
}
