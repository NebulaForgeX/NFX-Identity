package get

import (
	"context"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/grants/mapper"

	"github.com/google/uuid"
)

// BySubject 根据 SubjectType 和 SubjectID 获取 Grant 列表，实现 grants.Get 接口
func (h *Handler) BySubject(ctx context.Context, subjectType grants.SubjectType, subjectID uuid.UUID) ([]*grants.Grant, error) {
	var ms []models.Grant
	if err := h.db.WithContext(ctx).
		Where("subject_type = ? AND subject_id = ?", subjectType, subjectID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*grants.Grant, len(ms))
	for i, m := range ms {
		result[i] = mapper.GrantModelToDomain(&m)
	}
	return result, nil
}
