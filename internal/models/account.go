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

type Account struct {
	ID      int64
	Balance float64
}

type AccountUpdate struct {
	ID    int64
	Delta float64
}

type ApiCallData struct {
	AccountID int64
	Operation OpType
}

type OpType string
