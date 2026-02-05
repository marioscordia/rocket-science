package grpc

import (
	"context"

	userv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/user/v1"
	"github.com/marioscordia/rocket-science/user/internal/handler"
	"github.com/marioscordia/rocket-science/user/internal/handler/grpc/dto"
)

type Handler struct {
	userv1.UnimplementedUserServiceServer
	useCase handler.UseCase
}

func NewHandler(useCase handler.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	sessionID, err := h.useCase.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &userv1.LoginResponse{SessionId: sessionID}, nil
}
func (h *Handler) WhoAmI(ctx context.Context, req *userv1.WhoAmIRequest) (*userv1.WhoAmIResponse, error) {
	user, session, err := h.useCase.WhoAmI(ctx, req.GetSessionId())
	if err != nil {
		return nil, err
	}
	return &userv1.WhoAmIResponse{
		User:    dto.ToUserProto(user),
		Session: dto.ToSessionProto(session),
	}, nil
}
func (h *Handler) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	user := dto.ToUserModel(req)
	userID, err := h.useCase.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userv1.RegisterResponse{UserId: userID}, nil
}
func (h *Handler) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := h.useCase.GetUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &userv1.GetUserResponse{User: dto.ToUserProto(user)}, nil
}
