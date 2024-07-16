package service

import (
	"errors"
	"log/slog"
	"s3ai-testtask/internal/domain/interfaces"
	"s3ai-testtask/internal/infrastructure/logger/sl"
)

type BankService struct {
	repo interfaces.AccountRepository
}

func NewBankService(repo interfaces.AccountRepository) *BankService {
	return &BankService{
		repo: repo,
	}
}

func (s *BankService) CreateAccount() (string, error) {
	accountId, err := s.repo.CreateAccount()
	if err != nil {
		slog.Error("Error creating account in repository", sl.Err(err))
		return "", errors.New("account creation failed")
	}

	return accountId, nil
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
