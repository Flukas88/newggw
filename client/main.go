package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc/credentials"

	"github.com/Flukas88/newggw/proto/ggwpb"
	"google.golang.org/grpc"
)

var config ClientConfig
var version = "dev"

func main() {

	binArgs := os.Args[1:]
	if len(binArgs) < 2 {
		log.Printf("City or degrees not provided.")
		os.Exit(2)
	}

	// Reading config
	configFile, err := ioutil.ReadFile("client.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Error in un-marshalling config JSON")
	}

	address := fmt.Sprintf("%s:%d", config.Server, config.Port)
	log.Printf("Connecting client (version %s) to server on %s:%d ...", version, config.Server, config.Port)

	b, _ := ioutil.ReadFile("./certs/ca.cert")
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Printf("credentials: failed to append certificates")
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}
	cc, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	city := binArgs[0]
	degrees := binArgs[1]

	client := ggwpb.NewGgwClient(cc)
	request := &ggwpb.GgwRequest{City: city, Degrees: degrees}

	resp, clErr := client.Ggw(context.Background(), request)
	if clErr != nil {
		log.Fatal(clErr)
	}
	fmt.Printf("Temp in %s is %s (%s degrees)", resp.City, resp.Temp, resp.Degrees)
}
