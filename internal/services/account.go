package services

import (
	"awesomeProject/internal/models"
)

type AccountService struct {
	Acc models.Account
}

func NewAccountService(accID int64) AccountService {
	return AccountService{
		Acc: models.Account{
			ID: accID,
		},
	}
}

func (r AccountService) Deposite(amount float64) error {
	
	return nil
}

func (r AccountService) Withdraw(amount float64) error {
	return nil
}

func (r AccountService) GetBalance() float64 {
	return 0
}
