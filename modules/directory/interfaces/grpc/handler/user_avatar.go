package handler

import (
	"context"

	userAvatarApp "nfxid/modules/directory/application/user_avatars"
	userAvatarAppCommands "nfxid/modules/directory/application/user_avatars/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	useravatarpb "nfxid/protos/gen/directory/user_avatar"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserAvatarHandler struct {
	useravatarpb.UnimplementedUserAvatarServiceServer
	userAvatarAppSvc *userAvatarApp.Service
}

func NewUserAvatarHandler(userAvatarAppSvc *userAvatarApp.Service) *UserAvatarHandler {
	return &UserAvatarHandler{
		userAvatarAppSvc: userAvatarAppSvc,
	}
}

// CreateOrUpdateUserAvatar 创建或更新用户头像
func (h *UserAvatarHandler) CreateOrUpdateUserAvatar(ctx context.Context, req *useravatarpb.CreateOrUpdateUserAvatarRequest) (*useravatarpb.CreateOrUpdateUserAvatarResponse, error) {
	// 解析用户ID和图片ID
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image_id: %v", err)
	}

	// 创建命令
	cmd := userAvatarAppCommands.CreateOrUpdateUserAvatarCmd{
		UserID:  userID,
		ImageID: imageID,
	}

	// 调用应用服务创建或更新用户头像
	if err := h.userAvatarAppSvc.CreateOrUpdateUserAvatar(ctx, cmd); err != nil {
		logx.S().Errorf("failed to create or update user avatar: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create or update user avatar: %v", err)
	}

	// 获取创建/更新的用户头像
	userAvatarView, err := h.userAvatarAppSvc.GetUserAvatarByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user avatar: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user avatar: %v", err)
	}

	// 转换为 protobuf 响应
	userAvatar := mapper.UserAvatarROToProto(&userAvatarView)
	return &useravatarpb.CreateOrUpdateUserAvatarResponse{UserAvatar: userAvatar}, nil
}

// GetUserAvatarByUserID 根据用户ID获取用户头像
func (h *UserAvatarHandler) GetUserAvatarByUserID(ctx context.Context, req *useravatarpb.GetUserAvatarByUserIDRequest) (*useravatarpb.GetUserAvatarByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userAvatarView, err := h.userAvatarAppSvc.GetUserAvatarByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user avatar by user_id: %v", err)
		// Return empty response if not found
		return &useravatarpb.GetUserAvatarByUserIDResponse{}, nil
	}

	userAvatar := mapper.UserAvatarROToProto(&userAvatarView)
	return &useravatarpb.GetUserAvatarByUserIDResponse{UserAvatar: userAvatar}, nil
}

// GetUserAvatarByImageID 根据图片ID获取用户头像
func (h *UserAvatarHandler) GetUserAvatarByImageID(ctx context.Context, req *useravatarpb.GetUserAvatarByImageIDRequest) (*useravatarpb.GetUserAvatarByImageIDResponse, error) {
	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image_id: %v", err)
	}

	userAvatarView, err := h.userAvatarAppSvc.GetUserAvatarByImageID(ctx, imageID)
	if err != nil {
		logx.S().Errorf("failed to get user avatar by image_id: %v", err)
		// Return empty response if not found
		return &useravatarpb.GetUserAvatarByImageIDResponse{}, nil
	}

	userAvatar := mapper.UserAvatarROToProto(&userAvatarView)
	return &useravatarpb.GetUserAvatarByImageIDResponse{UserAvatar: userAvatar}, nil
}

// DeleteUserAvatar 删除用户头像
func (h *UserAvatarHandler) DeleteUserAvatar(ctx context.Context, req *useravatarpb.DeleteUserAvatarRequest) (*useravatarpb.DeleteUserAvatarResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	cmd := userAvatarAppCommands.DeleteUserAvatarCmd{UserID: userID}
	if err := h.userAvatarAppSvc.DeleteUserAvatar(ctx, cmd); err != nil {
		logx.S().Errorf("failed to delete user avatar: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete user avatar: %v", err)
	}

	return &useravatarpb.DeleteUserAvatarResponse{Success: true}, nil
}
