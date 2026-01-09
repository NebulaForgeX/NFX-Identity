package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/grants/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Grant，实现 grants.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*grants.Grant, error) {
	var m models.Grant
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, grants.ErrGrantNotFound
		}
		return nil, err
	}
	return mapper.GrantModelToDomain(&m), nil
}
