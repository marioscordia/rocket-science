package kafka

import (
	"context"

	"github.com/marioscordia/rocket-science/assembly/internal/handler"
	"github.com/marioscordia/rocket-science/platform/pkg/kafka"
)

type service struct {
	consumer kafka.Consumer
	useCase  handler.UseCase
}

func NewService(consumer kafka.Consumer, useCase handler.UseCase) *service {
	return &service{
		consumer: consumer,
		useCase:  useCase,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	err := s.consumer.Consume(ctx, s.handleOrderPaidMsg)
	if err != nil {
		return err
	}

	return nil
}
