package get

import (
	"context"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/mapper"

	"github.com/google/uuid"
)

// ActiveByMemberID 根据 MemberID 获取活跃的 MemberGroup 列表，实现 member_groups.Get 接口
func (h *Handler) ActiveByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member_groups.MemberGroup, error) {
	var ms []models.MemberGroup
	if err := h.db.WithContext(ctx).
		Where("member_id = ? AND revoked_at IS NULL", memberID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_groups.MemberGroup, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberGroupModelToDomain(&m)
	}
	return result, nil
}
