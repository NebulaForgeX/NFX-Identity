package list

import (
	"context"
	"github.com/google/uuid"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/enumx"
)

//* =============================== Authority profile_ids_by_authority_roles.go =============================== !//
func (h *AuthorityHandler) ProfileIDsByAuthorityRoles(ctx context.Context, roles []enums.AuthAuthorityRole) ([]uuid.UUID, error) {
	if len(roles) == 0 {
		return nil, nil
	}
	var ids []uuid.UUID
	if err := h.db.WithContext(ctx).
		Model(&rdbmodels.AuthorityProfile{}).
		Where(rdbmodels.AuthorityProfileCols.AuthorityRoles+" && ?", enumx.Array[enums.AuthAuthorityRole](roles)).
		Pluck(rdbmodels.AuthorityProfileCols.ID, &ids).Error; err != nil {
		return nil, authErr.ErrAuthorityProfileCountFailed.WithCause(err)
	}
	return ids, nil
}
