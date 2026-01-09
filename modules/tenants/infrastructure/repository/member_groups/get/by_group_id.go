package get

import (
	"context"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/mapper"

	"github.com/google/uuid"
)

// ByGroupID 根据 GroupID 获取 MemberGroup 列表，实现 member_groups.Get 接口
func (h *Handler) ByGroupID(ctx context.Context, groupID uuid.UUID) ([]*member_groups.MemberGroup, error) {
	var ms []models.MemberGroup
	if err := h.db.WithContext(ctx).
		Where("group_id = ?", groupID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_groups.MemberGroup, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberGroupModelToDomain(&m)
	}
	return result, nil
}
