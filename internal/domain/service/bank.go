package service

import (
	"errors"
	"s3ai-testtask/internal/domain/interfaces"
)

type BankService struct {
	repo interfaces.BankRepository
}

func NewBankService(repo interfaces.BankRepository) *BankService {
	return &BankService{
		repo: repo,
	}
}

func (s *BankService) CreateAccount() (string, error) {
	return "", errors.New("not implemented")
}

func (s *BankService) Deposit(id string, amount float64) error {
	return errors.New("not implemented")
}

func (s *BankService) Withdraw(id string, amount float64) error {
	return errors.New("not implemented")
}

func (s *BankService) GetBalance(id string) (float64, error) {
	return 0, errors.New("not implemented")
}
