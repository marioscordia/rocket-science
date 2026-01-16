package repository

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type Store interface {
	CreateOrder(ctx context.Context, order *dto.CreateOrder) error
	GetOrderByID(ctx context.Context, id string) (*dto.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) error
	UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod int32) error
}
