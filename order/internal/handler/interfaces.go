package handler

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type UseCase interface {
	CreateOrder(ctx context.Context, userID string, partIDs []string) (string, float64, error)
	GetOrderByID(ctx context.Context, id string) (*dto.Order, error)
	PayOrder(ctx context.Context, id string, transactionID string, paymentMethod string) (string, error)
	CancelOrder(ctx context.Context, id string) error
}
