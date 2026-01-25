package usecase

import (
	"context"

	"github.com/marioscordia/rocket-science/inventory/internal/dto"
)

type UseCase struct {
	repo Repo
}

func NewUseCase(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) GetPartByID(ctx context.Context, partID string) (*dto.Part, error) {
	return u.repo.GetPartByID(ctx, partID)
}

func (u *UseCase) ListParts(ctx context.Context, filter *dto.PartsFilter) ([]*dto.Part, error) {
	return u.repo.ListParts(ctx, filter)
}
