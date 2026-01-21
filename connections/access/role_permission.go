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
