package models

const (
	Deposite   OpType = "deposite"
	Withdraw   OpType = "withdraw"
	GetBalance OpType = "getbalance"
	Create     OpType = "create"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int64
	Balance float64
}

type ApiCallData struct {
	accountID int64
	operation OpType
}

type OpType string
