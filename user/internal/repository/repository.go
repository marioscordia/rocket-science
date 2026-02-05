package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/marioscordia/rocket-science/user/internal/constant"
	"github.com/marioscordia/rocket-science/user/internal/model"
)

type repository struct {
	store Store
	cache Cache
}

func NewRepository(store Store, cache Cache) *repository {
	return &repository{
		store: store,
		cache: cache,
	}
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

func (r *repository) CreateSession(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()
	session := &model.Session{
		ID:        sessionID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(constant.SessionTTL),
	}

	data, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	if err := r.cache.SetWithTTL(ctx, sessionID, data, constant.SessionTTL); err != nil {
		return "", err
	}

	return sessionID, nil
}
