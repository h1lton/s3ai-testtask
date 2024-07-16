package entity

import "s3ai-testtask/internal/domain/interfaces"

type Account struct {
	id      string
	balance float64
}

func NewAccount(id string, balance float64) interfaces.BankAccount {
	return &Account{
		id:      id,
		balance: balance,
	}
}

func (a *Account) Deposit(amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (a *Account) Withdraw(amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (a *Account) GetBalance() float64 {
	//TODO implement me
	panic("implement me")
}

func (a *Account) Id() string {
	return a.id
}

func (a *Account) SetId(id string) {
	a.id = id
}
