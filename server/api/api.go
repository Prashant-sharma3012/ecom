package api

import (
	"github.com/Prashant-sharma3012/ecom/tree/main/server/server"
	"github.com/gorilla/mux"
)

type API struct {
}

func SetupRoutes(s *server.Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", s.Ping)
	r.HandleFunc("/products", s.GetProductsList)
	r.HandleFunc("/productDetail/:productId", s.GetproductDetail)

	return r
}
