package usecase

import (
	"context"

	"github.com/marioscordia/rocket-science/user/internal/model"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error)
}
