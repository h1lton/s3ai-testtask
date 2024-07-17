package service

import (
	"errors"
	"s3ai-testtask/internal/domain/interfaces"
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

func (s *BankService) CreateAccount() string {
	return s.repo.CreateAccount()
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
