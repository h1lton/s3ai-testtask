package memory

import "s3ai-testtask/internal/domain/interfaces"

type AccountRepository struct {
}

func NewAccountRepository() interfaces.AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) GetAccount(id string) (interfaces.BankAccount, bool) {
	//TODO implement me
	panic("implement me")
}

func (r *AccountRepository) CreateAccount() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AccountRepository) SaveAccount(account interfaces.BankAccount) {
	//TODO implement me
	panic("implement me")
}
