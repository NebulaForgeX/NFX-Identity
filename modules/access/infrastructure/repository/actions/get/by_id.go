package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/actions"
	"nfxid/modules/access/infrastructure/repository/actions/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*actions.Action, error) {
	var m mapper.ActionModel
	if err := h.db.WithContext(ctx).Table("access.actions").Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, actions.ErrActionNotFound
		}
		return nil, err
	}
	return mapper.ActionModelToDomain(&m), nil
}
