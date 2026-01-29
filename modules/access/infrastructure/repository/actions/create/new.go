package create

import (
	"context"

	"nfxid/modules/access/domain/actions"
	"nfxid/modules/access/infrastructure/repository/actions/mapper"
)

func (h *Handler) New(ctx context.Context, a *actions.Action) error {
	m := mapper.ActionDomainToModel(a)
	return h.db.WithContext(ctx).Create(m).Error
}
