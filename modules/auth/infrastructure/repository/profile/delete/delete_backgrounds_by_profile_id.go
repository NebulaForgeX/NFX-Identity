package delete

import (
	"context"
	"github.com/google/uuid"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
)

//* =============================== Authority delete_backgrounds_by_profile_id =============================== !//
func (h *Handler) DeleteAuthorityBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileBackgroundCols.ProfileID+" = ?", profileID.String()).
		Delete(&rdbmodels.AuthorityProfileBackground{}).Error
}

//* =============================== Forger delete_backgrounds_by_profile_id =============================== !//
func (h *Handler) DeleteForgerBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileBackgroundCols.ProfileID+" = ?", profileID.String()).
		Delete(&rdbmodels.ForgerProfileBackground{}).Error
}
