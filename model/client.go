package model

import "time"

// Client model of table client
type Client struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Client) HasID() bool { return c.ID > 0 }

func (c Client) Validate() error {
	// implement validation of fields for creation and update
	return nil
}

// Clients slice of Client
type Clients []Client

func (c Clients) IsEmpty() bool { return len(c) == 0 }
