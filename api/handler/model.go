package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackkenney/gl-server/model"
)

// ModelHandler manages the endpoints for the model.
type ModelHandler struct {
	model model.Model
}

// ProvideModelHandler returns an instance of a ModelHandler.
func ProvideModelHandler(model model.Model) *ModelHandler {
	return &ModelHandler{
		model: model,
	}
}

// RegisterRoutes establishes the endpoints for the model.
func (h *ModelHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/predict", h.handlePredict).Methods(http.MethodPost)
	r.HandleFunc("/train", h.handleTrain).Methods(http.MethodPost)
}

func (h *ModelHandler) handlePredict(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *ModelHandler) handleTrain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
