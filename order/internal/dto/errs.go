package dto

import "errors"

var (
	// Order existence errors
	ErrOrderNotFound = errors.New("order not found")

	// Order state errors
	ErrOrderAlreadyPaid      = errors.New("order has already been paid")
	ErrOrderAlreadyCanceled  = errors.New("order has already been canceled")
	ErrOrderCannotBePaid     = errors.New("order cannot be paid in current state")
	ErrOrderCannotBeCanceled = errors.New("order cannot be canceled in current state")

	// Validation errors
	ErrInvalidOrderID       = errors.New("invalid order ID")
	ErrInvalidUserID        = errors.New("invalid user ID")
	ErrEmptyPartList        = errors.New("part list cannot be empty")
	ErrInvalidPartID        = errors.New("invalid part ID")
	ErrInvalidPaymentMethod = errors.New("invalid payment method")

	// Business logic errors
	ErrPartNotFound      = errors.New("one or more parts not found")
	ErrPartOutOfStock    = errors.New("one or more parts out of stock")
	ErrInsufficientStock = errors.New("insufficient stock for requested parts")

	// Payment errors
	ErrPaymentFailed      = errors.New("payment processing failed")
	ErrPaymentTimeout     = errors.New("payment processing timeout")
	ErrInvalidTransaction = errors.New("invalid transaction ID")

	// Repository errors
	ErrDatabaseConnection  = errors.New("database connection error")
	ErrDatabaseQuery       = errors.New("database query error")
	ErrDatabaseTransaction = errors.New("database transaction error")
)
