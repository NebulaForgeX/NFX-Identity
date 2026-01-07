package list

import (
	"context"
	permissionDomainViews "nfxid/modules/permission/domain/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"
)

// ByTags 根据 Tags 获取 Permission 列表，实现 permissionDomain.List 接口
func (h *Handler) ByTags(ctx context.Context, tags []string) ([]*permissionDomainViews.PermissionView, error) {
	var items []models.Permission
	if err := h.db.WithContext(ctx).
		Where("tag IN ?", tags).
		Where("deleted_at IS NULL").
		Order("tag ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	result := slice.MapP(items, func(m *models.Permission) *permissionDomainViews.PermissionView {
		view := permissionModelToDomainView(m)
		return &view
	})
	return result, nil
}
