package auth

import (
	"context"
	"fmt"

	passwordresetpb "nfxid/protos/gen/auth/password_reset"
)

// PasswordResetClient PasswordReset 客户端
type PasswordResetClient struct {
	client passwordresetpb.PasswordResetServiceClient
}

// NewPasswordResetClient 创建 PasswordReset 客户端
func NewPasswordResetClient(client passwordresetpb.PasswordResetServiceClient) *PasswordResetClient {
	return &PasswordResetClient{client: client}
}

// GetPasswordResetByID 根据ID获取密码重置
func (c *PasswordResetClient) GetPasswordResetByID(ctx context.Context, id string) (*passwordresetpb.PasswordReset, error) {
	req := &passwordresetpb.GetPasswordResetByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetPasswordResetByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.PasswordReset, nil
}

// GetPasswordResetByResetID 根据重置ID获取密码重置
func (c *PasswordResetClient) GetPasswordResetByResetID(ctx context.Context, resetID string) (*passwordresetpb.PasswordReset, error) {
	req := &passwordresetpb.GetPasswordResetByResetIDRequest{
		ResetId: resetID,
	}

	resp, err := c.client.GetPasswordResetByResetID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.PasswordReset, nil
}

// GetPasswordResetsByUserID 根据用户ID获取密码重置列表
func (c *PasswordResetClient) GetPasswordResetsByUserID(ctx context.Context, userID string) ([]*passwordresetpb.PasswordReset, error) {
	req := &passwordresetpb.GetPasswordResetsByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetPasswordResetsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.PasswordResets, nil
}