package handler

import (
	"context"

	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserPreferenceHandler struct {
	userpreferencepb.UnimplementedUserPreferenceServiceServer
	userPreferenceAppSvc *userPreferenceApp.Service
}

func NewUserPreferenceHandler(userPreferenceAppSvc *userPreferenceApp.Service) *UserPreferenceHandler {
	return &UserPreferenceHandler{
		userPreferenceAppSvc: userPreferenceAppSvc,
	}
}

// GetUserPreferenceByID 根据ID获取用户偏好
func (h *UserPreferenceHandler) GetUserPreferenceByID(ctx context.Context, req *userpreferencepb.GetUserPreferenceByIDRequest) (*userpreferencepb.GetUserPreferenceByIDResponse, error) {
	userPreferenceID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_preference_id: %v", err)
	}

	userPreferenceView, err := h.userPreferenceAppSvc.GetUserPreference(ctx, userPreferenceID)
	if err != nil {
		logx.S().Errorf("failed to get user preference by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user preference not found: %v", err)
	}

	userPreference := mapper.UserPreferenceROToProto(&userPreferenceView)
	return &userpreferencepb.GetUserPreferenceByIDResponse{UserPreference: userPreference}, nil
}

// GetUserPreferenceByUserID 根据用户ID获取用户偏好
func (h *UserPreferenceHandler) GetUserPreferenceByUserID(ctx context.Context, req *userpreferencepb.GetUserPreferenceByUserIDRequest) (*userpreferencepb.GetUserPreferenceByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userPreferenceView, err := h.userPreferenceAppSvc.GetUserPreferenceByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user preference by user_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user preference not found: %v", err)
	}

	userPreference := mapper.UserPreferenceROToProto(&userPreferenceView)
	return &userpreferencepb.GetUserPreferenceByUserIDResponse{UserPreference: userPreference}, nil
}
