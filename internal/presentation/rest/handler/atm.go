package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"s3ai-testtask/internal/infrastructure/logger/sl"
)

type CreateAccountResponse struct {
	ID string `json:"id"`
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := h.atmService.CreateAccount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating account", sl.Err(err))
	}

	w.WriteHeader(http.StatusCreated)

	resp := CreateAccountResponse{ID: id}
	json.NewEncoder(w).Encode(&resp)
}

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
}
