package repository

import (
	"context"
	"time"

	"github.com/marioscordia/rocket-science/user/internal/model"
)

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	SetWithTTL(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	HashSet(ctx context.Context, key string, values any) error
	HGetAll(ctx context.Context, key string) ([]any, error)
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
	Operator
}

type Store interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
}

type Operator interface {
	SAdd(ctx context.Context, key, value string) error
	SRem(ctx context.Context, key, value string) error
	SIsMember(ctx context.Context, key, value string) (bool, error)
	SMembers(ctx context.Context, key string) ([]string, error)
}
