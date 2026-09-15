package get

import (
	"context"

	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*email.Email, error) {
	var rows []models.Email
	if err := h.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Order("created_at").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*email.Email, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ToDomain(&rows[i]))
	}
	return out, nil
}
