package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 MemberGroup，实现 member_groups.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*member_groups.MemberGroup, error) {
	var m models.MemberGroup
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, member_groups.ErrMemberGroupNotFound
		}
		return nil, err
	}
	return mapper.MemberGroupModelToDomain(&m), nil
}
