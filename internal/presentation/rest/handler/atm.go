package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"s3ai-testtask/internal/infrastructure/logger/sl"
)

// CreateAccount
// @Summary create account
// @Tags accounts
// @Accept json
// @Produce json
// @Success 201 {object} CreateAccountResponse
// @Failure 500 {object} ErrorResponse
// @Router /accounts [post]
func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := h.atmService.CreateAccount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating account", sl.Err(err))
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(CreateAccountResponse{ID: id})
}

// Deposit
// @Summary Deposit
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param amount body AmountRequest true "Amount"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /accounts/{id}/deposit [post]
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
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Withdraw
// @Summary Withdraw
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param amount body AmountRequest true "Amount"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /accounts/{id}/withdraw [post]
func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")

	var req AmountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	err = h.atmService.Withdraw(accountId, req.Amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetBalance
// @Summary Get balance
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} BalanceResponse
// @Failure 400 {object} ErrorResponse
// @Router /accounts/{id}/balance [get]
func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")
	balance, err := h.atmService.GetBalance(accountId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(BalanceResponse{Balance: balance})
}
