package update

import (
	"context"
	"nfxid/modules/system/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reset 重置系统，实现 system_state.Update 接口
func (h *Handler) Reset(ctx context.Context, resetBy uuid.UUID) error {
	now := time.Now().UTC()
	return h.db.WithContext(ctx).
		Model(&models.SystemState{}).
		Where("id = (SELECT id FROM system.system_state ORDER BY created_at DESC LIMIT 1)").
		Updates(map[string]any{
			models.SystemStateCols.Initialized:   false,
			models.SystemStateCols.InitializedAt:  nil,
			models.SystemStateCols.LastResetAt:    &now,
			models.SystemStateCols.LastResetBy:    &resetBy,
			models.SystemStateCols.ResetCount:     gorm.Expr("reset_count + 1"),
			models.SystemStateCols.UpdatedAt:     now,
		}).Error
}
