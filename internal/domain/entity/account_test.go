package entity

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= 0.0001
}

func TestAccountOperations(t *testing.T) {
	account := NewAccount("id")
	if account == nil {
		t.Error("Account should be created")
		return
	}

	balance := account.GetBalance()
	if balance != 0 {
		t.Fatalf("Account balance should be 0.0 but was %f", balance)
	}

	tt := []struct {
		amount  float64
		deposit bool
		balance float64
		err     error
	}{
		{amount: 10.2, deposit: true, balance: 10.2, err: nil},
		{amount: 0.0, deposit: false, balance: 10.2, err: ErrGreaterThanZero},
		{amount: 10.0, deposit: false, balance: 0.2, err: nil},
		{amount: 0.1, deposit: false, balance: 0.1, err: nil},
		{amount: 0.2, deposit: false, balance: 0.1, err: ErrInsufficientFunds},
		{amount: -0.1, deposit: false, balance: 0.1, err: ErrGreaterThanZero},
	}

	for _, tc := range tt {
		var err error

		if tc.deposit {
			err = account.Deposit(tc.amount)
		} else {
			err = account.Withdraw(tc.amount)
		}

		if err != tc.err {
			t.Fatalf("Account operation error should be %v but was %v", tc.err, err)
		}

		if !almostEqual(account.GetBalance(), tc.balance) {
			t.Fatalf("Account balance should be %f but was %f", tc.balance, account.GetBalance())
		}
	}
}
