package get

import (
	"context"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/grants/mapper"

	"github.com/google/uuid"
)

// ByGrantTypeAndRefID 根据 GrantType 和 GrantRefID 获取 Grant 列表，实现 grants.Get 接口
func (h *Handler) ByGrantTypeAndRefID(ctx context.Context, grantType grants.GrantType, grantRefID uuid.UUID) ([]*grants.Grant, error) {
	var ms []models.Grant
	if err := h.db.WithContext(ctx).
		Where("grant_type = ? AND grant_ref_id = ?", grantType, grantRefID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*grants.Grant, len(ms))
	for i, m := range ms {
		result[i] = mapper.GrantModelToDomain(&m)
	}
	return result, nil
}
