package get

import (
	"context"
	"nfxid/modules/tenants/domain/member_app_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_app_roles/mapper"

	"github.com/google/uuid"
)

// ByMemberIDAndAppID 根据 MemberID 和 AppID 获取 MemberAppRole 列表，实现 member_app_roles.Get 接口
func (h *Handler) ByMemberIDAndAppID(ctx context.Context, memberID, appID uuid.UUID) ([]*member_app_roles.MemberAppRole, error) {
	var ms []models.MemberAppRole
	if err := h.db.WithContext(ctx).
		Where("member_id = ? AND app_id = ?", memberID, appID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_app_roles.MemberAppRole, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberAppRoleModelToDomain(&m)
	}
	return result, nil
}
