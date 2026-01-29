package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/marioscordia/rocket-science/order/internal/dto"
	"github.com/marioscordia/rocket-science/order/internal/repository/converter"
	"github.com/marioscordia/rocket-science/order/internal/repository/postgres"
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

	return r.store.CreateOrder(ctx, converter.CreateOrderToModel(order))
}

func (r *Repo) GetOrderByID(ctx context.Context, id string) (*dto.Order, error) {
	order, err := r.store.GetOrderByID(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrOrderNotFound) {
			return nil, dto.ErrOrderNotFound
		}
		return nil, err
	}
	return converter.ToDTO(order), nil
}

func (r *Repo) UpdateOrderStatus(ctx context.Context, id string, status string) error {
	return r.store.UpdateOrderStatus(ctx, id, status)
}

func (r *Repo) UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod string) error {
	return r.store.UpdateOrderPayment(ctx, id, transactionID, paymentMethod)
}
