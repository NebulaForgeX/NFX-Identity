package get

import (
	"context"
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/mapper"

	"github.com/google/uuid"
)

// ByMemberID 根据 MemberID 获取 MemberRole 列表，实现 member_roles.Get 接口
func (h *Handler) ByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member_roles.MemberRole, error) {
	var ms []models.MemberRole
	if err := h.db.WithContext(ctx).
		Where("member_id = ?", memberID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_roles.MemberRole, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberRoleModelToDomain(&m)
	}
	return result, nil
}
