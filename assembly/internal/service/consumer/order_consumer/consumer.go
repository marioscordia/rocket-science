package orderconsumer

import (
	"context"

	"github.com/marioscordia/rocket-science/platform/pkg/kafka"
	"github.com/marioscordia/rocket-science/platform/pkg/logger"
	"go.uber.org/zap"
)

type service struct {
	consumer kafka.Consumer
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting consumer")
	if err := s.consumer.Consume(ctx, s.OrderHandler); err != nil {
		logger.Error(ctx, "Error starting consumer: %v", zap.Error(err))
		return err
	}
	return nil
}
