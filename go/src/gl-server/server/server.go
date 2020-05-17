package server

import "github.com/gorilla/mux"

type Server interface {
	m *mux.Router
}
