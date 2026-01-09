package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Group，实现 groups.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*groups.Group, error) {
	var m models.Group
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, groups.ErrGroupNotFound
		}
		return nil, err
	}
	return mapper.GroupModelToDomain(&m), nil
}
