package handler

import (
	"context"

	userImageApp "nfxid/modules/directory/application/user_images"
	userImageAppCommands "nfxid/modules/directory/application/user_images/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	userimagepb "nfxid/protos/gen/directory/user_image"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
)

type UserImageHandler struct {
	userimagepb.UnimplementedUserImageServiceServer
	userImageAppSvc *userImageApp.Service
}

func NewUserImageHandler(userImageAppSvc *userImageApp.Service) *UserImageHandler {
	return &UserImageHandler{
		userImageAppSvc: userImageAppSvc,
	}
}

// CreateUserImage 创建用户图片
func (h *UserImageHandler) CreateUserImage(ctx context.Context, req *userimagepb.CreateUserImageRequest) (*userimagepb.CreateUserImageResponse, error) {
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
	cmd := userImageAppCommands.CreateUserImageCmd{
		UserID:       userID,
		ImageID:      imageID,
		DisplayOrder: int(req.DisplayOrder),
	}

	// 调用应用服务创建用户图片
	userImageID, err := h.userImageAppSvc.CreateUserImage(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// 获取创建的用户图片
	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		return nil, err
	}

	// 转换为 protobuf 响应
	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.CreateUserImageResponse{UserImage: userImage}, nil
}

// GetUserImageByID 根据ID获取用户图片
func (h *UserImageHandler) GetUserImageByID(ctx context.Context, req *userimagepb.GetUserImageByIDRequest) (*userimagepb.GetUserImageByIDResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		return nil, err
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.GetUserImageByIDResponse{UserImage: userImage}, nil
}

// GetUserImagesByUserID 根据用户ID获取用户图片列表
func (h *UserImageHandler) GetUserImagesByUserID(ctx context.Context, req *userimagepb.GetUserImagesByUserIDRequest) (*userimagepb.GetUserImagesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userImageViews, err := h.userImageAppSvc.GetUserImagesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userImages := mapper.UserImageListROToProto(userImageViews)
	return &userimagepb.GetUserImagesByUserIDResponse{UserImages: userImages}, nil
}

// GetUserImagesByImageID 根据图片ID获取用户图片列表
func (h *UserImageHandler) GetUserImagesByImageID(ctx context.Context, req *userimagepb.GetUserImagesByImageIDRequest) (*userimagepb.GetUserImagesByImageIDResponse, error) {
	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userImageViews, err := h.userImageAppSvc.GetUserImagesByImageID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	userImages := mapper.UserImageListROToProto(userImageViews)
	return &userimagepb.GetUserImagesByImageIDResponse{UserImages: userImages}, nil
}

// GetCurrentUserImageByUserID 获取用户当前图片（display_order = 0）
func (h *UserImageHandler) GetCurrentUserImageByUserID(ctx context.Context, req *userimagepb.GetCurrentUserImageByUserIDRequest) (*userimagepb.GetCurrentUserImageByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userImageView, err := h.userImageAppSvc.GetCurrentUserImageByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.GetCurrentUserImageByUserIDResponse{UserImage: userImage}, nil
}

// UpdateUserImageDisplayOrder 更新用户图片显示顺序
func (h *UserImageHandler) UpdateUserImageDisplayOrder(ctx context.Context, req *userimagepb.UpdateUserImageDisplayOrderRequest) (*userimagepb.UpdateUserImageDisplayOrderResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userImageAppCommands.UpdateUserImageDisplayOrderCmd{
		UserImageID: userImageID,
		DisplayOrder: int(req.DisplayOrder),
	}

	if err := h.userImageAppSvc.UpdateUserImageDisplayOrder(ctx, cmd); err != nil {
		return nil, err
	}

	// 获取更新的用户图片
	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		return nil, err
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.UpdateUserImageDisplayOrderResponse{UserImage: userImage}, nil
}

// UpdateUserImageImageID 更新用户图片ID
func (h *UserImageHandler) UpdateUserImageImageID(ctx context.Context, req *userimagepb.UpdateUserImageImageIDRequest) (*userimagepb.UpdateUserImageImageIDResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userImageAppCommands.UpdateUserImageImageIDCmd{
		UserImageID: userImageID,
		ImageID:     imageID,
	}

	if err := h.userImageAppSvc.UpdateUserImageImageID(ctx, cmd); err != nil {
		return nil, err
	}

	// 获取更新的用户图片
	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		return nil, err
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.UpdateUserImageImageIDResponse{UserImage: userImage}, nil
}

// DeleteUserImage 删除用户图片
func (h *UserImageHandler) DeleteUserImage(ctx context.Context, req *userimagepb.DeleteUserImageRequest) (*userimagepb.DeleteUserImageResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userImageAppCommands.DeleteUserImageCmd{UserImageID: userImageID}
	if err := h.userImageAppSvc.DeleteUserImage(ctx, cmd); err != nil {
		return nil, err
	}

	return &userimagepb.DeleteUserImageResponse{Success: true}, nil
}
