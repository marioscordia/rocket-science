package usecase

import (
	"context"
	"fmt"

	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type Usecase struct {
	repo             Repo
	inventoryService InventoryService
	paymentService   PaymentService
}

func NewOrderUseCase(repo Repo, paymentSvc PaymentService, inventorySvc InventoryService, db any) *Usecase {
	return &Usecase{
		repo:             repo,
		paymentService:   paymentSvc,
		inventoryService: inventorySvc,
	}
}

func (u *Usecase) CreateOrder(ctx context.Context, userID string, partIDs []string) (string, float64, error) {
	if userID == "" {
		return "", 0, dto.ErrInvalidUserID
	}

	if len(partIDs) == 0 {
		return "", 0, dto.ErrEmptyPartList
	}

	order := &dto.CreateOrder{
		UserID:  userID,
		PartIds: partIDs,
	}

	parts, err := u.inventoryService.GetParts(ctx, partIDs)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get parts: %w", err)
	}

	for _, partID := range partIDs {
		part, ok := parts[partID]
		if !ok {
			return "", 0, fmt.Errorf("%w: %s", dto.ErrPartNotFound, partID)
		}
		order.Price += part.Price
	}

	id, err := u.repo.CreateOrder(ctx, order)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create order: %w", err)
	}

	return id, order.Price, nil
}

func (u *Usecase) GetOrderByID(ctx context.Context, id string) (*dto.Order, error) {
	if id == "" {
		return nil, dto.ErrInvalidOrderID
	}

	order, err := u.repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", dto.ErrOrderNotFound, id)
	}

	return order, nil
}

func (u *Usecase) CancelOrder(ctx context.Context, id string) error {
	if id == "" {
		return dto.ErrInvalidOrderID
	}

	order, err := u.repo.GetOrderByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %s", dto.ErrOrderNotFound, id)
	}

	if order.Status == "paid" {
		return dto.ErrOrderCannotBeCanceled
	}

	if order.Status == "canceled" {
		return dto.ErrOrderAlreadyCanceled
	}

	if err := u.repo.UpdateOrderStatus(ctx, id, "canceled"); err != nil {
		return fmt.Errorf("failed to cancel order: %w", err)
	}

	// TODO: refund logic

	return nil
}

func (u *Usecase) PayOrder(ctx context.Context, id string, transactionID string, paymentMethod int32) (string, error) {
	if id == "" {
		return "", dto.ErrInvalidOrderID
	}

	order, err := u.repo.GetOrderByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("%w: %s", dto.ErrOrderNotFound, id)
	}

	if order.Status == "paid" {
		return "", dto.ErrOrderAlreadyPaid
	}

	if order.Status == "canceled" {
		return "", dto.ErrOrderCannotBePaid
	}

	transactionID, err = u.paymentService.ProcessPayment(ctx, id, order.UserID, paymentMethod)
	if err != nil {
		return "", fmt.Errorf("%w: %v", dto.ErrPaymentFailed, err)
	}

	if err := u.repo.UpdateOrderPayment(ctx, id, transactionID, paymentMethod); err != nil {
		return "", fmt.Errorf("failed to update order payment: %w", err)
	}

	return transactionID, nil
}
