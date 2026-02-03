package orderconsumer

import (
	"context"

	"github.com/marioscordia/rocket-science/assembly/internal/service/converter"
	"github.com/marioscordia/rocket-science/platform/pkg/kafka"
	"github.com/marioscordia/rocket-science/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) OrderHandler(ctx context.Context, msg kafka.Message) error {
	logger.Info(ctx, "Handling order message")
	order, err := converter.DecodeOrderMsg(msg.Value)
	if err != nil {
		logger.Error(ctx, "Error decoding order message", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Order message handled successfully",
		zap.String("event_id", order.EventID),
		zap.String("order_id", order.OrderID),
		zap.String("user_id", order.UserID),
		zap.String("payment_method", order.PaymentMethod),
		zap.String("transaction_id", order.TransactionID),
	)

	return nil
}
