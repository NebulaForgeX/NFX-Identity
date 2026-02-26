package handler

import (
	"context"

	userAvatarApp "nfxidentity/modules/directory/application/user_avatars"
	userAvatarAppCommands "nfxidentity/modules/directory/application/user_avatars/commands"
	"nfxidentity/modules/directory/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	useravatarpb "nfxidentity/protos/gen/directory/user_avatar"

	"github.com/google/uuid"
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
func (h *UserAvatarHandler) CreateOrUpdateUserAvatar(
	ctx context.Context,
	req *useravatarpb.CreateOrUpdateUserAvatarRequest,
) (*useravatarpb.CreateOrUpdateUserAvatarResponse, error) {
	// 解析用户ID和图片ID
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	// 创建命令
	cmd := userAvatarAppCommands.CreateOrUpdateUserAvatarCmd{
		UserID:  userID,
		ImageID: imageID,
	}

	// 调用应用服务创建或更新用户头像
	if err := h.userAvatarAppSvc.CreateOrUpdateUserAvatar(ctx, cmd); err != nil {
		return nil, err
	}

	// 获取创建/更新的用户头像
	userAvatarView, err := h.userAvatarAppSvc.GetUserAvatarByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 转换为 protobuf 响应
	userAvatar := mapper.UserAvatarROToProto(&userAvatarView)
	return &useravatarpb.CreateOrUpdateUserAvatarResponse{UserAvatar: userAvatar}, nil
}

// GetUserAvatarByUserID 根据用户ID获取用户头像
func (h *UserAvatarHandler) GetUserAvatarByUserID(
	ctx context.Context,
	req *useravatarpb.GetUserAvatarByUserIDRequest,
) (*useravatarpb.GetUserAvatarByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userAvatarView, err := h.userAvatarAppSvc.GetUserAvatarByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userAvatar := mapper.UserAvatarROToProto(&userAvatarView)
	return &useravatarpb.GetUserAvatarByUserIDResponse{UserAvatar: userAvatar}, nil
}

// GetUserAvatarByImageID 根据图片ID获取用户头像
func (h *UserAvatarHandler) GetUserAvatarByImageID(
	ctx context.Context,
	req *useravatarpb.GetUserAvatarByImageIDRequest,
) (*useravatarpb.GetUserAvatarByImageIDResponse, error) {
	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userAvatarView, err := h.userAvatarAppSvc.GetUserAvatarByImageID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	userAvatar := mapper.UserAvatarROToProto(&userAvatarView)
	return &useravatarpb.GetUserAvatarByImageIDResponse{UserAvatar: userAvatar}, nil
}

// DeleteUserAvatar 删除用户头像
func (h *UserAvatarHandler) DeleteUserAvatar(ctx context.Context, req *useravatarpb.DeleteUserAvatarRequest) (*useravatarpb.DeleteUserAvatarResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userAvatarAppCommands.DeleteUserAvatarCmd{UserID: userID}
	if err := h.userAvatarAppSvc.DeleteUserAvatar(ctx, cmd); err != nil {
		return nil, err
	}

	return &useravatarpb.DeleteUserAvatarResponse{Success: true}, nil
}
