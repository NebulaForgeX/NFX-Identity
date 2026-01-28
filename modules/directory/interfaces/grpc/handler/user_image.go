package handler

import (
	"context"

	userImageApp "nfxid/modules/directory/application/user_images"
	userImageAppCommands "nfxid/modules/directory/application/user_images/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userimagepb "nfxid/protos/gen/directory/user_image"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image_id: %v", err)
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
		logx.S().Errorf("failed to create user image: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create user image: %v", err)
	}

	// 获取创建的用户图片
	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		logx.S().Errorf("failed to get created user image: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created user image: %v", err)
	}

	// 转换为 protobuf 响应
	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.CreateUserImageResponse{UserImage: userImage}, nil
}

// GetUserImageByID 根据ID获取用户图片
func (h *UserImageHandler) GetUserImageByID(ctx context.Context, req *userimagepb.GetUserImageByIDRequest) (*userimagepb.GetUserImageByIDResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_image_id: %v", err)
	}

	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		logx.S().Errorf("failed to get user image by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user image not found: %v", err)
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.GetUserImageByIDResponse{UserImage: userImage}, nil
}

// GetUserImagesByUserID 根据用户ID获取用户图片列表
func (h *UserImageHandler) GetUserImagesByUserID(ctx context.Context, req *userimagepb.GetUserImagesByUserIDRequest) (*userimagepb.GetUserImagesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userImageViews, err := h.userImageAppSvc.GetUserImagesByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user images by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user images: %v", err)
	}

	userImages := mapper.UserImageListROToProto(userImageViews)
	return &userimagepb.GetUserImagesByUserIDResponse{UserImages: userImages}, nil
}

// GetUserImagesByImageID 根据图片ID获取用户图片列表
func (h *UserImageHandler) GetUserImagesByImageID(ctx context.Context, req *userimagepb.GetUserImagesByImageIDRequest) (*userimagepb.GetUserImagesByImageIDResponse, error) {
	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image_id: %v", err)
	}

	userImageViews, err := h.userImageAppSvc.GetUserImagesByImageID(ctx, imageID)
	if err != nil {
		logx.S().Errorf("failed to get user images by image_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user images: %v", err)
	}

	userImages := mapper.UserImageListROToProto(userImageViews)
	return &userimagepb.GetUserImagesByImageIDResponse{UserImages: userImages}, nil
}

// GetCurrentUserImageByUserID 获取用户当前图片（display_order = 0）
func (h *UserImageHandler) GetCurrentUserImageByUserID(ctx context.Context, req *userimagepb.GetCurrentUserImageByUserIDRequest) (*userimagepb.GetCurrentUserImageByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userImageView, err := h.userImageAppSvc.GetCurrentUserImageByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get current user image by user_id: %v", err)
		// Return empty response if not found
		return &userimagepb.GetCurrentUserImageByUserIDResponse{}, nil
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.GetCurrentUserImageByUserIDResponse{UserImage: userImage}, nil
}

// UpdateUserImageDisplayOrder 更新用户图片显示顺序
func (h *UserImageHandler) UpdateUserImageDisplayOrder(ctx context.Context, req *userimagepb.UpdateUserImageDisplayOrderRequest) (*userimagepb.UpdateUserImageDisplayOrderResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_image_id: %v", err)
	}

	cmd := userImageAppCommands.UpdateUserImageDisplayOrderCmd{
		UserImageID: userImageID,
		DisplayOrder: int(req.DisplayOrder),
	}

	if err := h.userImageAppSvc.UpdateUserImageDisplayOrder(ctx, cmd); err != nil {
		logx.S().Errorf("failed to update user image display order: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to update user image display order: %v", err)
	}

	// 获取更新的用户图片
	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		logx.S().Errorf("failed to get updated user image: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get updated user image: %v", err)
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.UpdateUserImageDisplayOrderResponse{UserImage: userImage}, nil
}

// UpdateUserImageImageID 更新用户图片ID
func (h *UserImageHandler) UpdateUserImageImageID(ctx context.Context, req *userimagepb.UpdateUserImageImageIDRequest) (*userimagepb.UpdateUserImageImageIDResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_image_id: %v", err)
	}

	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image_id: %v", err)
	}

	cmd := userImageAppCommands.UpdateUserImageImageIDCmd{
		UserImageID: userImageID,
		ImageID:     imageID,
	}

	if err := h.userImageAppSvc.UpdateUserImageImageID(ctx, cmd); err != nil {
		logx.S().Errorf("failed to update user image image_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to update user image image_id: %v", err)
	}

	// 获取更新的用户图片
	userImageView, err := h.userImageAppSvc.GetUserImage(ctx, userImageID)
	if err != nil {
		logx.S().Errorf("failed to get updated user image: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get updated user image: %v", err)
	}

	userImage := mapper.UserImageROToProto(&userImageView)
	return &userimagepb.UpdateUserImageImageIDResponse{UserImage: userImage}, nil
}

// DeleteUserImage 删除用户图片
func (h *UserImageHandler) DeleteUserImage(ctx context.Context, req *userimagepb.DeleteUserImageRequest) (*userimagepb.DeleteUserImageResponse, error) {
	userImageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_image_id: %v", err)
	}

	cmd := userImageAppCommands.DeleteUserImageCmd{UserImageID: userImageID}
	if err := h.userImageAppSvc.DeleteUserImage(ctx, cmd); err != nil {
		logx.S().Errorf("failed to delete user image: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete user image: %v", err)
	}

	return &userimagepb.DeleteUserImageResponse{Success: true}, nil
}
