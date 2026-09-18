package get

import (
	"context"
	"errors"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	identityDomain "nfxidentity/modules/auth/domain/identity"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/identity/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByProviderSubject(ctx context.Context, provider enums.AuthIdentityProvider, subject string) (*identityDomain.Identity, error) {
	var m rdbmodels.Identity
	if err := h.db.WithContext(ctx).
		Where(
			rdbmodels.IdentityCols.IdentityProvider+" = ? AND "+rdbmodels.IdentityCols.ProviderSubject+" = ?",
			provider, subject,
		).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrIdentityNotFound
		}
		return nil, err
	}
	return mapper.IdentityModelToDomain(&m), nil
}
