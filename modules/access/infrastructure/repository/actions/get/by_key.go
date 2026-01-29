package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/actions"
	"nfxid/modules/access/infrastructure/repository/actions/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByKey(ctx context.Context, key string) (*actions.Action, error) {
	var m mapper.ActionModel
	if err := h.db.WithContext(ctx).Table("access.actions").Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, actions.ErrActionNotFound
		}
		return nil, err
	}
	return mapper.ActionModelToDomain(&m), nil
}
