package repository

import (
	"context"

	"github.com/marioscordia/rocket-science/inventory/internal/dto"
	"github.com/marioscordia/rocket-science/inventory/internal/repository/converter"
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
	part, err := r.store.GetPartByID(ctx, partID)
	if err != nil {
		return nil, err
	}
	return converter.ToDTO(part), nil
}

func (r *Repository) ListParts(ctx context.Context, filter *dto.PartsFilter) ([]*dto.Part, error) {
	parts, err := r.store.ListParts(ctx, converter.FilterToModel(filter))
	if err != nil {
		return nil, err
	}
	return converter.ToDTOList(parts), nil
}
