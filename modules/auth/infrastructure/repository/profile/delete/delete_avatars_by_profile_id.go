package delete

import (
	"context"
	"github.com/google/uuid"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
)

//* =============================== Authority delete_avatars_by_profile_id =============================== !//
func (h *Handler) DeleteAuthorityAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileAvatarCols.ProfileID+" = ?", profileID.String()).
		Delete(&rdbmodels.AuthorityProfileAvatar{}).Error
}

//* =============================== Forger delete_avatars_by_profile_id =============================== !//
func (h *Handler) DeleteForgerAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileAvatarCols.ProfileID+" = ?", profileID.String()).
		Delete(&rdbmodels.ForgerProfileAvatar{}).Error
}
