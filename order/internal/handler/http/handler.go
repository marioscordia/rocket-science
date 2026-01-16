package http

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/handler"
	order_v1 "github.com/marioscordia/rocket-science/shared/pkg/openapi/order/v1"
)

type Handler struct {
	useCase handler.UseCase
}

func (h *Handler) CancelOrder(ctx context.Context, params order_v1.CancelOrderParams) (order_v1.CancelOrderRes, error) {
	if err := h.useCase.UpdateOrderStatus(ctx, params.OrderUUID, "canceled"); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Handler) CreateOrder(ctx context.Context, params order_v1.CreateOrderRequest) (order_v1.CreateOrderRes, error) {
	if err := h.useCase.CreateOrder(ctx, params.GetUserUUID(), params.GetPartUuids()); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Handler) GetOrder(ctx context.Context, params order_v1.GetOrderParams) (order_v1.GetOrderRes, error) {
	order, err := h.useCase.GetOrderByID(ctx, params.OrderUUID)
	if err != nil {
		return nil, err
	}

	_ = order // TODO: map to response

	return nil, nil
}

func (h *Handler) PayOrder(ctx context.Context, params order_v1.PayOrderParams) (order_v1.PayOrderRes, error) {
	return nil, nil
}
