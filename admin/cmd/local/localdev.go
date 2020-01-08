package main

import (
	"github.com/digithun/shio/admin/cmd/bootstrap"
	"os"
)

func main()  {
	app := bootstrap.New()

	err := app.Run(":4444")
	if err != nil {
		panic(err)
		return
	}

	os.Exit(0)
}
