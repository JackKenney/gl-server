package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackkenney/gl-server/api/handler"
)

// Router holds the injected model and handles the REST endpoints.
type Router http.Handler

// ProvideRouter returns a Server object.
func ProvideRouter(modelHandler *handler.ModelHandler) Router {
	router := mux.NewRouter()
	router = router.StrictSlash(false)

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	modelRoutes := router.PathPrefix("/api/model").Subrouter()
	modelHandler.RegisterRoutes(modelRoutes)

	return router
}
