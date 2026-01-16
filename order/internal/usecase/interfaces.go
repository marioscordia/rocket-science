package usecase

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type Repo interface {
	CreateOrder(ctx context.Context, order *dto.CreateOrder) (string, error)
	GetOrderByID(ctx context.Context, id string) (*dto.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) error
	UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod int32) error
}

type InventoryService interface {
	GetParts(ctx context.Context, partIDs []string) (map[string]*dto.Part, error)
}

type PaymentService interface {
	ProcessPayment(ctx context.Context, orderID, userID string, paymentMethod int32) (string, error)
}
