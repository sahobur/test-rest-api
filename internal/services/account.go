package services

import (
	"awesomeProject/internal/models"
)

type Account struct {
	Acc models.Account
	Ch  chan models.AccountUpdate
}

type BankAccount interface {
	Deposite(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

func NewAccountService(accID int64, ch chan models.AccountUpdate) Account {
	return Account{
		Acc: models.Account{
			ID: accID,
		},
		Ch: ch,
	}
}

func (r Account) Deposite(amount float64) error {
	r.Ch <- models.AccountUpdate{
		ID:    r.Acc.ID,
		Delta: amount,
	}
	return nil
}

func (r Account) Withdraw(amount float64) error {
	r.Ch <- models.AccountUpdate{
		ID:    r.Acc.ID,
		Delta: -amount,
	}
	return nil
}

func (r Account) GetBalance() float64 {
	return r.Acc.Balance
}
