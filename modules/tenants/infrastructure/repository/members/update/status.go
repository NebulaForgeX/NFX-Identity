package update

import (
	"context"
	"time"
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/members/mapper"

	"github.com/google/uuid"
)

// Status 更新状态，实现 members.Update 接口
func (h *Handler) Status(ctx context.Context, memberID uuid.UUID, status members.MemberStatus) error {
	statusEnum := mapper.MemberStatusDomainToEnum(status)
	updates := map[string]any{
		models.MemberCols.Status:    statusEnum,
		models.MemberCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.Member{}).
		Where("member_id = ?", memberID).
		Updates(updates).Error
}
