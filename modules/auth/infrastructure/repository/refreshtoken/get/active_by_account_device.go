package get

import (
	"context"
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ActiveByAccountDevice(ctx context.Context, accountID uuid.UUID, deviceID string) ([]*refreshtoken.RefreshToken, error) {
	var rows []models.Refreshtoken
	if err := h.db.WithContext(ctx).Where("account_id = ? AND device_id = ? AND revoked_at IS NULL AND deleted_at IS NULL", accountID, deviceID).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*refreshtoken.RefreshToken, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ToDomain(&rows[i]))
	}
	return out, nil
}
