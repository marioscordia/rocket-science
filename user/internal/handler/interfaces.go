package handler

import (
	"context"

	"github.com/marioscordia/rocket-science/user/internal/model"
)

type UseCase interface {
	Login(ctx context.Context, username, password string) (string, error)
	WhoAmI(ctx context.Context, sessionID string) (*model.User, *model.Session, error)
	Register(ctx context.Context, user *model.User) (string, error)
	GetUser(ctx context.Context, userID string) (*model.User, error)
}
