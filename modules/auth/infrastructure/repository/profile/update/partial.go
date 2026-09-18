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

//* =============================== Authority partial =============================== !//
func (h *Handler) AuthorityPartial(ctx context.Context, profileID uuid.UUID, p profileDomain.AuthorityProfilePatch) error {
	updates := mapper.AuthorityProfilePatchToUpdates(p)
	if len(updates) == 0 {
		return nil
	}
	updates[rdbmodels.AuthorityProfileCols.UpdatedAt] = time.Now().UTC()

	res := h.db.WithContext(ctx).Model(&rdbmodels.AuthorityProfile{}).
		Where(rdbmodels.AuthorityProfileCols.ID+" = ?", profileID.String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrAuthorityProfileNotFound
	}
	return nil
}

//* =============================== Forger partial =============================== !//
func (h *Handler) ForgerPartial(ctx context.Context, profileID uuid.UUID, p profileDomain.ForgerProfilePatch) error {
	updates := mapper.ForgerProfilePatchToUpdates(p)
	if len(updates) == 0 {
		return nil
	}
	updates[rdbmodels.ForgerProfileCols.UpdatedAt] = time.Now().UTC()

	res := h.db.WithContext(ctx).Model(&rdbmodels.ForgerProfile{}).
		Where(rdbmodels.ForgerProfileCols.ID+" = ?", profileID.String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrForgerProfileNotFound
	}
	return nil
}
