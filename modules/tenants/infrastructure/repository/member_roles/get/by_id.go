package get

import (
	"context"
	"errors"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/member_roles"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 MemberRole，实现 member_roles.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*member_roles.MemberRole, error) {
	var m models.MemberRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrMemberRoleNotFound
		}
		return nil, err
	}
	return mapper.MemberRoleModelToDomain(&m), nil
}
