package converter

import (
	"github.com/marioscordia/rocket-science/order/internal/dto"
	"github.com/marioscordia/rocket-science/order/internal/repository/model"
)

// ToDTO converts a repository model Order to a dto.Order
func ToDTO(m *model.Order) *dto.Order {
	if m == nil {
		return nil
	}

	order := &dto.Order{
		ID:        m.ID,
		UserID:    m.UserID,
		PartIDs:   m.PartIDs,
		Price:     m.Price,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	if m.TransactionID != nil {
		order.TransactionID = *m.TransactionID
	}

	if m.PaymentMethod != nil {
		order.PaymentMethod = *m.PaymentMethod
	}

	return order
}

// FromDTO converts a dto.Order to a repository model Order
func FromDTO(d *dto.Order) *model.Order {
	if d == nil {
		return nil
	}

	order := &model.Order{
		ID:        d.ID,
		UserID:    d.UserID,
		PartIDs:   d.PartIDs,
		Price:     d.Price,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	if d.TransactionID != "" {
		order.TransactionID = &d.TransactionID
	}

	if d.PaymentMethod != "" {
		order.PaymentMethod = &d.PaymentMethod
	}

	return order
}

// CreateOrderToModel converts a dto.CreateOrder to a repository model CreateOrder
func CreateOrderToModel(d *dto.CreateOrder) *model.CreateOrder {
	if d == nil {
		return nil
	}

	return &model.CreateOrder{
		ID:      d.ID,
		UserID:  d.UserID,
		PartIDs: d.PartIds,
		Price:   d.Price,
		Status:  d.Status,
	}
}

// CreateOrderFromModel converts a repository model CreateOrder to a dto.CreateOrder
func CreateOrderFromModel(m *model.CreateOrder) *dto.CreateOrder {
	if m == nil {
		return nil
	}

	return &dto.CreateOrder{
		ID:      m.ID,
		UserID:  m.UserID,
		PartIds: m.PartIDs,
		Price:   m.Price,
		Status:  m.Status,
	}
}
