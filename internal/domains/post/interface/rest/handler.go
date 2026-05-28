package rest

import (
	"encoding/json"
	"net/http"

	"github.com/homepage408/syncra/internal/domains/post/application/usecase"
	"github.com/homepage408/syncra/pkg/logger"
)

type Handler struct {
	service *usecase.Service
	log     logger.Logger
}

func New(service *usecase.Service, log logger.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "create-post endpoint"})
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "get-post endpoint"})
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "list-posts endpoint"})
}
