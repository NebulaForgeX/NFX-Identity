package auth

import (
	"context"
	"fmt"

	accountlockoutpb "nfxid/protos/gen/auth/account_lockout"
)

// AccountLockoutClient AccountLockout 客户端
type AccountLockoutClient struct {
	client accountlockoutpb.AccountLockoutServiceClient
}

// NewAccountLockoutClient 创建 AccountLockout 客户端
func NewAccountLockoutClient(client accountlockoutpb.AccountLockoutServiceClient) *AccountLockoutClient {
	return &AccountLockoutClient{client: client}
}

// GetAccountLockoutByUserID 根据用户ID获取账户锁定
func (c *AccountLockoutClient) GetAccountLockoutByUserID(ctx context.Context, userID string, tenantID *string) (*accountlockoutpb.AccountLockout, error) {
	req := &accountlockoutpb.GetAccountLockoutByUserIDRequest{
		UserId:   userID,
		TenantId: tenantID,
	}

	resp, err := c.client.GetAccountLockoutByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.AccountLockout, nil
}

// BatchGetAccountLockouts 批量获取账户锁定
func (c *AccountLockoutClient) BatchGetAccountLockouts(ctx context.Context, userIDs []string, tenantID *string) ([]*accountlockoutpb.AccountLockout, error) {
	req := &accountlockoutpb.BatchGetAccountLockoutsRequest{
		UserIds:  userIDs,
		TenantId: tenantID,
	}

	resp, err := c.client.BatchGetAccountLockouts(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.AccountLockouts, nil
}