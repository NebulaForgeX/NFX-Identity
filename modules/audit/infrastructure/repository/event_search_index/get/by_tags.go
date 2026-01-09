package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"
	"strings"

	"gorm.io/gorm"
)

// ByTags 根据 Tags 获取 EventSearchIndexes，实现 event_search_index.Get 接口
func (h *Handler) ByTags(ctx context.Context, tags []string) ([]*event_search_index.EventSearchIndex, error) {
	if len(tags) == 0 {
		return []*event_search_index.EventSearchIndex{}, nil
	}
	
	// PostgreSQL array contains operator
	tagsStr := "{" + strings.Join(tags, ",") + "}"
	query := h.db.WithContext(ctx).Where("tags && ?", tagsStr)
	
	var ms []models.EventSearchIndex
	if err := query.Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*event_search_index.EventSearchIndex{}, nil
		}
		return nil, err
	}
	
	result := make([]*event_search_index.EventSearchIndex, len(ms))
	for i := range ms {
		result[i] = mapper.EventSearchIndexModelToDomain(&ms[i])
	}
	return result, nil
}
