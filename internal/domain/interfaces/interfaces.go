package interfaces

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type ATMService interface {
	Deposit(id string, amount float64) error
	Withdraw(id string, amount float64) error
	GetBalance(id string) float64
}
