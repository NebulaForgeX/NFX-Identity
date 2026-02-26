package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/members"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/members/mapper"
	"time"

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
