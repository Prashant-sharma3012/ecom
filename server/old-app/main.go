package main

import (
	"fmt"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/api"
	"github.com/Prashant-sharma3012/ecom/tree/main/server/server"
)

func main() {
	fmt.Println("starting server")

	s := server.SetupServe()
	s.Handler = api.SetupRoutes(s)
	s.Start()
}
