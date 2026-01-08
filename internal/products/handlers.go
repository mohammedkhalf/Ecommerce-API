package products

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call the service -> list Product
	// 2. Return json response
	products := []string{"hello", "world"}
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}
