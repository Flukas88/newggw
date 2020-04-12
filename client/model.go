package main

// ClientConfig is the client config
type ClientConfig struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

type App struct {
	Config ClientConfig
}
