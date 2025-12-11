package mapper

import (
	userAppViews "nfxid/modules/auth/application/user/views"
	authpb "nfxid/protos/gen/auth/auth"
	rolepb "nfxid/protos/gen/auth/role"
	userpb "nfxid/protos/gen/auth/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserViewToProtoAuth 将 UserView 转换为 proto Auth 消息
func UserViewToProtoAuth(v *userAppViews.UserView) *authpb.Auth {
	if v == nil {
		return nil
	}

	auth := &authpb.Auth{
		Id:         v.ID.String(),
		Username:   v.Username,
		Email:      v.Email,
		Phone:      v.Phone,
		Status:     statusToProto(v.Status),
		IsVerified: v.IsVerified,
		CreatedAt:  timestamppb.New(v.CreatedAt),
		UpdatedAt:  timestamppb.New(v.UpdatedAt),
	}

	if v.LastLoginAt != nil {
		auth.LastLoginAt = timestamppb.New(*v.LastLoginAt)
	}

	// 检查是否有密码（通过 UserView 无法判断，假设已设置）
	auth.PasswordSet = true

	// 嵌套用户信息
	auth.User = UserViewToProtoUser(v)

	// 嵌套角色信息（多个角色）
	if len(v.Roles) > 0 {
		auth.Roles = make([]*rolepb.Role, 0, len(v.Roles))
		for _, role := range v.Roles {
			auth.Roles = append(auth.Roles, RoleViewToProtoRole(&role))
		}
	}

	return auth
}

// UserViewToProtoUser 将 UserView 转换为 proto User 消息
func UserViewToProtoUser(v *userAppViews.UserView) *userpb.User {
	if v == nil {
		return nil
	}

	user := &userpb.User{
		Id:         v.ID.String(),
		Username:   v.Username,
		Email:      v.Email,
		Phone:      v.Phone,
		Status:     userStatusToProto(v.Status),
		IsVerified: v.IsVerified,
		CreatedAt:  timestamppb.New(v.CreatedAt),
		UpdatedAt:  timestamppb.New(v.UpdatedAt),
	}

	if v.LastLoginAt != nil {
		user.LastLoginAt = timestamppb.New(*v.LastLoginAt)
	}

	// 嵌套角色信息（多个角色）
	if len(v.Roles) > 0 {
		user.Roles = make([]*rolepb.Role, 0, len(v.Roles))
		for _, role := range v.Roles {
			user.Roles = append(user.Roles, RoleViewToProtoRole(&role))
		}
	}

	return user
}

// RoleViewToProtoRole 将 RoleView 转换为 proto Role 消息
// 注意：这里使用 userAppViews.RoleView（因为 UserView 中嵌套的是这个类型）
func RoleViewToProtoRole(v *userAppViews.RoleView) *rolepb.Role {
	if v == nil {
		return nil
	}

	role := &rolepb.Role{
		Id:       v.ID.String(),
		Name:     v.Name,
		IsSystem: v.IsSystem,
	}

	if v.Description != nil {
		role.Description = v.Description
	}
	if v.Permissions != nil {
		// Note: Permissions is *datatypes.JSON, need to convert to []string
		// For now, leave it empty or implement conversion if needed
		role.Permissions = []string{}
	}

	return role
}

// statusToProto 将状态字符串转换为 proto AuthStatus
func statusToProto(status string) authpb.AuthStatus {
	switch status {
	case "pending":
		return authpb.AuthStatus_AUTH_STATUS_PENDING
	case "active":
		return authpb.AuthStatus_AUTH_STATUS_ACTIVE
	case "deactive":
		return authpb.AuthStatus_AUTH_STATUS_DEACTIVE
	default:
		return authpb.AuthStatus_AUTH_STATUS_UNSPECIFIED
	}
}

// userStatusToProto 将状态字符串转换为 proto UserStatus
func userStatusToProto(status string) userpb.UserStatus {
	switch status {
	case "pending":
		return userpb.UserStatus_USER_STATUS_PENDING
	case "active":
		return userpb.UserStatus_USER_STATUS_ACTIVE
	case "deactive":
		return userpb.UserStatus_USER_STATUS_DEACTIVE
	default:
		return userpb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}
