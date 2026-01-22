package handler

import (
	"context"

	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userPreferenceAppCommands "nfxid/modules/directory/application/user_preferences/commands"
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

// CreateUserPreference 创建用户偏好
func (h *UserPreferenceHandler) CreateUserPreference(ctx context.Context, req *userpreferencepb.CreateUserPreferenceRequest) (*userpreferencepb.CreateUserPreferenceResponse, error) {
	// 解析用户ID
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	// 转换 protobuf Struct 到 map[string]interface{}
	var notifications map[string]interface{}
	if req.Notifications != nil {
		notifications = req.Notifications.AsMap()
	}

	var privacy map[string]interface{}
	if req.Privacy != nil {
		privacy = req.Privacy.AsMap()
	}

	var display map[string]interface{}
	if req.Display != nil {
		display = req.Display.AsMap()
	}

	var other map[string]interface{}
	if req.Other != nil {
		other = req.Other.AsMap()
	}

	// 创建命令
	cmd := userPreferenceAppCommands.CreateUserPreferenceCmd{
		UserID:        userID,
		Theme:         "",
		Language:      "",
		Timezone:      "",
		Notifications: notifications,
		Privacy:       privacy,
		Display:       display,
		Other:         other,
	}

	// 设置可选字段
	if req.Theme != nil {
		cmd.Theme = *req.Theme
	}
	if req.Language != nil {
		cmd.Language = *req.Language
	}
	if req.Timezone != nil {
		cmd.Timezone = *req.Timezone
	}

	// 调用应用服务创建用户偏好
	userPreferenceID, err := h.userPreferenceAppSvc.CreateUserPreference(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to create user preference: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create user preference: %v", err)
	}

	// 获取创建的用户偏好
	userPreferenceView, err := h.userPreferenceAppSvc.GetUserPreference(ctx, userPreferenceID)
	if err != nil {
		logx.S().Errorf("failed to get created user preference: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created user preference: %v", err)
	}

	// 转换为 protobuf 响应
	userPreference := mapper.UserPreferenceROToProto(&userPreferenceView)
	return &userpreferencepb.CreateUserPreferenceResponse{UserPreference: userPreference}, nil
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
