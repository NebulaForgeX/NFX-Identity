package get

import (
	"context"

	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByAccountIDAndDeviceID(ctx context.Context, accountID uuid.UUID, deviceID string) ([]*refreshtokenDomain.RefreshToken, error) {
	var models []rdbmodels.RefreshToken
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.RefreshTokenCols.AccountID+" = ?", accountID.String()).
		Where(rdbmodels.RefreshTokenCols.DeviceID+" = ?", deviceID).
		Where(rdbmodels.RefreshTokenCols.RevokedAt + " IS NULL").
		Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*refreshtokenDomain.RefreshToken, 0, len(models))
	for i := range models {
		result = append(result, mapper.RefreshTokenModelToDomain(&models[i]))
	}
	return result, nil
}
