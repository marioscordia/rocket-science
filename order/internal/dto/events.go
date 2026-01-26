package dto

type OrderPaidEvent struct {
	EventID       string
	OrderID       string
	UserID        string
	PaymentMethod string
	TransactionID string
}
