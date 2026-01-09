package create

import (
	"context"
	"nfxid/modules/system/domain/system_state"
	"nfxid/modules/system/infrastructure/repository/system_state/mapper"
)

// New 创建新的 SystemState，实现 system_state.Create 接口
func (h *Handler) New(ctx context.Context, ss *system_state.SystemState) error {
	m := mapper.SystemStateDomainToModel(ss)
	return h.db.WithContext(ctx).Create(&m).Error
}
