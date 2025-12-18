package list

import (
	"context"
	"nfxid/enums"
	userPermissionDomainViews "nfxid/modules/permission/domain/user_permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// userPermissionWithPermission 用于接收 JOIN 查询结果
type userPermissionWithPermission struct {
	models.UserPermission
	PermissionTag      string `gorm:"column:permission_tag"`
	PermissionName     string `gorm:"column:permission_name"`
	PermissionCategory string `gorm:"column:permission_category"`
}

// ByUserID 根据 UserID 获取 UserPermission 列表，实现 userPermissionDomain.List 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*userPermissionDomainViews.UserPermissionView, error) {
	var items []userPermissionWithPermission
	if err := h.db.WithContext(ctx).
		Table("permission.user_permissions").
		Select("user_permissions.*, p.tag as permission_tag, p.name as permission_name, p.category as permission_category").
		Joins("JOIN permission.permissions p ON p.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ?", userID).
		Where("user_permissions.deleted_at IS NULL").
		Where("p.deleted_at IS NULL").
		Order("user_permissions.created_at DESC").
		Scan(&items).Error; err != nil {
		return nil, err
	}
	result := make([]*userPermissionDomainViews.UserPermissionView, len(items))
	for i, item := range items {
		view := userPermissionModelToDomainView(&item)
		result[i] = &view
	}
	return result, nil
}

func userPermissionModelToDomainView(item *userPermissionWithPermission) userPermissionDomainViews.UserPermissionView {
	return userPermissionDomainViews.UserPermissionView{
		ID:           item.ID,
		UserID:       item.UserID,
		PermissionID: item.PermissionID,
		Tag:          item.PermissionTag,
		Name:         item.PermissionName,
		Category:     enums.PermissionCategory(item.PermissionCategory), // Convert string to enum
		CreatedAt:    item.CreatedAt,
	}
}
