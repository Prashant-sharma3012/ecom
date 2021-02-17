package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/db"
	"github.com/Prashant-sharma3012/ecom/tree/main/server/interfaces"
	"github.com/gorilla/mux"
)

func main() {

	db := db.GetDBClient()

	log.Println("Setting up App")
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}).Methods("GET")

	interfaces.SetupProductRoutes(r, db.Product)

	s := &http.Server{
		Addr:           ":4000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening on port 4000")
	log.Fatal(s.ListenAndServe())

	log.Println("hey")
}
