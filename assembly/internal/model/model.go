package model

type Order struct {
	EventID       string `json:"event_id"`
	OrderID       string `json:"order_id"`
	UserID        string `json:"user_id"`
	PaymentMethod string `json:"payment_method"`
	TransactionID string `json:"transaction_id"`
}
