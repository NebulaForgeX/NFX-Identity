package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/invitations/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Invitation，实现 invitations.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*invitations.Invitation, error) {
	var m models.Invitation
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, invitations.ErrInvitationNotFound
		}
		return nil, err
	}
	return mapper.InvitationModelToDomain(&m), nil
}
