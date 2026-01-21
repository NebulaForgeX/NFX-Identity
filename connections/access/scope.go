package access

import (
	"context"
	"fmt"

	scopepb "nfxid/protos/gen/access/scope"
)

// ScopeClient Scope 客户端
type ScopeClient struct {
	client scopepb.ScopeServiceClient
}

// NewScopeClient 创建 Scope 客户端
func NewScopeClient(client scopepb.ScopeServiceClient) *ScopeClient {
	return &ScopeClient{client: client}
}

// GetScopeByScope 根据Scope标识符获取范围
func (c *ScopeClient) GetScopeByScope(ctx context.Context, scope string) (*scopepb.Scope, error) {
	req := &scopepb.GetScopeByScopeRequest{
		Scope: scope,
	}

	resp, err := c.client.GetScopeByScope(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Scope, nil
}

// GetAllScopes 获取所有范围列表
func (c *ScopeClient) GetAllScopes(ctx context.Context, isSystem *bool) ([]*scopepb.Scope, error) {
	req := &scopepb.GetAllScopesRequest{
		IsSystem: isSystem,
	}

	resp, err := c.client.GetAllScopes(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Scopes, nil
}

// BatchGetScopes 批量获取范围
func (c *ScopeClient) BatchGetScopes(ctx context.Context, scopes []string) ([]*scopepb.Scope, error) {
	req := &scopepb.BatchGetScopesRequest{
		Scopes: scopes,
	}

	resp, err := c.client.BatchGetScopes(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Scopes, nil
}