package get

import (
	"context"

	"nfxid/modules/access/domain/super_admins"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/super_admins/mapper"
)

func (h *Handler) All(ctx context.Context, limit, offset int) ([]*super_admins.SuperAdmin, error) {
	var list []*models.SuperAdmin
	q := h.db.WithContext(ctx).Order("created_at ASC")
	if limit > 0 {
		q = q.Limit(limit)
	}
	if offset > 0 {
		q = q.Offset(offset)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*super_admins.SuperAdmin, len(list))
	for i := range list {
		out[i] = mapper.SuperAdminModelToDomain(list[i])
	}
	return out, nil
}
