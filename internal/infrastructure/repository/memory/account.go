package memory

import "s3ai-testtask/internal/domain/interfaces"

type AccountRepository struct {
}

func NewAccountRepository() interfaces.AccountRepository {
	return &AccountRepository{}
}

func (b *AccountRepository) GetAccount(id string) (interfaces.BankAccount, bool) {
	//TODO implement me
	panic("implement me")
}

func (b *AccountRepository) CreateAccount() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b *AccountRepository) SaveAccount(account interfaces.BankAccount) {
	//TODO implement me
	panic("implement me")
}
