package get

import (
	"context"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/users/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 User 列表，实现 users.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*users.User, error) {
	var ms []models.User
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*users.User, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserModelToDomain(&m)
	}
	return result, nil
}
