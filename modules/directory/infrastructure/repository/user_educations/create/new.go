package create

import (
	"context"
	"nfxid/modules/directory/domain/user_educations"
	"nfxid/modules/directory/infrastructure/repository/user_educations/mapper"
)

// New 创建新的 UserEducation，实现 user_educations.Create 接口
func (h *Handler) New(ctx context.Context, ue *user_educations.UserEducation) error {
	m := mapper.UserEducationDomainToModel(ue)
	return h.db.WithContext(ctx).Create(&m).Error
}
