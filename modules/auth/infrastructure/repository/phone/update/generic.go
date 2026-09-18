package update

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"
)

func (h *Handler) Generic(ctx context.Context, p *phoneDomain.Phone) error {
	updates := mapper.PhoneDomainToUpdates(p)
	res := h.db.WithContext(ctx).Model(&rdbmodels.Phone{}).
		Where(rdbmodels.PhoneCols.ID+" = ?", p.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrPhoneBindingNotFound
	}
	return nil
}
