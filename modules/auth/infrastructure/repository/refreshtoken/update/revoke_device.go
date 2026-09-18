package update

import (
	"context"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) RevokeDevice(ctx context.Context, accountID uuid.UUID, deviceID string, at time.Time) error {
	return h.db.WithContext(ctx).Model(&models.Refreshtoken{}).Where("account_id = ? AND device_id = ? AND revoked_at IS NULL", accountID, deviceID).Update("revoked_at", at).Error
}
