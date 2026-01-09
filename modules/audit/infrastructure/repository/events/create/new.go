package create

import (
	"context"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"
)

// New 创建新的 Event，实现 events.Create 接口
func (h *Handler) New(ctx context.Context, e *events.Event) error {
	m := mapper.EventDomainToModel(e)
	return h.db.WithContext(ctx).Create(&m).Error
}
