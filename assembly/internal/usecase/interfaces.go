package usecase

import "context"

type Repository interface {
	BuildShip(ctx context.Context) error
}
