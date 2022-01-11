package model

import "time"

// Transaction model of table transaction
type Transaction struct {
	ID            uint      `json:"id"`
	Money         float32   `json:"money"`
	Type          string    `json:"type"`
	AccountHolder uint      `json:"account_holder"`
	Subject       uint      `json:"subject"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (t Transaction) HasID() bool { return t.ID > 0 }

func (t Transaction) Validate() error {
	// implement validation of fields for creation and update
	return nil
}

// Transactions slice of Transaction
type Transactions []Transaction

func (t Transactions) IsEmpty() bool { return len(t) == 0 }

type TypeTransaction string

var TypeTransactionMap = map[TypeTransaction]string{
	Deposite: "DEPOSIT",
	Withdraw: "WITHDRAW",
	Transfer: "TRANSFER",
}

const (
	Deposite TypeTransaction = "DEPOSIT"
	Withdraw TypeTransaction = "WITHDRAW"
	Transfer TypeTransaction = "TRANSFER"
)

func IsValidTransactionType(t string) bool {
	if _, ok:= TypeTransactionMap[TypeTransaction(t)]; ok {
		return true
	}

	return false
}
