package grpc

import (
	"context"

	userv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/user/v1"
	"github.com/marioscordia/rocket-science/user/internal/handler"
)

type Handler struct {
	userv1.UnimplementedUserServiceServer
	useCase handler.UseCase
}

func (h *Handler) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	// Implementation here
	sessionID, err := h.useCase.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &userv1.LoginResponse{SessionId: sessionID}, nil
}
func (h *Handler) WhoAmI(ctx context.Context, req *userv1.WhoAmIRequest) (*userv1.WhoAmIResponse, error) {
	// Implementation here
	return &userv1.WhoAmIResponse{}, nil
}
func (h *Handler) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	// Implementation here
	return &userv1.RegisterResponse{}, nil
}
func (h *Handler) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	// Implementation here
	return &userv1.GetUserResponse{}, nil
}
