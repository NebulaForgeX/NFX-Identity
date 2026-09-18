package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) AuthorityByAccountAndID(ctx context.Context, accountID, profileID uuid.UUID) (*profileDomain.AuthorityProfile, error) {
	var m rdbmodels.AuthorityProfile
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileCols.ID+" = ? AND "+rdbmodels.AuthorityProfileCols.AccountID+" = ?", profileID.String(), accountID.String()).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrAuthorityProfileNotFound
		}
		return nil, err
	}
	return mapper.AuthorityProfileModelToDomain(&m), nil
}

func (h *Handler) ForgerByAccountAndID(ctx context.Context, accountID, profileID uuid.UUID) (*profileDomain.ForgerProfile, error) {
	var m rdbmodels.ForgerProfile
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileCols.ID+" = ? AND "+rdbmodels.ForgerProfileCols.AccountID+" = ?", profileID.String(), accountID.String()).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrForgerProfileNotFound
		}
		return nil, err
	}
	return mapper.ForgerProfileModelToDomain(&m), nil
}
