package payment

import (
	"context"

	paymentv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
)

var PaymentMethodMap = map[string]paymentv1.PaymentMethod{
	"unspecified":    paymentv1.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED,
	"card":           paymentv1.PaymentMethod_PAYMENT_METHOD_CARD,
	"sbp":            paymentv1.PaymentMethod_PAYMENT_METHOD_SBP,
	"credit_card":    paymentv1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD,
	"investor_money": paymentv1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY,
}

type Service struct {
	client paymentv1.PaymentServiceClient
}

func NewService(conn grpc.ClientConnInterface) *Service {
	client := paymentv1.NewPaymentServiceClient(conn)

	return &Service{
		client: client,
	}
}

func (s *Service) ProcessPayment(ctx context.Context, orderID, userID string, paymentMethod string) (string, error) {
	resp, err := s.client.PayOrder(ctx, &paymentv1.PayOrderRequest{
		OrderUuid:     orderID,
		UserUuid:      userID,
		PaymentMethod: PaymentMethodMap[paymentMethod],
	})
	if err != nil {
		return "", err
	}

	return resp.TransactionId, nil
}
