package application

import "s3ai-testtask/internal/domain/service"

type ATMService struct {
	bank *service.BankService
}

func NewATMService(bank *service.BankService) *ATMService {
	return &ATMService{
		bank: &service.BankService{},
	}
}

func (a *ATMService) CreateAccount() (string, error) {
	accountId, err := a.bank.CreateAccount()
	if err != nil {
		return "", err
	}

	return accountId, nil
}

func (a *ATMService) Deposit(id string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (a *ATMService) Withdraw(id string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (a *ATMService) GetBalance(id string) (float64, error) {
	//TODO implement me
	panic("implement me")
}
