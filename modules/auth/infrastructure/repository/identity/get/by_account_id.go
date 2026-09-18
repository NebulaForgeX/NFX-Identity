package get

import (
	"context"

	identityDomain "nfxidentity/modules/auth/domain/identity"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/identity/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*identityDomain.Identity, error) {
	var ms []rdbmodels.Identity
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.IdentityCols.AccountID+" = ?", accountID.String()).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	out := make([]*identityDomain.Identity, 0, len(ms))
	for i := range ms {
		out = append(out, mapper.IdentityModelToDomain(&ms[i]))
	}
	return out, nil
}
