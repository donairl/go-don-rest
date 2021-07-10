package main

import (
	"github.com/donairl/go-don-rest/lib"
	"github.com/donairl/go-don-rest/router"
)

func main() {
	// Echo instance

	db := lib.ConnectDb()
	e := router.New(db)
	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
