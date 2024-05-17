package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) initOrdersRoutes(api chi.Router) {
	api.Route("/orders", func(r chi.Router) {
		r.Get("/", h.GetOrders)
	})
	api.Get("/{id}/", h.GetOrderByID)
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.services.Orders.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	// search params
	id := chi.URLParam(r, "id")

	order, err := h.services.Orders.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
