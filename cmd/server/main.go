package main

import (
	"context"
	"fmt"
	"net"

	srv "github.com/Flukas88/newggw/pkg/models/server"
	"github.com/Flukas88/newggw/proto/ggwpb"
)

type server struct {
}

func (*server) Ggw(ctx context.Context, request *ggwpb.GgwRequest) (*ggwpb.GgwResponse, error) {
	city := request.City
	degrees := request.Degrees
	var ct srv.CityTemp
	data, dataErr := GetRespJSON(city)
	if dataErr != nil {
		return nil, dataErr
	}
	getErr := ct.GetCityInfo(city, degrees, data)
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
	app := NewApp("./certs/service.key", "./certs/service.pem")
	address := fmt.Sprintf("%s:%d", app.Config.Host, app.Config.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		app.ErrLogger.Fatalf("Error %v", err)
	}
	app.OutLogger.Printf("Server (version %s) is listening on %v ...\n", version, address)

	app.SetupCloseHandler()

	ggwpb.RegisterGgwServer(app.Server, &server{})

	srvErr := app.Server.Serve(lis)
	if srvErr != nil {
		app.ErrLogger.Fatal(srvErr)
	}

}
