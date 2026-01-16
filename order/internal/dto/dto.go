package dto

import "time"

type Order struct {
	ID            string
	UserID        string
	PartIDs       []string
	Price         float64
	Status        string
	TransactionID string
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreateOrder struct {
	ID      string
	UserID  string
	PartIds []string
	Price   float64
	Status  string
}

type Part struct {
	ID       string
	Price    float64
	Quantity int32
}
