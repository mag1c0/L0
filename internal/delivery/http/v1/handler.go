package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/mag1c0/L0/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api chi.Router) {
	api.Route("/api/v1", func(r chi.Router) {
		h.initOrdersRoutes(r)
	})
}
