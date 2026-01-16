package payment

import (
	"context"

	paymentv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/payment/v1"
)

type Service struct {
	client paymentv1.PaymentServiceClient
}

func (s *Service) ProcessPayment(ctx context.Context, orderID, userID string, paymentMethod int32) (string, error) {
	resp, err := s.client.PayOrder(ctx, &paymentv1.PayOrderRequest{
		OrderUuid:     orderID,
		UserUuid:      userID,
		PaymentMethod: paymentv1.PaymentMethod(paymentMethod),
	})
	if err != nil {
		return "", err
	}

	return resp.TransactionUuid, nil
}
