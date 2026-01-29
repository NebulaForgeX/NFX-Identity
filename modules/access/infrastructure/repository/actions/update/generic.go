package update

import (
	"context"

	"nfxid/modules/access/domain/actions"
	"nfxid/modules/access/infrastructure/repository/actions/mapper"
)

func (h *Handler) Generic(ctx context.Context, a *actions.Action) error {
	m := mapper.ActionDomainToModel(a)
	updates := mapper.ActionModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&mapper.ActionModel{}).
		Where("id = ?", a.ID()).
		Updates(updates).Error
}
