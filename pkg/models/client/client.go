package client

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
)

// Config is the client config
type Config struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

// App is the app
type App struct {
	Config    Config
	OutLogger *log.Logger
	ErrLogger *log.Logger
	CertFile  string
}

func (a App) SetCreds() (*tls.Config, error) {
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
