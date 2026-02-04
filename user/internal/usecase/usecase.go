package usecase

import (
	"context"

	"github.com/marioscordia/rocket-science/user/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type useCase struct {
	repo Repository
}

func (u *useCase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// TODO: Create session and return session ID
	return "", nil
}

func (u *useCase) WhoAmI(ctx context.Context, sessionID string) (*model.User, *model.Session, error) {
	session, err := u.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.repo.GetUserByID(ctx, session.UserID)
	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}

func (u *useCase) Register(ctx context.Context, user *model.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	userID, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (u *useCase) GetUser(ctx context.Context, userID string) (*model.User, error) {
	return u.repo.GetUserByID(ctx, userID)
}
