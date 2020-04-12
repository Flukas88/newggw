package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/grpc/credentials"

	"github.com/Flukas88/newggw/proto/ggwpb"
	"google.golang.org/grpc"
)

var version = "dev"

func main() {
	app := NewApp()
	binArgs := os.Args[1:]
	if len(binArgs) < 2 {
		app.ErrLogger.Printf("City or degrees not provided.")
		os.Exit(2)
	}

	address := fmt.Sprintf("%s:%d", app.Config.Server, app.Config.Port)
	app.OutLogger.Printf("Connecting client (version %s) to server on %s:%d ...", version, app.Config.Server, app.Config.Port)

	b, _ := ioutil.ReadFile("./certs/ca.cert")
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		app.ErrLogger.Printf("credentials: failed to append certificates")
		os.Exit(2)
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}
	cc, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	if err != nil {
		app.ErrLogger.Fatal(err)
	}
	defer cc.Close()

	city := binArgs[0]
	degrees := binArgs[1]

	client := ggwpb.NewGgwClient(cc)
	request := &ggwpb.GgwRequest{City: city, Degrees: degrees}

	resp, clErr := client.Ggw(context.Background(), request)
	if clErr != nil {
		app.ErrLogger.Fatal(clErr)
	}
	app.OutLogger.Printf("Temp in %s is %s (%s degrees)", resp.City, resp.Temp, resp.Degrees)
}
