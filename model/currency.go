package model

import "time"

// Currency model of table currency
type Currency struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Currency) HasID() bool { return c.ID > 0 }

func (c Currency) Validate() error {
	// implement validation of fields for creation and update
	return nil
}

// Currencies slice of Currency
type Currencies []Currency

func (c Currencies) IsEmpty() bool { return len(c) == 0 }
