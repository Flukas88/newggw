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
	var config srv.Config
	outLogger := log.New(os.Stdout, "ServerApp - ", log.LstdFlags)
	errLogger := log.New(os.Stderr, "ServerApp - ", log.LstdFlags)
	// Reading config
	configFile, loadErr := ioutil.ReadFile("server.json")
	if loadErr != nil {
		errLogger.Fatalf("error in reading server config file: %v", loadErr)
		return nil
	}
	jsonErr := json.Unmarshal(configFile, &config)
	if jsonErr != nil {
		errLogger.Fatalf("error in un-marshalling config JSON: %v", jsonErr)
		return nil
	}

	creds, credErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if credErr != nil {
		errLogger.Fatalf("failed to setup TLS: %v", credErr)
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
