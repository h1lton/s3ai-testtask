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
	return &AccountRepository{
		accounts: make(map[string]*entity.Account),
		mu:       sync.Mutex{},
	}
}

func (r *AccountRepository) GetAccount(id string) (interfaces.BankAccount, bool) {
	r.mu.Lock()
	account, ok := r.accounts[id]
	r.mu.Unlock()
	return account, ok
}

func (r *AccountRepository) CreateAccount() string {
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

	return id
}
