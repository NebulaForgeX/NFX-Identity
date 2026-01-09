package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 MemberGroup，实现 member_groups.Update 接口
func (h *Handler) Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID) error {
	// 先检查 MemberGroup 是否存在
	var m models.MemberGroup
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return member_groups.ErrMemberGroupNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return member_groups.ErrMemberGroupAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.MemberGroupCols.RevokedAt: &now,
		models.MemberGroupCols.RevokedBy: &revokedBy,
	}

	return h.db.WithContext(ctx).
		Model(&models.MemberGroup{}).
		Where("id = ?", id).
		Updates(updates).Error
}
