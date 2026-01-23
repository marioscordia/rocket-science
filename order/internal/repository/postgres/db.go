package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marioscordia/rocket-science/order/internal/dto"
)

type DB struct {
	con *pgxpool.Pool
}

func NewDB(url string) (*DB, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to pgx pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging pgx pool: %w", err)
	}

	return &DB{con: pool}, nil
}

func (db *DB) CreateOrder(ctx context.Context, order *dto.CreateOrder) (string, error) {
	_, err := db.con.Exec(ctx, "INSERT INTO orders (id, user_id, part_ids, price, status) VALUES ($1, $2, $3, $4, $5)",
		order.ID, order.UserID, order.PartIds, order.Price, order.Status)
	if err != nil {
		return "", err
	}
	return order.ID, nil
}

func (db *DB) GetOrderByID(ctx context.Context, id string) (*dto.Order, error) {
	order := &dto.Order{}
	row := db.con.QueryRow(ctx, "SELECT id, user_id, part_ids, price, transaction_id, payment_method, status, created_at, updated_at FROM orders WHERE id=$1", id)
	err := row.Scan(&order.ID, &order.UserID, &order.PartIDs, &order.Price, &order.TransactionID, &order.PaymentMethod, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (db *DB) UpdateOrderStatus(ctx context.Context, id string, status string) error {
	_, err := db.con.Exec(ctx, "UPDATE orders SET status=$1, updated_at=NOW() WHERE id=$2", status, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateOrderPayment(ctx context.Context, id string, transactionID string, paymentMethod string) error {
	_, err := db.con.Exec(ctx, "UPDATE orders SET transaction_id=$1, payment_method=$2, status=$3, updated_at=NOW() WHERE id=$4", transactionID, paymentMethod, "paid", id)
	if err != nil {
		return err
	}
	return nil
}
