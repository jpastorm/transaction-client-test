package model

import "time"

// Account model of table account
type Account struct {
	ID         uint      `json:"id"`
	ClientID   uint      `json:"client_id"`
	CurrencyID uint      `json:"currency_id"`
	Money      float32   `json:"money"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (a Account) HasID() bool { return a.ID > 0 }

func (a Account) Validate() error {
	// implement validation of fields for creation and update
	return nil
}

// Accounts slice of Account
type Accounts []Account

func (a Accounts) IsEmpty() bool { return len(a) == 0 }
