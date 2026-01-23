package auth

import (
	"context"
	"fmt"

	refreshtokenpb "nfxid/protos/gen/auth/refresh_token"
)

// RefreshTokenClient RefreshToken 客户端
type RefreshTokenClient struct {
	client refreshtokenpb.RefreshTokenServiceClient
}

// NewRefreshTokenClient 创建 RefreshToken 客户端
func NewRefreshTokenClient(client refreshtokenpb.RefreshTokenServiceClient) *RefreshTokenClient {
	return &RefreshTokenClient{client: client}
}

// GetRefreshTokenByID 根据ID获取刷新令牌
func (c *RefreshTokenClient) GetRefreshTokenByID(ctx context.Context, id string) (*refreshtokenpb.RefreshToken, error) {
	req := &refreshtokenpb.GetRefreshTokenByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetRefreshTokenByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RefreshToken, nil
}

// GetRefreshTokenByTokenID 根据令牌ID获取刷新令牌
func (c *RefreshTokenClient) GetRefreshTokenByTokenID(ctx context.Context, tokenID string) (*refreshtokenpb.RefreshToken, error) {
	req := &refreshtokenpb.GetRefreshTokenByTokenIDRequest{
		TokenId: tokenID,
	}

	resp, err := c.client.GetRefreshTokenByTokenID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RefreshToken, nil
}

// GetRefreshTokensByUserID 根据用户ID获取刷新令牌列表
func (c *RefreshTokenClient) GetRefreshTokensByUserID(ctx context.Context, userID string) ([]*refreshtokenpb.RefreshToken, error) {
	req := &refreshtokenpb.GetRefreshTokensByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetRefreshTokensByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RefreshTokens, nil
}