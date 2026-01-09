package get

import (
	"context"
	"time"
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/mapper"

	"github.com/google/uuid"
)

// ActiveByMemberID 根据 MemberID 获取活跃的 MemberRole 列表，实现 member_roles.Get 接口
func (h *Handler) ActiveByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member_roles.MemberRole, error) {
	now := time.Now().UTC()
	var ms []models.MemberRole
	if err := h.db.WithContext(ctx).
		Where("member_id = ? AND (expires_at IS NULL OR expires_at > ?) AND revoked_at IS NULL", memberID, now).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_roles.MemberRole, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberRoleModelToDomain(&m)
	}
	return result, nil
}
