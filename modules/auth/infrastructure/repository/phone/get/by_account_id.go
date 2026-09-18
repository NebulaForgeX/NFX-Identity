package get

import (
	"context"
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*phone.Phone, error) {
	var rows []models.Phone
	if err := h.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Order("created_at").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*phone.Phone, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ToDomain(&rows[i]))
	}
	return out, nil
}
