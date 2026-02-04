package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marioscordia/rocket-science/user/internal/model"
)

type store struct {
	pool *pgxpool.Pool
}

func (s *store) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	// Implement the logic to get a user by ID from the database
}

func (s *store) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	// Implement the logic to get a user by username from the database
}

func (s *store) CreateUser(ctx context.Context, user *model.User) (string, error) {
	// Implement the logic to create a new user in the database
}
