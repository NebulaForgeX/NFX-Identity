package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 EventSearchIndex，实现 event_search_index.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*event_search_index.EventSearchIndex, error) {
	var m models.EventSearchIndex
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, event_search_index.ErrEventSearchIndexNotFound
		}
		return nil, err
	}
	return mapper.EventSearchIndexModelToDomain(&m), nil
}
