package http

import (
	"context"
	"errors"

	"github.com/marioscordia/rocket-science/order/internal/dto"
	"github.com/marioscordia/rocket-science/order/internal/handler"
	order_v1 "github.com/marioscordia/rocket-science/shared/pkg/openapi/order/v1"
)

type Handler struct {
	order_v1.UnimplementedHandler
	useCase handler.UseCase
}

func (h *Handler) CancelOrder(ctx context.Context, params order_v1.CancelOrderParams) (order_v1.CancelOrderRes, error) {
	if err := h.useCase.CancelOrder(ctx, params.OrderUUID); err != nil {
		switch {
		case errors.Is(err, dto.ErrOrderNotFound):
			return &order_v1.CancelOrderNotFound{
				Message: "Order not found",
			}, nil
		case errors.Is(err, dto.ErrInvalidOrderID),
			errors.Is(err, dto.ErrOrderAlreadyCanceled),
			errors.Is(err, dto.ErrOrderCannotBeCanceled):
			return &order_v1.CancelOrderConflict{
				Message: err.Error(),
			}, nil
		default:
			// Return as generic error since we don't know the specific error type
			return &order_v1.CancelOrderConflict{
				Message: "Internal server error",
			}, nil
		}
	}
	return &order_v1.CancelOrderNoContent{}, nil
}

func (h *Handler) CreateOrder(ctx context.Context, req order_v1.CreateOrderRequest) (order_v1.CreateOrderRes, error) {
	orderID, totalPrice, err := h.useCase.CreateOrder(ctx, req.GetUserUUID(), req.GetPartUuids())
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrInvalidUserID), errors.Is(err, dto.ErrEmptyPartList):
			return &order_v1.CreateOrderBadRequest{
				Message: err.Error(),
			}, nil
		case errors.Is(err, dto.ErrPartNotFound):
			return &order_v1.CreateOrderNotFound{
				Message: "One or more parts not found",
			}, nil
		case errors.Is(err, dto.ErrPartOutOfStock), errors.Is(err, dto.ErrInsufficientStock):
			return &order_v1.CreateOrderBadRequest{
				Message: err.Error(),
			}, nil
		default:
			return &order_v1.CreateOrderBadRequest{
				Message: "Failed to create order",
			}, nil
		}
	}

	return &order_v1.CreateOrderResponse{
		OrderUUID:  orderID,
		TotalPrice: totalPrice,
	}, nil
}

func (h *Handler) GetOrder(ctx context.Context, params order_v1.GetOrderParams) (order_v1.GetOrderRes, error) {
	order, err := h.useCase.GetOrderByID(ctx, params.OrderUUID)
	if err != nil {
		// Since GetOrderRes interface likely only has OrderResponse and Error types
		// we return a generic error response
		return &order_v1.Error{
			Message: err.Error(),
		}, nil
	}

	// Map transaction ID and payment method to optional fields
	var transactionUUID order_v1.OptNilString
	if order.TransactionID != "" {
		transactionUUID.SetTo(order.TransactionID)
	} else {
		transactionUUID.SetToNull()
	}

	var paymentMethod order_v1.OptNilPaymentMethod
	if order.PaymentMethod != "" {
		paymentMethod.SetTo(order_v1.PaymentMethod(order.PaymentMethod))
	} else {
		paymentMethod.SetToNull()
	}

	return &order_v1.OrderResponse{
		OrderUUID:       order.ID,
		UserUUID:        order.UserID,
		PartUuids:       order.PartIDs,
		TotalPrice:      order.Price,
		TransactionUUID: transactionUUID,
		PaymentMethod:   paymentMethod,
		Status:          order_v1.OrderStatus(order.Status),
	}, nil
}

func (h *Handler) PayOrder(ctx context.Context, req order_v1.PayOrderRequest, params order_v1.PayOrderParams) (order_v1.PayOrderRes, error) {
	transactionID := ""

	transactionID, err := h.useCase.PayOrder(ctx, params.OrderUUID, transactionID, string(req.GetPaymentMethod()))
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrOrderNotFound):
			return &order_v1.PayOrderNotFound{
				Message: "Order not found",
			}, nil
		case errors.Is(err, dto.ErrInvalidOrderID),
			errors.Is(err, dto.ErrInvalidPaymentMethod):
			return &order_v1.PayOrderBadRequest{
				Message: err.Error(),
			}, nil
		default:
			return &order_v1.PayOrderBadRequest{
				Message: "Payment processing failed",
			}, nil
		}
	}

	return &order_v1.PayOrderResponse{
		TransactionUUID: transactionID,
	}, nil
}
