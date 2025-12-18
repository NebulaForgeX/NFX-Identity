package single

import (
	"context"
	"errors"
	userPermissionDomainErrors "nfxid/modules/permission/domain/user_permission/errors"
	userPermissionDomainViews "nfxid/modules/permission/domain/user_permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// userPermissionWithPermission 用于接收 JOIN 查询结果
type userPermissionWithPermission struct {
	models.UserPermission
	PermissionTag      string `gorm:"column:permission_tag"`
	PermissionName     string `gorm:"column:permission_name"`
	PermissionCategory string `gorm:"column:permission_category"`
}

// ByUserIDAndPermissionID 根据 UserID 和 PermissionID 获取 UserPermission，实现 userPermissionDomain.Single 接口
func (h *Handler) ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (*userPermissionDomainViews.UserPermissionView, error) {
	var item userPermissionWithPermission
	if err := h.db.WithContext(ctx).
		Table("permission.user_permissions").
		Select("user_permissions.*, p.tag as permission_tag, p.name as permission_name, p.category as permission_category").
		Joins("JOIN permission.permissions p ON p.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ? AND user_permissions.permission_id = ?", userID, permissionID).
		Where("user_permissions.deleted_at IS NULL").
		Where("p.deleted_at IS NULL").
		First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userPermissionDomainErrors.ErrUserPermissionNotFound
		}
		return nil, err
	}
	view := userPermissionModelToDomainView(&item)
	return &view, nil
}

func userPermissionModelToDomainView(item *userPermissionWithPermission) userPermissionDomainViews.UserPermissionView {
	return userPermissionDomainViews.UserPermissionView{
		ID:           item.ID,
		UserID:       item.UserID,
		PermissionID: item.PermissionID,
		Tag:          item.PermissionTag,
		Name:         item.PermissionName,
		Category:     item.PermissionCategory,
		CreatedAt:    item.CreatedAt,
	}
}
