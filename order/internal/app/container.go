package app

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marioscordia/rocket-science/order/internal/handler"
	"github.com/marioscordia/rocket-science/order/internal/usecase"
)

type container struct {
	usecase handler.UseCase

	repo usecase.Repo

	paymentSvc   usecase.PaymentService
	inventorySvc usecase.InventoryService

	db *pgxpool.Pool
}

func NewContainer() *container {
	return &container{}
}

func (c *container) UseCase(ctx context.Context) handler.UseCase {
	if c.usecase == nil {
		// TODO: Initialize usecase properly
	}
	return c.usecase
}

func (c *container) Repository(ctx context.Context) usecase.Repo {
	if c.repo == nil {
		// TODO: Initialize repository properly
	}

	return c.repo
}

func (c *container) PaymentService(ctx context.Context) usecase.PaymentService {
	if c.paymentSvc == nil {
		// TODO: Initialize payment service properly
	}

	return c.paymentSvc
}

func (c *container) InventoryService(ctx context.Context) usecase.InventoryService {
	if c.inventorySvc == nil {
		// TODO: Initialize inventory service properly
	}

	return c.inventorySvc
}

func (c *container) DB(ctx context.Context) *pgxpool.Pool {
	if c.db == nil {
		// TODO: Initialize database connection properly
	}

	return c.db
}
