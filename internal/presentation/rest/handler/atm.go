package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"s3ai-testtask/internal/infrastructure/logger/sl"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := h.atmService.CreateAccount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating account", sl.Err(err))
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(CreateAccountResponse{ID: id})
}

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")

	var req AmountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	err = h.atmService.Deposit(accountId, req.Amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	}
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
}
