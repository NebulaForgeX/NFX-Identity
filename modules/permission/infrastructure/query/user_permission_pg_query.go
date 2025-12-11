package query

import (
	"context"
	userPermissionAppViews "nfxid/modules/permission/application/user_permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userPermissionPGQuery struct {
	db *gorm.DB
}

func NewUserPermissionPGQuery(db *gorm.DB) *userPermissionPGQuery {
	return &userPermissionPGQuery{db: db}
}

// userPermissionWithPermission 用于接收 JOIN 查询结果
type userPermissionWithPermission struct {
	models.UserPermission
	PermissionTag     string `gorm:"column:permission_tag"`
	PermissionName    string `gorm:"column:permission_name"`
	PermissionCategory string `gorm:"column:permission_category"`
}

func (q *userPermissionPGQuery) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*userPermissionAppViews.UserPermissionView, error) {
	var items []userPermissionWithPermission
	if err := q.db.WithContext(ctx).
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
	result := make([]*userPermissionAppViews.UserPermissionView, len(items))
	for i, item := range items {
		category := ""
		if item.PermissionCategory != "" {
			category = item.PermissionCategory
		}
		result[i] = &userPermissionAppViews.UserPermissionView{
			ID:           item.ID,
			UserID:       item.UserID,
			PermissionID: item.PermissionID,
			Tag:          item.PermissionTag,
			Name:         item.PermissionName,
			Category:     category,
			CreatedAt:    item.CreatedAt,
		}
	}
	return result, nil
}

func (q *userPermissionPGQuery) GetPermissionTagsByUserID(ctx context.Context, userID uuid.UUID) ([]string, error) {
	var tags []string
	err := q.db.WithContext(ctx).
		Model(&models.UserPermission{}).
		Select("p.tag").
		Joins("JOIN permission.permissions p ON p.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ?", userID).
		Where("user_permissions.deleted_at IS NULL").
		Where("p.deleted_at IS NULL").
		Pluck("p.tag", &tags).Error
	return tags, err
}

