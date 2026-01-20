package usecase

import (
	"context"

	"github.com/marioscordia/rocket-science/inventory/internal/dto"
)

type Repo interface {
	GetPartByID(ctx context.Context, partID string) (*dto.Part, error)
	ListParts(ctx context.Context, filter *dto.PartsFilter) ([]*dto.Part, error)
}
