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

//* =============================== Authority settings_by_profile_id =============================== !//
func (h *Handler) AuthoritySettingsByProfileID(ctx context.Context, profileID uuid.UUID) (*profileDomain.AuthorityProfileSettings, error) {
	var m rdbmodels.AuthorityProfileSetting
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileSettingCols.ID+" = ?", profileID.String()).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAuthorityProfileSettingsNotFound
		}
		return nil, authErr.ErrAuthorityProfileSettingsGetFailed
	}
	return mapper.AuthorityProfileSettingsModelToDomain(&m), nil
}

//* =============================== Forger settings_by_profile_id =============================== !//
func (h *Handler) ForgerSettingsByProfileID(ctx context.Context, profileID uuid.UUID) (*profileDomain.ForgerProfileSettings, error) {
	var m rdbmodels.ForgerProfileSetting
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileSettingCols.ID+" = ?", profileID.String()).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrForgerProfileSettingsNotFound
		}
		return nil, authErr.ErrForgerProfileSettingsGetFailed
	}
	return mapper.ForgerProfileSettingsModelToDomain(&m), nil
}
