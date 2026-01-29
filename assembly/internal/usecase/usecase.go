package usecase

import "context"

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *usecase {
	return &usecase{repo: repo}
}

func (u *usecase) BuildShip(ctx context.Context) error {
	return u.repo.BuildShip(ctx)
}
