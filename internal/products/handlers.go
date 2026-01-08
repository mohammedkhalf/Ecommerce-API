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

	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products := struct {
		Products []string `json:"products"`
	}{}
	json.Write(w, http.StatusOk, products)
}
