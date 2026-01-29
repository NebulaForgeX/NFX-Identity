package create

import (
	"context"

	"nfxid/modules/access/domain/action_requirements"
	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"
)

func (h *Handler) New(ctx context.Context, ar *action_requirements.ActionRequirement) error {
	m := mapper.ActionRequirementDomainToModel(ar)
	return h.db.WithContext(ctx).Create(m).Error
}
