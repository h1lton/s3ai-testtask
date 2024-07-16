package application

import "s3ai-testtask/internal/domain/service"

type ATMService struct {
	bank *service.BankService
}

func NewATMService() *ATMService {
	return &ATMService{
		bank: &service.BankService{},
	}
}

func (A *ATMService) CreateAccount() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (A *ATMService) Deposit(id string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (A *ATMService) Withdraw(id string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (A *ATMService) GetBalance(id string) float64 {
	//TODO implement me
	panic("implement me")
}
