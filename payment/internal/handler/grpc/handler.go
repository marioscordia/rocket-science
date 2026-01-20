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
	// Implement your payment processing logic here
	transactionID := uuid.NewString()

	// For demonstration, returning a dummy response
	return &paymentv1.PayOrderResponse{
		TransactionUuid: transactionID,
	}, nil
}
