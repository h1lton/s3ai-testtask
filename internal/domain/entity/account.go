package entity

import (
	"errors"
	"sync"
)

var ErrLock = errors.New("another operation is already in progress for this account")
var ErrGreaterThanZero = errors.New("amount must be greater than zero")
var ErrInsufficientFunds = errors.New("insufficient funds")

type Account struct {
	id      string
	balance float64
	mu      sync.Mutex
}

func NewAccount(id string) *Account {
	return &Account{
		id:      id,
		balance: 0,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrGreaterThanZero
	}

	isLock := a.mu.TryLock()
	if !isLock {
		return ErrLock
	}
	a.balance += amount
	a.mu.Unlock()
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrGreaterThanZero
	}

	isLock := a.mu.TryLock()
	defer a.mu.Unlock()
	if !isLock {
		return ErrLock
	}

	if amount > a.balance {
		return ErrInsufficientFunds
	}

	a.balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	balance := a.balance
	a.mu.Unlock()
	return balance
}
