package mapper

import (
	roleAppViews "nfxid/modules/auth/application/role/views"
	rolepb "nfxid/protos/gen/auth/role"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// RoleViewToProto 将 RoleView 转换为 proto Role 消息
func RoleViewToProto(v *roleAppViews.RoleView) *rolepb.Role {
	if v == nil {
		return nil
	}

	role := &rolepb.Role{
		Id:        v.ID.String(),
		Name:      v.Name,
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		role.Description = v.Description
	}
	if v.Permissions != nil {
		// Note: Permissions is *datatypes.JSON, need to convert to []string
		// For now, leave it empty or implement conversion if needed
		role.Permissions = []string{}
	}
	if v.DeletedAt != nil {
		role.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return role
}
