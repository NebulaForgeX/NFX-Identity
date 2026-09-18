package get

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	authErr "nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority by_profile_id =============================== !//
func (h *Handler) AuthorityByProfileID(ctx context.Context, profileID uuid.UUID) (*profileDomain.AuthorityProfile, error) {
	var m rdbmodels.AuthorityProfile
	if err := h.db.WithContext(ctx).Where(rdbmodels.AuthorityProfileCols.ID+" = ?", profileID.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAuthorityProfileNotFound
		}
		return nil, err
	}
	return mapper.AuthorityProfileModelToDomain(&m), nil
}

//* =============================== Forger by_profile_id =============================== !//
func (h *Handler) ForgerByProfileID(ctx context.Context, profileID uuid.UUID) (*profileDomain.ForgerProfile, error) {
	var m rdbmodels.ForgerProfile
	if err := h.db.WithContext(ctx).Where(rdbmodels.ForgerProfileCols.ID+" = ?", profileID.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrForgerProfileNotFound
		}
		return nil, err
	}
	return mapper.ForgerProfileModelToDomain(&m), nil
}
