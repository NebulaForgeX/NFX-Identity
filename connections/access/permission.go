package access

import (
	"context"
	"fmt"

	permissionpb "nfxid/protos/gen/access/permission"
)

// PermissionClient Permission 客户端
type PermissionClient struct {
	client permissionpb.PermissionServiceClient
}

// NewPermissionClient 创建 Permission 客户端
func NewPermissionClient(client permissionpb.PermissionServiceClient) *PermissionClient {
	return &PermissionClient{client: client}
}

// CreatePermission 创建权限
func (c *PermissionClient) CreatePermission(ctx context.Context, key, name string, description *string, isSystem bool) (string, error) {
	req := &permissionpb.CreatePermissionRequest{
		Key:         key,
		Name:        name,
		Description: description,
		IsSystem:    isSystem,
	}

	resp, err := c.client.CreatePermission(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Permission.Id, nil
}
