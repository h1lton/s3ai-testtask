package application

import "s3ai-testtask/internal/domain/service"

type ATMService struct {
	bank *service.BankService
}

func NewATMService(bank *service.BankService) *ATMService {
	return &ATMService{
		bank: bank,
	}
}

// TODO: implement channels

func (a *ATMService) CreateAccount() string {
	return a.bank.CreateAccount()
}

func (a *ATMService) Deposit(accountId string, amount float64) error {
	return a.bank.Deposit(accountId, amount)
}

func (a *ATMService) Withdraw(accountId string, amount float64) error {
	return a.bank.Withdraw(accountId, amount)
}

func (a *ATMService) GetBalance(accountId string) (float64, error) {
	return a.bank.GetBalance(accountId)
}
