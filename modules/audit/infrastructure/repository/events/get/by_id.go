package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Event，实现 events.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*events.Event, error) {
	var m models.Event
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, events.ErrEventNotFound
		}
		return nil, err
	}
	return mapper.EventModelToDomain(&m), nil
}
