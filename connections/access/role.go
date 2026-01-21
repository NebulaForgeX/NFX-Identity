package access

import (
	"context"
	"fmt"

	rolepb "nfxid/protos/gen/access/role"
)

// RoleClient Role 客户端
type RoleClient struct {
	client rolepb.RoleServiceClient
}

// NewRoleClient 创建 Role 客户端
func NewRoleClient(client rolepb.RoleServiceClient) *RoleClient {
	return &RoleClient{client: client}
}

// CreateRole 创建角色
func (c *RoleClient) CreateRole(ctx context.Context, key, name string, description *string, scopeType string, isSystem bool) (string, error) {
	// 转换范围类型枚举
	var scope rolepb.AccessScopeType
	switch scopeType {
	case "tenant":
		scope = rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_TENANT
	case "app":
		scope = rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_APP
	case "global":
		scope = rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_GLOBAL
	default:
		scope = rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_GLOBAL
	}

	req := &rolepb.CreateRoleRequest{
		Key:         key,
		Name:        name,
		Description: description,
		ScopeType:   scope,
		IsSystem:    isSystem,
	}

	resp, err := c.client.CreateRole(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Role.Id, nil
}

// GetRoleByID 根据ID获取角色
func (c *RoleClient) GetRoleByID(ctx context.Context, id string) (*rolepb.Role, error) {
	req := &rolepb.GetRoleByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetRoleByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Role, nil
}

// GetRoleByKey 根据Key获取角色
func (c *RoleClient) GetRoleByKey(ctx context.Context, key string) (*rolepb.Role, error) {
	req := &rolepb.GetRoleByKeyRequest{
		Key: key,
	}

	resp, err := c.client.GetRoleByKey(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Role, nil
}

// GetAllRoles 获取所有角色列表
func (c *RoleClient) GetAllRoles(ctx context.Context, scopeType *rolepb.AccessScopeType, isSystem *bool) ([]*rolepb.Role, error) {
	req := &rolepb.GetAllRolesRequest{
		ScopeType: scopeType,
		IsSystem:  isSystem,
	}

	resp, err := c.client.GetAllRoles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Roles, nil
}

// BatchGetRoles 批量获取角色
func (c *RoleClient) BatchGetRoles(ctx context.Context, ids []string) ([]*rolepb.Role, error) {
	req := &rolepb.BatchGetRolesRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetRoles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Roles, nil
}
