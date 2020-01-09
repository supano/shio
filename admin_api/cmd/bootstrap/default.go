package bootstrap

import (
	"cloud.google.com/go/datastore"
	"errors"
	"github.com/digithun/shio/admin_api"
	endpointv1 "github.com/digithun/shio/admin_api/pkg/endpoint/v1"
	"github.com/digithun/shio/admin_api/pkg/repository"
	"github.com/digithun/shio/admin_api/pkg/transport"
	"github.com/digithun/shio/admin_api/pkg/usecase"
	"github.com/go-chi/chi"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
)

type App struct {
	serv *grpc.Server
}

func New(db *datastore.Client) *App {
	if db == nil {
		panic(errors.New(admin_api.ErrorDbIsNotProvided))
	}

	assetRepo := repository.NewDefaultAssetRepository(db)
	transactionRepo := repository.NewDefaultTransactionRepository(db)

	inventoryUseCase := usecase.NewDefaultInventoryUseCase(assetRepo)
	transactionUseCase := usecase.NewDefaultTransactionUseCase(transactionRepo)
	adminEndpoint := endpointv1.NewAdminEndPoint(inventoryUseCase, transactionUseCase)
	grpcServ := transport.New(adminEndpoint)

	return &App{serv:grpcServ}
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

