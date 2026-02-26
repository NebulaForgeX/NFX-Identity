package get

import (
	"context"
	"nfxidentity/modules/tenants/domain/members"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/members/mapper"

	"github.com/google/uuid"
)

// ByTenantIDAndStatus 根据 TenantID 和 Status 获取 Member 列表，实现 members.Get 接口
func (h *Handler) ByTenantIDAndStatus(ctx context.Context, tenantID uuid.UUID, status members.MemberStatus) ([]*members.Member, error) {
	statusEnum := mapper.MemberStatusDomainToEnum(status)
	var ms []models.Member
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND status = ?", tenantID, statusEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*members.Member, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberModelToDomain(&m)
	}
	return result, nil
}
