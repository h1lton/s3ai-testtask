package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"s3ai-testtask/internal/presentation/rest/handler"

	"net/http"
)

func NewRouter(h *handler.Handler) (http.Handler, error) {
	r := chi.NewRouter()

	r.Use(middleware.SetHeader("Content-Type", "application/json; charset=utf-8"))

	r.Post("/accounts", h.CreateAccount)
	r.Post("/accounts/{id}/deposit", h.Deposit)
	r.Post("/accounts/{id}/withdraw", h.Withdraw)
	r.Get("/accounts/{id}/balance", h.GetBalance)

	return r, nil
}
