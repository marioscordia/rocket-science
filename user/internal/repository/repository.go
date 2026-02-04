package repository

import (
	"context"
	"encoding/json"

	"github.com/marioscordia/rocket-science/user/internal/model"
)

type repository struct {
	store Store
	cache Cache
}

func (r *repository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return r.store.GetUserByID(ctx, userID)
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return r.store.GetUserByUsername(ctx, username)
}

func (r *repository) CreateUser(ctx context.Context, user *model.User) (string, error) {
	return r.store.CreateUser(ctx, user)
}

func (r *repository) GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error) {
	data, err := r.cache.Get(ctx, sessionID)
	if err == nil {
		return nil, err
	}

	var session *model.Session
	if err := json.Unmarshal(data, session); err != nil {
		return nil, err
	}

	return session, nil
}
