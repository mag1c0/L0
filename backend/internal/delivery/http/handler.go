package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/mag1c0/L0/backend/internal/delivery/http/v1"
	"github.com/mag1c0/L0/backend/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *chi.Mux {
	// Init chi handler
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	// Init router
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router chi.Router) {
	handlerV1 := v1.NewHandler(h.services)
	handlerV1.Init(router)
}
