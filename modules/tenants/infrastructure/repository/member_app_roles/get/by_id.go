package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/member_app_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_app_roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 MemberAppRole，实现 member_app_roles.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*member_app_roles.MemberAppRole, error) {
	var m models.MemberAppRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, member_app_roles.ErrMemberAppRoleNotFound
		}
		return nil, err
	}
	return mapper.MemberAppRoleModelToDomain(&m), nil
}
