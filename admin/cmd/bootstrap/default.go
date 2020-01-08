package bootstrap

import (
	endpointv1 "github.com/digithun/shio/admin/pkg/endpoint/v1"
	"github.com/digithun/shio/admin/pkg/transport"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	serv *grpc.Server
}

func New() *App {

	adminEndpoint := endpointv1.NewAdminEndPoint()
	grpsServ := transport.New(adminEndpoint)

	return &App{serv:grpsServ}
}

func (a *App) Run(grpcPort string) error {
	tcp, err := net.Listen("tcp", grpcPort)
	if err != nil {
		panic(err)
	}
	return a.serv.Serve(tcp)
}

func (a *App) Stop()  {
	a.serv.Stop()
}

