package handler

import "context"

type UseCase interface {
	BuildShip(ctx context.Context) error
}
