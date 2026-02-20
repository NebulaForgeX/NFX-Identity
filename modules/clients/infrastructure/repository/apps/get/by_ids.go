package get

import (
	"context"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"

	"github.com/google/uuid"
)

// ByIDs 根据 ID 列表批量获取 App，实现 apps.Get 接口
func (h *Handler) ByIDs(ctx context.Context, ids []uuid.UUID) ([]*apps.App, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var list []models.Application
	if err := h.db.WithContext(ctx).Where("id IN ?", ids).Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*apps.App, 0, len(list))
	for i := range list {
		out = append(out, mapper.AppModelToDomain(&list[i]))
	}
	return out, nil
}
