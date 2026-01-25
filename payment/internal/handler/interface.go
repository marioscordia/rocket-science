package handler

import "context"

type UseCase interface {
	Pay(ctx context.Context, orderUUID string, amount float64, method string) error
}
