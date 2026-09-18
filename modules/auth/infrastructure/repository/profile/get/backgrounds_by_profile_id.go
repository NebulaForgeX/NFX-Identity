package get

import (
	"context"
	"github.com/google/uuid"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority backgrounds_by_profile_id =============================== !//
func (h *Handler) AuthorityBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*profileDomain.AuthorityProfileBackground, error) {
	var rows []rdbmodels.AuthorityProfileBackground
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileBackgroundCols.ProfileID+" = ?", profileID.String()).
		Order(rdbmodels.AuthorityProfileBackgroundCols.SortOrder + " ASC, " + rdbmodels.AuthorityProfileBackgroundCols.CreatedAt + " ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*profileDomain.AuthorityProfileBackground, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.AuthorityProfileBackgroundModelToDomain(&rows[i]))
	}
	return out, nil
}

//* =============================== Forger backgrounds_by_profile_id =============================== !//
func (h *Handler) ForgerBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*profileDomain.ForgerProfileBackground, error) {
	var rows []rdbmodels.ForgerProfileBackground
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileBackgroundCols.ProfileID+" = ?", profileID.String()).
		Order(rdbmodels.ForgerProfileBackgroundCols.SortOrder + " ASC, " + rdbmodels.ForgerProfileBackgroundCols.CreatedAt + " ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*profileDomain.ForgerProfileBackground, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ForgerProfileBackgroundModelToDomain(&rows[i]))
	}
	return out, nil
}
