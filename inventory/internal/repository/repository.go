package repository

import (
	"context"

	"github.com/marioscordia/rocket-science/inventory/internal/dto"
)

type Repository struct {
	store Store
}

func NewRepository(store Store) *Repository {
	return &Repository{
		store: store,
	}
}

func (r *Repository) GetPartByID(ctx context.Context, partID string) (*dto.Part, error) {
	return r.store.GetPartByID(ctx, partID)
}

func (r *Repository) ListParts(ctx context.Context, filter *dto.PartsFilter) ([]*dto.Part, error) {
	return r.store.ListParts(ctx, filter)
}
