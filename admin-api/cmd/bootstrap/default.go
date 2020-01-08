package bootstrap

import (
	endpointv1 "github.com/digithun/shio/admin-api/pkg/endpoint/v1"
	"github.com/digithun/shio/admin-api/pkg/transport"
	"github.com/go-chi/chi"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
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

	r := chi.NewRouter()
	r.Handle("/*", http.StripPrefix("/", grpcweb.WrapServer(a.serv)))
	handler := cors.AllowAll().Handler(r)

	return http.ListenAndServe(grpcPort,handler)
}

func (a *App) Stop()  {
	a.serv.Stop()
}

