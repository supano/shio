package main

import (
	"fmt"
	"github.com/digithun/shio/admin-api/cmd/bootstrap"
	"os"
)

func main()  {
	app := bootstrap.New()

	err := app.Run(":4444")
	if err != nil {
		fmt.Println("err", err)
		panic(err)
		return
	}

	os.Exit(0)
}
