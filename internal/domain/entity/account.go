package entity

import (
	"errors"
	"sync"
)

var LockError = errors.New("another operation is already in progress for this account")

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
	isLock := a.mu.TryLock()
	if !isLock {
		return LockError
	}
	a.balance += amount
	a.mu.Unlock()
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	isLock := a.mu.TryLock()
	if !isLock {
		return LockError
	}

	if amount > a.balance {
		return errors.New("insufficient funds")
	}

	a.balance -= amount
	a.mu.Unlock()
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func (a *Account) Id() string {
	return a.id
}

func (a *Account) SetId(id string) {
	a.id = id
}
