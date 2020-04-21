package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Flukas88/newggw/pkg/models/models/server"
	json "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewApp(keyFile, certFile string) *server.App {
	var config server.ServerConfig
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

	return &server.App{
		Config:    config,
		OutLogger: outLogger,
		ErrLogger: errLogger,
		CertFile:  certFile,
		KeyFile:   keyFile,
		server:    grpc.NewServer(grpc.Creds(creds)),
	}
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func (a *server.App) SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		a.OutLogger.Println("\nClosing...")
		a.server.GracefulStop()
		os.Exit(0)
	}()
}
