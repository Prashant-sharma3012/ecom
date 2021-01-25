package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Addr           string
	Logger         *logrus.Logger
	DB             *db.DB
	Handler        *mux.Router
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func SetupServe() *Server {
	fmt.Println("setting up server")

	return &Server{
		Logger:         logrus.New(),
		DB:             db.GetDBClient(),
		Addr:           ":3000",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
}
