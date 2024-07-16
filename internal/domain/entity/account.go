package entity

import (
	"errors"
	"math"
	"sync"
)

var ErrLock = errors.New("another operation is already in progress for this account")
var ErrGreaterThanZero = errors.New("amount must be greater than zero")
var ErrInsufficientFunds = errors.New("insufficient funds")

type Account struct {
	id      string
	balance int64
	mu      sync.Mutex
}

func NewAccount(id string) *Account {
	return &Account{
		id:      id,
		balance: 0,
	}
}

func (a *Account) floatToCents(f float64) int64 {
	multiplied := f * 1000
	rounded := math.Round(multiplied)
	return int64(rounded)
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrGreaterThanZero
	}

	amountInCents := a.floatToCents(amount)

	isLock := a.mu.TryLock()
	if !isLock {
		return ErrLock
	}
	a.balance += amountInCents
	a.mu.Unlock()
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrGreaterThanZero
	}

	amountInCents := a.floatToCents(amount)

	isLock := a.mu.TryLock()
	defer a.mu.Unlock()
	if !isLock {
		return ErrLock
	}

	if amountInCents > a.balance {
		return ErrInsufficientFunds
	}

	a.balance -= amountInCents
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	balance := a.balance
	a.mu.Unlock()
	return float64(balance) / 1000
}
