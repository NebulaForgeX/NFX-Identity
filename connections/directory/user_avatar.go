package directory

import (
	"context"
	"fmt"

	useravatarpb "nfxid/protos/gen/directory/user_avatar"
)

// UserAvatarClient UserAvatar 客户端
type UserAvatarClient struct {
	client useravatarpb.UserAvatarServiceClient
}

// NewUserAvatarClient 创建 UserAvatar 客户端
func NewUserAvatarClient(client useravatarpb.UserAvatarServiceClient) *UserAvatarClient {
	return &UserAvatarClient{client: client}
}

// CreateOrUpdateUserAvatar 创建或更新用户头像
func (c *UserAvatarClient) CreateOrUpdateUserAvatar(ctx context.Context, userID, imageID string) (*useravatarpb.UserAvatar, error) {
	req := &useravatarpb.CreateOrUpdateUserAvatarRequest{
		UserId:  userID,
		ImageId: imageID,
	}

	resp, err := c.client.CreateOrUpdateUserAvatar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserAvatar, nil
}

// GetUserAvatarByUserID 根据用户ID获取用户头像
func (c *UserAvatarClient) GetUserAvatarByUserID(ctx context.Context, userID string) (*useravatarpb.UserAvatar, error) {
	req := &useravatarpb.GetUserAvatarByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserAvatarByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	if resp.UserAvatar == nil {
		return nil, fmt.Errorf("user avatar not found")
	}

	return resp.UserAvatar, nil
}

// GetUserAvatarByImageID 根据图片ID获取用户头像
func (c *UserAvatarClient) GetUserAvatarByImageID(ctx context.Context, imageID string) (*useravatarpb.UserAvatar, error) {
	req := &useravatarpb.GetUserAvatarByImageIDRequest{
		ImageId: imageID,
	}

	resp, err := c.client.GetUserAvatarByImageID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	if resp.UserAvatar == nil {
		return nil, fmt.Errorf("user avatar not found")
	}

	return resp.UserAvatar, nil
}

// DeleteUserAvatar 删除用户头像
func (c *UserAvatarClient) DeleteUserAvatar(ctx context.Context, userID string) error {
	req := &useravatarpb.DeleteUserAvatarRequest{
		UserId: userID,
	}

	_, err := c.client.DeleteUserAvatar(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}

	return nil
}
