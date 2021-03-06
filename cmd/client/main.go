package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/Flukas88/newggw/proto/ggwpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var version = "dev"

func main() {
	app := NewApp("./certs/ca.cert")

	city := flag.String("city", "Dublin", "The City")
	degrees := flag.String("degrees", "C", "C or F")

	flag.Parse()
	if flag.NFlag() < 2 {
		flag.Usage()
		os.Exit(2)
	}

	address := fmt.Sprintf("%s:%d", app.Config.Server, app.Config.Port)
	app.OutLogger.Printf("Connecting client (version %s) to server on %s:%d ...", version, app.Config.Server, app.Config.Port)

	config, configErr := app.SetCreds()
	if configErr != nil {
		app.ErrLogger.Fatalf(configErr.Error())
	}

	cc, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	if err != nil {
		app.ErrLogger.Fatal(err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			app.ErrLogger.Fatal(err)
		}
	}()

	client := ggwpb.NewGgwClient(cc)
	request := &ggwpb.GgwRequest{City: *city, Degrees: *degrees}

	resp, clErr := client.Ggw(context.Background(), request)
	if clErr != nil {
		app.ErrLogger.Fatal(clErr)
	}
	app.OutLogger.Printf("Temp in %s is %s (%s degrees)", resp.City, resp.Temp, resp.Degrees)
}
