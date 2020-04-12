package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"github.com/Flukas88/newggw/proto/ggwpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Ggw(ctx context.Context, request *ggwpb.GgwRequest) (*ggwpb.GgwResponse, error) {
	city := request.City
	degrees := request.Degrees
	var ct CityTemp
	data, dataErr := getRespJSON(city)
	if dataErr != nil {
		return nil, dataErr
	}
	getErr := ct.getCityInfo(city, degrees, data)
	log.Printf("Responding for (%s,%s)", city, degrees)
	if getErr != nil {
		return nil, getErr
	}
	response := &ggwpb.GgwResponse{
		City:    ct.City,
		Temp:    ct.Temp,
		Degrees: ct.Degrees,
	}
	return response, nil
}

var version = "dev"

func main() {

	app := NewApp()

	address := fmt.Sprintf("%s:%d", app.Config.Host, app.Config.Port)

	// Server
	creds, err := credentials.NewServerTLSFromFile("./certs/service.pem", "./certs/service.key")
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server (version %s) is listening on %v ...\n", version, address)

	s := grpc.NewServer(grpc.Creds(creds))

	SetupCloseHandler(s)

	ggwpb.RegisterGgwServer(s, &server{})

	srvErr := s.Serve(lis)
	if srvErr != nil {
		log.Fatal(srvErr)
	}

}
