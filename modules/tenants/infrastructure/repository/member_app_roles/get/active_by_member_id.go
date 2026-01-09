package get

import (
	"context"
	"nfxid/modules/tenants/domain/member_app_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_app_roles/mapper"
	"time"

	"github.com/google/uuid"
)

// ActiveByMemberID 根据 MemberID 获取活跃的 MemberAppRole 列表，实现 member_app_roles.Get 接口
func (h *Handler) ActiveByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member_app_roles.MemberAppRole, error) {
	now := time.Now().UTC()
	var ms []models.MemberAppRole
	if err := h.db.WithContext(ctx).
		Where("member_id = ? AND (expires_at IS NULL OR expires_at > ?) AND revoked_at IS NULL", memberID, now).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_app_roles.MemberAppRole, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberAppRoleModelToDomain(&m)
	}
	return result, nil
}
