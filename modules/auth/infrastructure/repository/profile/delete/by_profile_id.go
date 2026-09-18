package delete

import (
	"context"
	"github.com/google/uuid"
	authErr "nfxidentity/errors/src/auth"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
)

//* =============================== Authority by_profile_id =============================== !//
func (h *Handler) AuthorityByProfileID(ctx context.Context, profileID uuid.UUID) error {
	res := h.db.WithContext(ctx).Where(rdbmodels.AuthorityProfileCols.ID+" = ?", profileID.String()).Delete(&rdbmodels.AuthorityProfile{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrAuthorityProfileNotFound
	}
	return nil
}

//* =============================== Forger by_profile_id =============================== !//
func (h *Handler) ForgerByProfileID(ctx context.Context, profileID uuid.UUID) error {
	res := h.db.WithContext(ctx).Where(rdbmodels.ForgerProfileCols.ID+" = ?", profileID.String()).Delete(&rdbmodels.ForgerProfile{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrForgerProfileNotFound
	}
	return nil
}
