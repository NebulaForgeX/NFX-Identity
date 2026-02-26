package get

import (
	"context"
	"errors"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/member_groups"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/member_groups/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 MemberGroup，实现 member_groups.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*member_groups.MemberGroup, error) {
	var m models.MemberGroup
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrMemberGroupNotFound
		}
		return nil, err
	}
	return mapper.MemberGroupModelToDomain(&m), nil
}
