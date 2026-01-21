package directory

import (
	"context"
	"fmt"

	useremailpb "nfxid/protos/gen/directory/user_email"
)

// UserEmailClient UserEmail 客户端
type UserEmailClient struct {
	client useremailpb.UserEmailServiceClient
}

// NewUserEmailClient 创建 UserEmail 客户端
func NewUserEmailClient(client useremailpb.UserEmailServiceClient) *UserEmailClient {
	return &UserEmailClient{client: client}
}

// GetUserEmailByID 根据ID获取用户邮箱
func (c *UserEmailClient) GetUserEmailByID(ctx context.Context, id string) (*useremailpb.UserEmail, error) {
	req := &useremailpb.GetUserEmailByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserEmailByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmail, nil
}

// GetUserEmailByEmail 根据邮箱地址获取用户邮箱
func (c *UserEmailClient) GetUserEmailByEmail(ctx context.Context, email string) (*useremailpb.UserEmail, error) {
	req := &useremailpb.GetUserEmailByEmailRequest{
		Email: email,
	}

	resp, err := c.client.GetUserEmailByEmail(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmail, nil
}

// GetUserEmailsByUserID 根据用户ID获取用户邮箱列表
func (c *UserEmailClient) GetUserEmailsByUserID(ctx context.Context, userID string) ([]*useremailpb.UserEmail, error) {
	req := &useremailpb.GetUserEmailsByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserEmailsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmails, nil
}