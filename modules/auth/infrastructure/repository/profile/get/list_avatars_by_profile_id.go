package get

import (
	"context"
	"github.com/google/uuid"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority list_avatars_by_profile_id =============================== !//
func (h *Handler) ListAuthorityAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*profileDomain.AuthorityProfileAvatar, error) {
	var rows []rdbmodels.AuthorityProfileAvatar
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileAvatarCols.ProfileID+" = ?", profileID.String()).
		Where(rdbmodels.AuthorityProfileAvatarCols.DeletedAt + " IS NULL"). //nolint:goconst // GORM column name
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*profileDomain.AuthorityProfileAvatar, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.AuthorityProfileAvatarModelToDomain(&rows[i]))
	}
	return out, nil
}

//* =============================== Forger list_avatars_by_profile_id =============================== !//
func (h *Handler) ListForgerAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*profileDomain.ForgerProfileAvatar, error) {
	var rows []rdbmodels.ForgerProfileAvatar
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileAvatarCols.ProfileID+" = ?", profileID.String()).
		Where(rdbmodels.ForgerProfileAvatarCols.DeletedAt + " IS NULL"). //nolint:goconst // GORM column name
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*profileDomain.ForgerProfileAvatar, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ForgerProfileAvatarModelToDomain(&rows[i]))
	}
	return out, nil
}
