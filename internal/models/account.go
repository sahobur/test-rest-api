package models

const (
	Deposite   OpType = "deposite"
	Withdraw   OpType = "withdraw"
	GetBalance OpType = "getbalance"
	Create     OpType = "create"
)

type Amount struct {
	Amount float64 `json:"amount"`
}

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
	AccountID int64
	Operation OpType
}

type OpType string
