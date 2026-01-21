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
