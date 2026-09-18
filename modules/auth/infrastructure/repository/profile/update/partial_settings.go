package update

import (
	"context"
	"github.com/google/uuid"
	authErr "nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
	"time"
)

//* =============================== Authority partial_settings =============================== !//
func (h *Handler) AuthorityPartialSettings(ctx context.Context, profileID uuid.UUID, p profileDomain.AuthorityProfileSettingsPatch) error {
	updates := mapper.AuthorityProfileSettingsPatchToUpdates(p)
	if len(updates) == 0 {
		return nil
	}
	updates[rdbmodels.AuthorityProfileSettingCols.UpdatedAt] = time.Now().UTC()
	res := h.db.WithContext(ctx).
		Model(&rdbmodels.AuthorityProfileSetting{}).
		Where(rdbmodels.AuthorityProfileSettingCols.ID+" = ?", profileID.String()).
		Updates(updates)
	if res.Error != nil {
		return authErr.ErrAuthorityProfileSettingsUpdateFailed.WithCause(res.Error)
	}
	if res.RowsAffected == 0 {
		return authErr.ErrAuthorityProfileSettingsNotFound
	}
	return nil
}

//* =============================== Forger partial_settings =============================== !//
func (h *Handler) ForgerPartialSettings(ctx context.Context, profileID uuid.UUID, p profileDomain.ForgerProfileSettingsPatch) error {
	updates := mapper.ForgerProfileSettingsPatchToUpdates(p)
	if len(updates) == 0 {
		return nil
	}
	updates[rdbmodels.ForgerProfileSettingCols.UpdatedAt] = time.Now().UTC()
	res := h.db.WithContext(ctx).
		Model(&rdbmodels.ForgerProfileSetting{}).
		Where(rdbmodels.ForgerProfileSettingCols.ID+" = ?", profileID.String()).
		Updates(updates)
	if res.Error != nil {
		return authErr.ErrForgerProfileSettingsUpdateFailed.WithCause(res.Error)
	}
	if res.RowsAffected == 0 {
		return authErr.ErrForgerProfileSettingsNotFound
	}
	return nil
}
