package update

import (
	"context"
	"nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/mapper"
	"time"

	"github.com/google/uuid"
)

// UpdateStatus 更新状态，实现 user_credentials.Update 接口
func (h *Handler) UpdateStatus(ctx context.Context, userID uuid.UUID, status user_credentials.CredentialStatus) error {
	statusEnum := mapper.CredentialStatusDomainToEnum(status)
	updates := map[string]any{
		models.UserCredentialCols.Status:    statusEnum,
		models.UserCredentialCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.UserCredential{}).
		Where("id = ?", userID).
		Updates(updates).Error
}
