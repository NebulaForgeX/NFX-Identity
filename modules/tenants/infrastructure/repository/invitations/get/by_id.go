package get

import (
	"context"
	"errors"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/invitations"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Invitation，实现 invitations.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*invitations.Invitation, error) {
	var m models.Invitation
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrInvitationNotFound
		}
		return nil, err
	}
	return mapper.InvitationModelToDomain(&m), nil
}
