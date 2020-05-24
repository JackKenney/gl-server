package api

import (
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/jackkenney/gl-server/api/handler"
)

// Server defines the router and the configuration variables.
type Server struct {
	// router Router
	modelHandler *handler.ModelHandler
	wait         time.Duration
}

// ProvideServer returns the location of the server instance on the heap.
func ProvideServer(modelHandler *handler.ModelHandler) *Server {
	return &Server{
		modelHandler: modelHandler,
	}
}

// Run starts the server listening with the default parameters
func (s *Server) Run(port string) {
	log.Println("Listening on port", port)

	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	modelRoutes := router.PathPrefix("/api/model").Subrouter()
	s.modelHandler.RegisterRoutes(modelRoutes)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println(err)
	}
}
