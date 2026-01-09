package get

import (
	"context"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/mapper"

	"github.com/google/uuid"
)

// ByMemberID 根据 MemberID 获取 MemberGroup 列表，实现 member_groups.Get 接口
func (h *Handler) ByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member_groups.MemberGroup, error) {
	var ms []models.MemberGroup
	if err := h.db.WithContext(ctx).
		Where("member_id = ?", memberID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_groups.MemberGroup, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberGroupModelToDomain(&m)
	}
	return result, nil
}
