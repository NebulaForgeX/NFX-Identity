package get

import (
	"context"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"

	"github.com/google/uuid"
)

// ByParentGroupID 根据 ParentGroupID 获取 Group 列表，实现 groups.Get 接口
func (h *Handler) ByParentGroupID(ctx context.Context, parentGroupID uuid.UUID) ([]*groups.Group, error) {
	var ms []models.Group
	if err := h.db.WithContext(ctx).
		Where("parent_group_id = ?", parentGroupID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*groups.Group, len(ms))
	for i, m := range ms {
		result[i] = mapper.GroupModelToDomain(&m)
	}
	return result, nil
}
