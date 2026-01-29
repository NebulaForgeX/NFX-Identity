package check

import (
	"context"

	"nfxid/modules/access/infrastructure/repository/actions/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).Model(&mapper.ActionModel{}).
		Where("id = ?", id).Count(&count).Error
	return count > 0, err
}
