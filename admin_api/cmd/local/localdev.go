package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/digithun/shio/admin_api/cmd/bootstrap"
	"os"
)

func main()  {
	err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:5545")

	//projectId := os.Getenv("DATASTORE_PROJECT_ID")
	//if projectId == "" {
		projectId := "catcat-development"
	//}

	db, err := datastore.NewClient(context.Background(), projectId)
	if err != nil {
		panic(err)
		return
	}

	app := bootstrap.New(db)

	if err != nil {
		panic(err)
		return
	}

	err = app.Run(":4444")
	if err != nil {
		panic(err)
		return
	}

	os.Exit(0)
}
