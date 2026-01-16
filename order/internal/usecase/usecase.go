package usecase

import (
	"context"
	"errors"

	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type Usecase struct {
	repo             Repo
	inventoryService InventoryService
	paymentService   PaymentService
}

func (u *Usecase) CreateOrder(ctx context.Context, userID string, partIDs []string) (string, float64, error) {
	order := &dto.CreateOrder{
		UserID:  userID,
		PartIds: partIDs,
	}

	parts, err := u.inventoryService.GetParts(ctx, partIDs)
	if err != nil {
		return "", 0, err
	}

	for _, partID := range partIDs {
		part, ok := parts[partID]
		if !ok {
			return "", 0, errors.New("Part with ID " + partID + " does not exist.")
		}
		order.Price += part.Price
	}

	id, err := u.repo.CreateOrder(ctx, order)
	if err != nil {
		return "", 0, err
	}

	return id, order.Price, nil
}

func (u *Usecase) GetOrderByID(ctx context.Context, id string) (*dto.Order, error) {
	return u.repo.GetOrderByID(ctx, id)
}

func (u *Usecase) CancelOrder(ctx context.Context, id string) error {
	order, err := u.repo.GetOrderByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status == "paid" {
		return errors.New("Order has already been paid, can not cancel it.")
	}

	if order.Status == "canceled" {
		return errors.New("Order has already been canceled.")
	}

	if err := u.repo.UpdateOrderStatus(ctx, id, "canceled"); err != nil {
		return err
	}

	// TODO: refund logic

	return nil
}

func (u *Usecase) PayOrder(ctx context.Context, id string, transactionID string, paymentMethod int32) (string, error) {
	order, err := u.repo.GetOrderByID(ctx, id)
	if err != nil {
		return "", err
	}

	if order.Status == "paid" {
		return "", errors.New("Order has already been paid!")
	}

	if order.Status == "canceled" {
		return "", errors.New("Order has already been canceled, you can not pay for it.")
	}

	transactionID, err = u.paymentService.ProcessPayment(ctx, id, order.UserID, paymentMethod)
	if err != nil {
		return "", err
	}

	if err := u.repo.UpdateOrderPayment(ctx, id, transactionID, paymentMethod); err != nil {
		return "", err
	}

	return transactionID, nil
}
