package get

import (
	"context"

	"nfxidentity/enums"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/enumx"

	"github.com/google/uuid"
)

func (h *Handler) HasOwnerRole(ctx context.Context, accountID uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.AuthorityProfile{}).
		Where(models.AuthorityProfileCols.AccountID+" = ? AND "+models.AuthorityProfileCols.AuthorityRoles+" @> ?",
			accountID.String(), enumx.Array[enums.AuthAuthorityRole]{enums.AuthAuthorityRoleOwner}).
		Count(&n).Error
	return n > 0, err
}

func (h *Handler) AnyOwnerExists(ctx context.Context) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.AuthorityProfile{}).
		Where(models.AuthorityProfileCols.AuthorityRoles+" @> ?",
			enumx.Array[enums.AuthAuthorityRole]{enums.AuthAuthorityRoleOwner}).
		Count(&n).Error
	return n > 0, err
}
