package assembly

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/dto"
	"github.com/marioscordia/rocket-science/platform/pkg/kafka"
	"github.com/marioscordia/rocket-science/platform/pkg/logger"
	eventsv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/events/v1"
	paymentv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/payment/v1"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	producer kafka.Producer
}

func NewService(producer kafka.Producer) *Service {
	return &Service{
		producer: producer,
	}
}

func (p *Service) ProduceUFORecorded(ctx context.Context, event dto.OrderPaidEvent) error {

	msg := &eventsv1.OrderPaid{
		EventId:       event.EventID,
		OrderId:       event.OrderID,
		UserId:        event.UserID,
		PaymentMethod: paymentv1.PaymentMethod(paymentv1.PaymentMethod_value[event.PaymentMethod]),
		TransactionId: event.TransactionID,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal UFORecorded", zap.Error(err))
		return err
	}

	err = p.producer.Send(ctx, []byte(event.EventID), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish UFORecorded", zap.Error(err))
		return err
	}

	return nil
}
