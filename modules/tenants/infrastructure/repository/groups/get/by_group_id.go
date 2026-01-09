package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"

	"gorm.io/gorm"
)

// ByGroupID 根据 GroupID 获取 Group，实现 groups.Get 接口
func (h *Handler) ByGroupID(ctx context.Context, groupID string) (*groups.Group, error) {
	var m models.Group
	if err := h.db.WithContext(ctx).Where("group_id = ?", groupID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, groups.ErrGroupNotFound
		}
		return nil, err
	}
	return mapper.GroupModelToDomain(&m), nil
}
