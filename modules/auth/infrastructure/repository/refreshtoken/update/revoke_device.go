package update

import (
	"context"
	"time"

	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) RevokeDevice(ctx context.Context, accountID uuid.UUID, deviceID string, at time.Time) error {
	return h.db.WithContext(ctx).Model(&rdbmodels.RefreshToken{}).
		Where(rdbmodels.RefreshTokenCols.AccountID+" = ? AND "+rdbmodels.RefreshTokenCols.DeviceID+" = ? AND "+rdbmodels.RefreshTokenCols.RevokedAt+" IS NULL", accountID, deviceID).
		Update(rdbmodels.RefreshTokenCols.RevokedAt, at).Error
}
