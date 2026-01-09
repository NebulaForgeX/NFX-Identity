package update

import (
	"context"
	"nfxid/modules/system/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
)

// Initialize 初始化系统，实现 system_state.Update 接口
func (h *Handler) Initialize(ctx context.Context, version string) error {
	now := time.Now().UTC()
	return h.db.WithContext(ctx).
		Model(&models.SystemState{}).
		Where("id = (SELECT id FROM system.system_state ORDER BY created_at DESC LIMIT 1)").
		Updates(map[string]any{
			models.SystemStateCols.Initialized:           true,
			models.SystemStateCols.InitializedAt:         &now,
			models.SystemStateCols.InitializationVersion: &version,
			models.SystemStateCols.UpdatedAt:             now,
		}).Error
}
