package directory

import (
	"context"
	"fmt"

	userimagepb "nfxid/protos/gen/directory/user_image"
)

// UserImageClient UserImage 客户端
type UserImageClient struct {
	client userimagepb.UserImageServiceClient
}

// NewUserImageClient 创建 UserImage 客户端
func NewUserImageClient(client userimagepb.UserImageServiceClient) *UserImageClient {
	return &UserImageClient{client: client}
}

// CreateUserImage 创建用户图片
func (c *UserImageClient) CreateUserImage(ctx context.Context, userID, imageID string, displayOrder int32) (*userimagepb.UserImage, error) {
	req := &userimagepb.CreateUserImageRequest{
		UserId:       userID,
		ImageId:      imageID,
		DisplayOrder: displayOrder,
	}

	resp, err := c.client.CreateUserImage(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserImage, nil
}

// GetUserImageByID 根据ID获取用户图片
func (c *UserImageClient) GetUserImageByID(ctx context.Context, id string) (*userimagepb.UserImage, error) {
	req := &userimagepb.GetUserImageByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserImageByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserImage, nil
}

// GetUserImagesByUserID 根据用户ID获取用户图片列表
func (c *UserImageClient) GetUserImagesByUserID(ctx context.Context, userID string) ([]*userimagepb.UserImage, error) {
	req := &userimagepb.GetUserImagesByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserImagesByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserImages, nil
}

// GetUserImagesByImageID 根据图片ID获取用户图片列表
func (c *UserImageClient) GetUserImagesByImageID(ctx context.Context, imageID string) ([]*userimagepb.UserImage, error) {
	req := &userimagepb.GetUserImagesByImageIDRequest{
		ImageId: imageID,
	}

	resp, err := c.client.GetUserImagesByImageID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserImages, nil
}

// GetCurrentUserImageByUserID 获取用户当前图片（display_order = 0）
func (c *UserImageClient) GetCurrentUserImageByUserID(ctx context.Context, userID string) (*userimagepb.UserImage, error) {
	req := &userimagepb.GetCurrentUserImageByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetCurrentUserImageByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	if resp.UserImage == nil {
		return nil, fmt.Errorf("current user image not found")
	}

	return resp.UserImage, nil
}

// UpdateUserImageDisplayOrder 更新用户图片显示顺序
func (c *UserImageClient) UpdateUserImageDisplayOrder(ctx context.Context, id string, displayOrder int32) (*userimagepb.UserImage, error) {
	req := &userimagepb.UpdateUserImageDisplayOrderRequest{
		Id:           id,
		DisplayOrder: displayOrder,
	}

	resp, err := c.client.UpdateUserImageDisplayOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserImage, nil
}

// UpdateUserImageImageID 更新用户图片ID
func (c *UserImageClient) UpdateUserImageImageID(ctx context.Context, id, imageID string) (*userimagepb.UserImage, error) {
	req := &userimagepb.UpdateUserImageImageIDRequest{
		Id:      id,
		ImageId: imageID,
	}

	resp, err := c.client.UpdateUserImageImageID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserImage, nil
}

// DeleteUserImage 删除用户图片
func (c *UserImageClient) DeleteUserImage(ctx context.Context, id string) error {
	req := &userimagepb.DeleteUserImageRequest{
		Id: id,
	}

	_, err := c.client.DeleteUserImage(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}

	return nil
}
