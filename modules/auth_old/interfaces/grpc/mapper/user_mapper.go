package mapper

import (
	userAppViews "nfxid/modules/auth/application/user/views"
	rolepb "nfxid/protos/gen/auth/role"
	userpb "nfxid/protos/gen/auth/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserViewToProto 将 UserView 转换为 proto User 消息
func UserViewToProto(v *userAppViews.UserView) *userpb.User {
	if v == nil {
		return nil
	}

	phone := ""
	if v.Phone != nil {
		phone = *v.Phone
	}
	user := &userpb.User{
		Id:         v.ID.String(),
		Username:   v.Username,
		Email:      v.Email,
		Phone:      phone,
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
			user.Roles = append(user.Roles, userRoleViewToProtoRole(&role))
		}
	}

	return user
}

// userRoleViewToProtoRole 将 userAppViews.RoleView 转换为 proto Role 消息
func userRoleViewToProtoRole(v *userAppViews.RoleView) *rolepb.Role {
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
