package server

import "net/http"

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WIP"))
}

func (s *Server) GetProductsList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WIP"))
}

func (s *Server) GetproductDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WIP"))
}
