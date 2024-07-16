package interfaces

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type ATMService interface {
	CreateAccount() (string, error)
	Deposit(id string, amount float64) error
	Withdraw(id string, amount float64) error
	GetBalance(id string) (float64, error)
}

type AccountRepository interface {
	GetAccount(id string) (BankAccount, bool)
	CreateAccount() (BankAccount, error)
	SaveAccount(account BankAccount)
}
