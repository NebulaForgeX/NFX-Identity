package get

import (
	"context"
	"nfxidentity/modules/directory/domain/badges"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/badges/mapper"
)

// ByCategory 根据 Category 获取 Badge 列表，实现 badges.Get 接口
func (h *Handler) ByCategory(ctx context.Context, category string) ([]*badges.Badge, error) {
	var ms []models.Badge
	if err := h.db.WithContext(ctx).
		Where("category = ?", category).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*badges.Badge, len(ms))
	for i, m := range ms {
		result[i] = mapper.BadgeModelToDomain(&m)
	}
	return result, nil
}
