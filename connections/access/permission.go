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

// GetPermissionByID 根据ID获取权限
func (c *PermissionClient) GetPermissionByID(ctx context.Context, id string) (*permissionpb.Permission, error) {
	req := &permissionpb.GetPermissionByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetPermissionByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Permission, nil
}

// GetPermissionByKey 根据Key获取权限
func (c *PermissionClient) GetPermissionByKey(ctx context.Context, key string) (*permissionpb.Permission, error) {
	req := &permissionpb.GetPermissionByKeyRequest{
		Key: key,
	}

	resp, err := c.client.GetPermissionByKey(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Permission, nil
}

// GetAllPermissions 获取所有权限列表
func (c *PermissionClient) GetAllPermissions(ctx context.Context, isSystem *bool) ([]*permissionpb.Permission, error) {
	req := &permissionpb.GetAllPermissionsRequest{
		IsSystem: isSystem,
	}

	resp, err := c.client.GetAllPermissions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Permissions, nil
}

// BatchGetPermissions 批量获取权限
func (c *PermissionClient) BatchGetPermissions(ctx context.Context, ids []string) ([]*permissionpb.Permission, error) {
	req := &permissionpb.BatchGetPermissionsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetPermissions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Permissions, nil
}
