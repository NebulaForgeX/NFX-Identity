package create

import (
	"context"
	education "nfxid/modules/auth/domain/profile_education"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 Education，实现 education.Create 接口
func (h *Handler) New(ctx context.Context, e *education.Education) error {
	m := mapper.EducationDomainToModel(e)
	return h.db.WithContext(ctx).Create(&m).Error
}
