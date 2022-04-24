package server

import (
	"flamingo-authService/auth/repository"
	"github.com/gorilla/mux"
	"transaction-service/config"
)

// Server struct
type Server struct {
	svc repository.DBOps
	cfg *config.Config
}

// NewServer constructor
func NewServer(cfg *config.Config, s repository.DBOps) *Server {
	return &Server{cfg: cfg, svc: s}
}

// Run server
func (s *Server) Run() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/accounts", s.createAccount)
	r.HandleFunc("/api/accounts/{account_id}", s.getAccount)
	r.HandleFunc("/api/transactions", s.createTransaction)

	return r
}
