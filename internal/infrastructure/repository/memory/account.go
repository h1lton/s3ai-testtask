package memory

import (
	"github.com/google/uuid"
	"s3ai-testtask/internal/domain/entity"
	"s3ai-testtask/internal/domain/interfaces"
	"sync"
)

type AccountRepository struct {
	accounts map[string]*entity.Account
	mu       sync.Mutex
}

func NewAccountRepository() interfaces.AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) GetAccount(id string) (interfaces.BankAccount, bool) {
	//TODO implement me
	panic("implement me")
}

func (r *AccountRepository) CreateAccount() (string, error) {
	r.mu.Lock()
	var id string
	for {
		id = uuid.New().String()
		if _, ok := r.accounts[id]; !ok {
			break
		}
	}

	r.accounts[id] = entity.NewAccount(id)
	r.mu.Unlock()

	return id, nil
}

func (r *AccountRepository) SaveAccount(account interfaces.BankAccount) {
	//TODO implement me
	panic("implement me")
}
