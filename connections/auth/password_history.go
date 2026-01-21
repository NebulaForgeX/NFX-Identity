package auth

import (
	"context"
	"fmt"

	passwordhistorypb "nfxid/protos/gen/auth/password_history"
)

// PasswordHistoryClient PasswordHistory 客户端
type PasswordHistoryClient struct {
	client passwordhistorypb.PasswordHistoryServiceClient
}

// NewPasswordHistoryClient 创建 PasswordHistory 客户端
func NewPasswordHistoryClient(client passwordhistorypb.PasswordHistoryServiceClient) *PasswordHistoryClient {
	return &PasswordHistoryClient{client: client}
}

// GetPasswordHistoryByID 根据ID获取密码历史
func (c *PasswordHistoryClient) GetPasswordHistoryByID(ctx context.Context, id string) (*passwordhistorypb.PasswordHistory, error) {
	req := &passwordhistorypb.GetPasswordHistoryByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetPasswordHistoryByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.PasswordHistory, nil
}

// GetPasswordHistoriesByUserID 根据用户ID获取密码历史列表
func (c *PasswordHistoryClient) GetPasswordHistoriesByUserID(ctx context.Context, userID string, tenantID *string) ([]*passwordhistorypb.PasswordHistory, error) {
	req := &passwordhistorypb.GetPasswordHistoriesByUserIDRequest{
		UserId:   userID,
		TenantId: tenantID,
	}

	resp, err := c.client.GetPasswordHistoriesByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.PasswordHistories, nil
}