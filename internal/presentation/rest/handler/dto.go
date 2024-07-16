package handler

type CreateAccountResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AmountRequest struct {
	Amount float64 `json:"amount"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}
