package service

import (
	"errors"
	"log/slog"
	"s3ai-testtask/internal/domain/interfaces"
	"s3ai-testtask/internal/infrastructure/logger/sl"
)

var NotFoundError = errors.New("account not found")

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

func (s *BankService) Deposit(accountId string, amount float64) error {
	account, exist := s.repo.GetAccount(accountId)
	if !exist {
		return NotFoundError
	}

	return account.Deposit(amount)
}

func (s *BankService) Withdraw(accountId string, amount float64) error {
	account, exist := s.repo.GetAccount(accountId)
	if !exist {
		return NotFoundError
	}

	return account.Withdraw(amount)
}

func (s *BankService) GetBalance(accountId string) (float64, error) {
	account, exist := s.repo.GetAccount(accountId)
	if !exist {
		return 0, NotFoundError
	}

	return account.GetBalance(), nil
}
