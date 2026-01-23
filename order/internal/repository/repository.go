package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type Repo struct {
	store Store
}

func NewRepo(store Store) *Repo {
	return &Repo{store: store}
}

func (r *Repo) CreateOrder(ctx context.Context, order *dto.CreateOrder) (string, error) {
	order.ID = uuid.NewString()
	order.Status = "pending"

	return r.store.CreateOrder(ctx, order)
}

func (r *Repo) GetOrderByID(ctx context.Context, id string) (*dto.Order, error) {
	return r.store.GetOrderByID(ctx, id)
}

func (r *Repo) UpdateOrderStatus(ctx context.Context, id string, status string) error {
	return r.store.UpdateOrderStatus(ctx, id, status)
}

func (r *Repo) UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod string) error {
	return r.store.UpdateOrderPayment(ctx, id, transactionID, paymentMethod)
}
