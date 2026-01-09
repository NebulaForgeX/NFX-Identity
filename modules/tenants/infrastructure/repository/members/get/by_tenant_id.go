package get

import (
	"context"
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/members/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 Member 列表，实现 members.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*members.Member, error) {
	var ms []models.Member
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*members.Member, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberModelToDomain(&m)
	}
	return result, nil
}
