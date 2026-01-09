package update

import (
	"context"
	"time"
	"nfxid/enums"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Join 加入租户，实现 members.Update 接口
func (h *Handler) Join(ctx context.Context, memberID uuid.UUID) error {
	now := time.Now().UTC()
	status := enums.TenantsMemberStatusActive
	updates := map[string]any{
		models.MemberCols.Status:   status,
		models.MemberCols.JoinedAt: &now,
		models.MemberCols.UpdatedAt: now,
	}

	return h.db.WithContext(ctx).
		Model(&models.Member{}).
		Where("member_id = ?", memberID).
		Updates(updates).Error
}
