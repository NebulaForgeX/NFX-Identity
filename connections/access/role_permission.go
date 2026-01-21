package access

import (
	"context"
	"fmt"

	rolepermissionpb "nfxid/protos/gen/access/role_permission"
)

// RolePermissionClient RolePermission 客户端
type RolePermissionClient struct {
	client rolepermissionpb.RolePermissionServiceClient
}

// NewRolePermissionClient 创建 RolePermission 客户端
func NewRolePermissionClient(client rolepermissionpb.RolePermissionServiceClient) *RolePermissionClient {
	return &RolePermissionClient{client: client}
}

// CreateRolePermission 创建角色权限关联
func (c *RolePermissionClient) CreateRolePermission(ctx context.Context, roleID, permissionID string) (string, error) {
	req := &rolepermissionpb.CreateRolePermissionRequest{
		RoleId:       roleID,
		PermissionId: permissionID,
	}

	resp, err := c.client.CreateRolePermission(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RolePermission.Id, nil
}

// GetRolePermissionByID 根据ID获取角色权限关联
func (c *RolePermissionClient) GetRolePermissionByID(ctx context.Context, id string) (*rolepermissionpb.RolePermission, error) {
	req := &rolepermissionpb.GetRolePermissionByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetRolePermissionByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RolePermission, nil
}

// GetPermissionsByRole 根据角色获取权限列表
func (c *RolePermissionClient) GetPermissionsByRole(ctx context.Context, roleID string) ([]*rolepermissionpb.RolePermission, error) {
	req := &rolepermissionpb.GetPermissionsByRoleRequest{
		RoleId: roleID,
	}

	resp, err := c.client.GetPermissionsByRole(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RolePermissions, nil
}

// GetRolesByPermission 根据权限获取角色列表
func (c *RolePermissionClient) GetRolesByPermission(ctx context.Context, permissionID string) ([]*rolepermissionpb.RolePermission, error) {
	req := &rolepermissionpb.GetRolesByPermissionRequest{
		PermissionId: permissionID,
	}

	resp, err := c.client.GetRolesByPermission(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RolePermissions, nil
}
