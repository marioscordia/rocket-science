package model

import "time"

// Order represents the database model for orders table
type Order struct {
	ID            string
	UserID        string
	PartIDs       []string
	Price         float64
	Status        string
	TransactionID *string
	PaymentMethod *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// CreateOrder represents the model for creating a new order
type CreateOrder struct {
	ID      string
	UserID  string
	PartIDs []string
	Price   float64
	Status  string
}
