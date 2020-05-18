package server

import "github.com/gorilla/mux"

type GLServer interface {
}

type Server struct {
	m *mux.Router
}
