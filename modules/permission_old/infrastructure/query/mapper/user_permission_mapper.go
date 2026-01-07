package mapper

import (
	userPermissionAppViews "nfxid/modules/permission/application/user_permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"
)

// UserPermissionModelToAppView 将 UserPermission model 转换为 Application View
// 注意：这个函数只转换基本字段，tag/name/category 需要通过 JOIN 查询获取
func UserPermissionModelToAppView(m *models.UserPermission) userPermissionAppViews.UserPermissionView {
	return userPermissionAppViews.UserPermissionView{
		ID:           m.ID,
		UserID:       m.UserID,
		PermissionID: m.PermissionID,
		Tag:          "", // 需要通过 JOIN 查询获取
		Name:         "", // 需要通过 JOIN 查询获取
		Category:     "", // 需要通过 JOIN 查询获取
		CreatedAt:    m.CreatedAt,
	}
}

