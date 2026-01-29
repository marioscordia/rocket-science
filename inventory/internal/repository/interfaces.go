package repository

import (
	"context"

	"github.com/marioscordia/rocket-science/inventory/internal/repository/model"
)

type Store interface {
	GetPartByID(ctx context.Context, partID string) (*model.Part, error)
	ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error)
}
