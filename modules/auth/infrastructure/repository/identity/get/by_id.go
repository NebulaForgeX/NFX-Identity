package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	identityDomain "nfxidentity/modules/auth/domain/identity"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/identity/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*identityDomain.Identity, error) {
	var m rdbmodels.Identity
	if err := h.db.WithContext(ctx).Where(rdbmodels.IdentityCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrIdentityNotFound
		}
		return nil, err
	}
	return mapper.IdentityModelToDomain(&m), nil
}
