package auth

import (
	"context"
	"fmt"

	loginattemptpb "nfxid/protos/gen/auth/login_attempt"
)

// LoginAttemptClient LoginAttempt 客户端
type LoginAttemptClient struct {
	client loginattemptpb.LoginAttemptServiceClient
}

// NewLoginAttemptClient 创建 LoginAttempt 客户端
func NewLoginAttemptClient(client loginattemptpb.LoginAttemptServiceClient) *LoginAttemptClient {
	return &LoginAttemptClient{client: client}
}

// GetLoginAttemptByID 根据ID获取登录尝试
func (c *LoginAttemptClient) GetLoginAttemptByID(ctx context.Context, id string) (*loginattemptpb.LoginAttempt, error) {
	req := &loginattemptpb.GetLoginAttemptByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetLoginAttemptByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.LoginAttempt, nil
}

// GetLoginAttemptsByUserID 根据用户ID获取登录尝试列表
func (c *LoginAttemptClient) GetLoginAttemptsByUserID(ctx context.Context, userID string, limit *int32) ([]*loginattemptpb.LoginAttempt, error) {
	req := &loginattemptpb.GetLoginAttemptsByUserIDRequest{
		UserId: userID,
		Limit:  limit,
	}

	resp, err := c.client.GetLoginAttemptsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.LoginAttempts, nil
}

// GetLoginAttemptsByIdentifier 根据标识符获取登录尝试列表
func (c *LoginAttemptClient) GetLoginAttemptsByIdentifier(ctx context.Context, identifier string, limit *int32) ([]*loginattemptpb.LoginAttempt, error) {
	req := &loginattemptpb.GetLoginAttemptsByIdentifierRequest{
		Identifier: identifier,
		Limit:      limit,
	}

	resp, err := c.client.GetLoginAttemptsByIdentifier(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.LoginAttempts, nil
}