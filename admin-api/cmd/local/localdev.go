package main

import (
	"github.com/digithun/shio/admin-api/cmd/bootstrap"
	"os"
)

func main()  {
	app := bootstrap.New()

	err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:5545")
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
