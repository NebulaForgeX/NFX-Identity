package grpc

import (
	"context"
	"fmt"

	"nfxid/modules/system/application/bootstrap"

	grantpb "nfxid/protos/gen/access/grant"
	permissionpb "nfxid/protos/gen/access/permission"
	rolepb "nfxid/protos/gen/access/role"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"
	userpb "nfxid/protos/gen/directory/user"
)

// Adapter 将 infrastructure 层的 gRPC 客户端适配为 application 层的接口
// 实现 bootstrap.GRPCClients 接口
type Adapter struct {
	clients *Clients
}

// NewAdapter 创建适配器
func NewAdapter(clients *Clients) bootstrap.GRPCClients {
	return &Adapter{clients: clients}
}

// CreateUser 创建用户
func (a *Adapter) CreateUser(ctx context.Context, username string, status string, isVerified bool) (string, error) {
	// 转换状态枚举
	var userStatus userpb.DirectoryUserStatus
	switch status {
	case "active":
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_ACTIVE
	case "pending":
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_PENDING
	case "deactive":
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_DEACTIVE
	default:
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_ACTIVE
	}

	req := &userpb.CreateUserRequest{
		Username:   username,
		Status:     userStatus,
		IsVerified: isVerified,
	}

	resp, err := a.clients.DirectoryClient.CreateUser(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.User.Id, nil
}

// CreateRole 创建角色
func (a *Adapter) CreateRole(ctx context.Context, key, name string, description *string, scopeType string, isSystem bool) (string, error) {
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

	resp, err := a.clients.RoleClient.CreateRole(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Role.Id, nil
}

// CreatePermission 创建权限
func (a *Adapter) CreatePermission(ctx context.Context, key, name string, description *string, isSystem bool) (string, error) {
	req := &permissionpb.CreatePermissionRequest{
		Key:         key,
		Name:        name,
		Description: description,
		IsSystem:    isSystem,
	}

	resp, err := a.clients.PermissionClient.CreatePermission(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Permission.Id, nil
}

// CreateRolePermission 创建角色权限关联
func (a *Adapter) CreateRolePermission(ctx context.Context, roleID, permissionID string) (string, error) {
	req := &rolepermissionpb.CreateRolePermissionRequest{
		RoleId:       roleID,
		PermissionId: permissionID,
	}

	resp, err := a.clients.RolePermissionClient.CreateRolePermission(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RolePermission.Id, nil
}

// CreateGrant 创建授权
func (a *Adapter) CreateGrant(ctx context.Context, subjectType, subjectID, grantType, grantRefID string, tenantID, appID *string) (string, error) {
	// 转换主体类型枚举
	var subjType grantpb.AccessSubjectType
	switch subjectType {
	case "user":
		subjType = grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_USER
	case "client":
		subjType = grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_CLIENT
	default:
		return "", fmt.Errorf("invalid subject type: %s", subjectType)
	}

	// 转换授权类型枚举
	var gType grantpb.AccessGrantType
	switch grantType {
	case "role":
		gType = grantpb.AccessGrantType_ACCESS_GRANT_TYPE_ROLE
	case "permission":
		gType = grantpb.AccessGrantType_ACCESS_GRANT_TYPE_PERMISSION
	default:
		return "", fmt.Errorf("invalid grant type: %s", grantType)
	}

	req := &grantpb.CreateGrantRequest{
		SubjectType: subjType,
		SubjectId:   subjectID,
		GrantType:   gType,
		GrantRefId:  grantRefID,
		TenantId:    tenantID,
		AppId:       appID,
		Effect:      grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_ALLOW,
	}

	resp, err := a.clients.GrantClient.CreateGrant(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Grant.Id, nil
}

// CreateUserCredential 创建用户凭证
func (a *Adapter) CreateUserCredential(ctx context.Context, userID, password string, tenantID *string, mustChangePassword bool) error {

	req := &usercredentialpb.CreateUserCredentialRequest{
		UserId:             userID,
		TenantId:           tenantID,
		CredentialType:     usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSWORD,
		Password:           password,
		MustChangePassword: mustChangePassword,
	}

	_, err := a.clients.UserCredentialClient.CreateUserCredential(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}

	return nil
}
