package update

import (
	"context"
	"errors"
	tenantsErr "nfxid/errors/src/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 MemberAppRole，实现 member_app_roles.Update 接口
func (h *Handler) Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error {
	// 先检查 MemberAppRole 是否存在
	var m models.MemberAppRole
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tenantsErr.ErrMemberAppRoleNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return tenantsErr.ErrMemberAppRoleAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.MemberAppRoleCols.RevokedAt:    &now,
		models.MemberAppRoleCols.RevokedBy:    &revokedBy,
		models.MemberAppRoleCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.MemberAppRole{}).
		Where("id = ?", id).
		Updates(updates).Error
}
