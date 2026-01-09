package delete

import (
	"context"
	"nfxid/modules/system/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 SystemState，实现 system_state.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.SystemState{}).Error
}
