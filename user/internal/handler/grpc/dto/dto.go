package dto

import (
	commonv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/common/v1"
	userv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/user/v1"
	"github.com/marioscordia/rocket-science/user/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserModel converts a UserRegistrationInfo proto to a model.User
func ToUserModel(req *userv1.RegisterRequest) *model.User {
	if req == nil || req.RegistrationInfo == nil {
		return nil
	}

	user := &model.User{
		Username: req.RegistrationInfo.Info.Username,
		Email:    req.RegistrationInfo.Info.Email,
		Password: req.RegistrationInfo.Password,
	}

	if len(req.RegistrationInfo.Info.NotificationMethods) > 0 {
		user.NotificationMethods = make([]model.NotificationMethod, 0, len(req.RegistrationInfo.Info.NotificationMethods))
		for _, nm := range req.RegistrationInfo.Info.NotificationMethods {
			user.NotificationMethods = append(user.NotificationMethods, model.NotificationMethod{
				ProviderName: nm.ProviderName,
				Target:       nm.Target,
			})
		}
	}

	return user
}

// ToUserProto converts a model.User to a common.v1.User proto
func ToUserProto(user *model.User) *commonv1.User {
	if user == nil {
		return nil
	}

	protoUser := &commonv1.User{
		Id: user.ID,
		Info: &commonv1.UserInfo{
			Username: user.Username,
			Email:    user.Email,
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	if len(user.NotificationMethods) > 0 {
		protoUser.Info.NotificationMethods = make([]*commonv1.NotificationMethod, 0, len(user.NotificationMethods))
		for _, nm := range user.NotificationMethods {
			protoUser.Info.NotificationMethods = append(protoUser.Info.NotificationMethods, &commonv1.NotificationMethod{
				ProviderName: nm.ProviderName,
				Target:       nm.Target,
			})
		}
	}

	return protoUser
}

// ToSessionProto converts a model.Session to a common.v1.Session proto
func ToSessionProto(session *model.Session) *commonv1.Session {
	if session == nil {
		return nil
	}

	return &commonv1.Session{
		Id:        session.ID,
		CreatedAt: timestamppb.New(session.CreatedAt),
		UpdatedAt: timestamppb.New(session.UpdatedAt),
		ExpiresAt: timestamppb.New(session.ExpiresAt),
	}
}
