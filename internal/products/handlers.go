package products

import (
	"github.com/mohammedkhalf/Ecommerce-API/internal/json"
	"log"
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

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)
}
