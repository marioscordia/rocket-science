package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marioscordia/rocket-science/order/internal/repository/model"
)

// ErrOrderNotFound is returned when an order is not found in the database
var ErrOrderNotFound = errors.New("order not found")

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) (*DB, error) {
	return &DB{pool: pool}, nil
}

func (db *DB) CreateOrder(ctx context.Context, order *model.CreateOrder) (string, error) {
	_, err := db.pool.Exec(ctx, "INSERT INTO orders (id, user_id, part_ids, price, status) VALUES ($1, $2, $3, $4, $5)",
		order.ID, order.UserID, order.PartIDs, order.Price, order.Status)
	if err != nil {
		return "", err
	}
	return order.ID, nil
}

func (db *DB) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	order := &model.Order{}
	row := db.pool.QueryRow(ctx, "SELECT id, user_id, part_ids, price, transaction_id, payment_method, status, created_at, updated_at FROM orders WHERE id=$1", id)
	err := row.Scan(&order.ID, &order.UserID, &order.PartIDs, &order.Price, &order.TransactionID, &order.PaymentMethod, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	return order, nil
}

func (db *DB) UpdateOrderStatus(ctx context.Context, id string, status string) error {
	_, err := db.pool.Exec(ctx, "UPDATE orders SET status=$1, updated_at=NOW() WHERE id=$2", status, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod string) error {
	_, err := db.pool.Exec(ctx, "UPDATE orders SET transaction_id=$1, payment_method=$2, status=$3, updated_at=NOW() WHERE id=$4", transactionID, paymentMethod, "paid", id)
	if err != nil {
		return err
	}
	return nil
}
