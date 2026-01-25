package grpc

import (
	"context"

	"github.com/google/uuid"
	paymentv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/payment/v1"
)

type Handler struct {
	paymentv1.UnimplementedPaymentServiceServer
}

func (h *Handler) PayOrder(ctx context.Context, req *paymentv1.PayOrderRequest) (*paymentv1.PayOrderResponse, error) {
	// TODO: Implement your payment processing logic here

	return &paymentv1.PayOrderResponse{
		Success:       true,
		Message:       "Payment processed successfully",
		TransactionId: uuid.NewString(), // You'll need to import "github.com/google/uuid"
	}, nil
}
