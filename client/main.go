package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Flukas88/newggw/proto/ggwpb"
	"google.golang.org/grpc"
)

var config ClientConfig
var version = "dev"

func main() {

	// Reading config
	configFile, err := ioutil.ReadFile("client.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(configFile), &config)
	if err != nil {
		log.Fatal("Error in un-marshalling config JSON")
	}

	address := fmt.Sprintf("%s:%d", config.Server, config.Port)
	log.Printf("Connecting client (version %s) to server on %s:%d ...", version, config.Server, config.Port)
	argsWithoutProg := os.Args[1:]

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(address, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := ggwpb.NewGgwClient(cc)
	request := &ggwpb.WheaterRequest{City: argsWithoutProg[0], Degrees: argsWithoutProg[1]}

	resp, clErr := client.Now(context.Background(), request)
	if clErr != nil {
		log.Fatal(clErr)
	}
	fmt.Printf("Temp in %s is %s (%s degrees)", resp.City, resp.Temp, resp.Degrees)
}
