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

// ByMemberIDAndGroupID 根据 MemberID 和 GroupID 获取 MemberGroup，实现 member_groups.Get 接口
func (h *Handler) ByMemberIDAndGroupID(ctx context.Context, memberID, groupID uuid.UUID) (*member_groups.MemberGroup, error) {
	var m models.MemberGroup
	if err := h.db.WithContext(ctx).
		Where("member_id = ? AND group_id = ?", memberID, groupID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, member_groups.ErrMemberGroupNotFound
		}
		return nil, err
	}
	return mapper.MemberGroupModelToDomain(&m), nil
}
