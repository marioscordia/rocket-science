package kafka

import (
	"context"

	"github.com/marioscordia/rocket-science/platform/pkg/kafka"
)

func (s *service) handleOrderPaidMsg(ctx context.Context, msg kafka.Message) error {
	return s.useCase.BuildShip(ctx)
}
