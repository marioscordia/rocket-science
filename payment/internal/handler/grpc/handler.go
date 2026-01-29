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
	// TODO: Call usecase to process payment

	transactionID := uuid.NewString()

	return &paymentv1.PayOrderResponse{
		TransactionId: transactionID,
	}, nil
}
