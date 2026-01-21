package access

import (
	"context"
	"fmt"

	scopepermissionpb "nfxid/protos/gen/access/scope_permission"
)

// ScopePermissionClient ScopePermission 客户端
type ScopePermissionClient struct {
	client scopepermissionpb.ScopePermissionServiceClient
}

// NewScopePermissionClient 创建 ScopePermission 客户端
func NewScopePermissionClient(client scopepermissionpb.ScopePermissionServiceClient) *ScopePermissionClient {
	return &ScopePermissionClient{client: client}
}

// GetScopePermissionByID 根据ID获取范围权限关联
func (c *ScopePermissionClient) GetScopePermissionByID(ctx context.Context, id string) (*scopepermissionpb.ScopePermission, error) {
	req := &scopepermissionpb.GetScopePermissionByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetScopePermissionByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ScopePermission, nil
}

// GetPermissionsByScope 根据范围获取权限列表
func (c *ScopePermissionClient) GetPermissionsByScope(ctx context.Context, scope string) ([]*scopepermissionpb.ScopePermission, error) {
	req := &scopepermissionpb.GetPermissionsByScopeRequest{
		Scope: scope,
	}

	resp, err := c.client.GetPermissionsByScope(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ScopePermissions, nil
}

// GetScopesByPermission 根据权限获取范围列表
func (c *ScopePermissionClient) GetScopesByPermission(ctx context.Context, permissionID string) ([]*scopepermissionpb.ScopePermission, error) {
	req := &scopepermissionpb.GetScopesByPermissionRequest{
		PermissionId: permissionID,
	}

	resp, err := c.client.GetScopesByPermission(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ScopePermissions, nil
}