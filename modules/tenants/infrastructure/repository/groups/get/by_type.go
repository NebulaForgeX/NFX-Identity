package get

import (
	"context"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"
)

// ByType 根据 Type 获取 Group 列表，实现 groups.Get 接口
func (h *Handler) ByType(ctx context.Context, groupType groups.GroupType) ([]*groups.Group, error) {
	typeEnum := mapper.GroupTypeDomainToEnum(groupType)
	var ms []models.Group
	if err := h.db.WithContext(ctx).
		Where("type = ?", typeEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*groups.Group, len(ms))
	for i, m := range ms {
		result[i] = mapper.GroupModelToDomain(&m)
	}
	return result, nil
}
