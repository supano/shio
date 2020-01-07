package bootstrap

import "google.golang.org/grpc"

type App struct {
	serv *grpc.Server
}

func (a *App) New()  {
	grpsServ := transport

}

func (a *App) Run()  {
	
}

func (a *App) Stop()  {
	
}

