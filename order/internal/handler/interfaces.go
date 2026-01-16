package handler

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type UseCase interface {
	CreateOrder(ctx context.Context, userID string, partIDs []string) error
	GetOrderByID(ctx context.Context, id string) (*dto.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) error
	UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod string) error
}
