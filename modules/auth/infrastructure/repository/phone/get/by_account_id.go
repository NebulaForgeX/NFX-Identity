package get

import (
	"context"

	phoneDomain "nfxidentity/modules/auth/domain/phone"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*phoneDomain.Phone, error) {
	var ms []rdbmodels.Phone
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.PhoneCols.AccountID+" = ?", accountID.String()).
		Find(&ms).Error; err != nil {
		return nil, err
	}
	out := make([]*phoneDomain.Phone, 0, len(ms))
	for i := range ms {
		out = append(out, mapper.PhoneModelToDomain(&ms[i]))
	}
	return out, nil
}
