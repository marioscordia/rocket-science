package app

import (
	"context"

	"github.com/marioscordia/rocket-science/assembly/internal/handler"
	"github.com/marioscordia/rocket-science/assembly/internal/repository"
	"github.com/marioscordia/rocket-science/assembly/internal/usecase"
)

type container struct {
	useCase handler.UseCase
	repo    usecase.Repository
}

func (c *container) UseCase(ctx context.Context) handler.UseCase {
	if c.useCase == nil {
		c.useCase = usecase.NewUsecase(c.Repository(ctx))
	}
	return c.useCase
}

func (c *container) Repository(ctx context.Context) usecase.Repository {
	if c.repo == nil {
		c.repo = repository.NewRepository(ctx)
	}
	return c.repo
}
