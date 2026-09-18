package update

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	authErr "nfxidentity/errors/src/auth"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
)

//* =============================== Authority preference =============================== !//
// Preference overwrites only the preference column for the given profile.
func (h *Handler) AuthorityPreference(ctx context.Context, profileID uuid.UUID, preference *datatypes.JSON) error {
	res := h.db.WithContext(ctx).
		Model(&rdbmodels.AuthorityProfile{}).
		Where(rdbmodels.AuthorityProfileCols.ID+" = ?", profileID.String()).
		Update(rdbmodels.AuthorityProfileCols.Preference, preference)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrAuthorityProfileNotFound
	}
	return nil
}

//* =============================== Forger preference =============================== !//
// Preference overwrites only the preference column for the given profile.
func (h *Handler) ForgerPreference(ctx context.Context, profileID uuid.UUID, preference *datatypes.JSON) error {
	res := h.db.WithContext(ctx).
		Model(&rdbmodels.ForgerProfile{}).
		Where(rdbmodels.ForgerProfileCols.ID+" = ?", profileID.String()).
		Update(rdbmodels.ForgerProfileCols.Preference, preference)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrForgerProfileNotFound
	}
	return nil
}
