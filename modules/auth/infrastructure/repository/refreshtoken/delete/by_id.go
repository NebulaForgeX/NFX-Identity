package delete

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	res := h.db.WithContext(ctx).Where(rdbmodels.RefreshTokenCols.ID+" = ?", id.String()).Delete(&rdbmodels.RefreshToken{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrRefreshTokenNotFound
	}
	return nil
}
