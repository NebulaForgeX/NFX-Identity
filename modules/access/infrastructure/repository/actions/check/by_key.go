package check

import (
	"context"

	"nfxid/modules/access/infrastructure/repository/actions/mapper"
)

func (h *Handler) ByKey(ctx context.Context, key string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).Model(&mapper.ActionModel{}).
		Where("key = ?", key).Count(&count).Error
	return count > 0, err
}
