package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/members/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Member，实现 members.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*members.Member, error) {
	var m models.Member
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, members.ErrMemberNotFound
		}
		return nil, err
	}
	return mapper.MemberModelToDomain(&m), nil
}
