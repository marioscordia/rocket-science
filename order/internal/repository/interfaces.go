package repository

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/repository/model"
)

type Store interface {
	CreateOrder(ctx context.Context, order *model.CreateOrder) (string, error)
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) error
	UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod string) error
}
