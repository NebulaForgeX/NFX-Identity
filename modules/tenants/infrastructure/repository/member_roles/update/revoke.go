package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 MemberRole，实现 member_roles.Update 接口
func (h *Handler) Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error {
	// 先检查 MemberRole 是否存在
	var m models.MemberRole
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return member_roles.ErrMemberRoleNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return member_roles.ErrMemberRoleAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.MemberRoleCols.RevokedAt:    &now,
		models.MemberRoleCols.RevokedBy:    &revokedBy,
		models.MemberRoleCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.MemberRole{}).
		Where("id = ?", id).
		Updates(updates).Error
}
