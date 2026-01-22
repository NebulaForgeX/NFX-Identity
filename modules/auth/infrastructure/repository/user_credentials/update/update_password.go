package update

import (
	"context"
	"encoding/json"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// UpdatePassword 更新密码，实现 user_credentials.Update 接口
func (h *Handler) UpdatePassword(ctx context.Context, userID uuid.UUID, passwordHash string, hashAlg string, hashParams map[string]interface{}) error {
	now := time.Now().UTC()

	var hashParamsJSON *datatypes.JSON
	if len(hashParams) > 0 {
		paramsBytes, _ := json.Marshal(hashParams)
		jsonData := datatypes.JSON(paramsBytes)
		hashParamsJSON = &jsonData
	}

	updates := map[string]any{
		models.UserCredentialCols.PasswordHash:      passwordHash,
		models.UserCredentialCols.HashAlg:           hashAlg,
		models.UserCredentialCols.HashParams:        hashParamsJSON,
		models.UserCredentialCols.PasswordUpdatedAt: &now,
		models.UserCredentialCols.UpdatedAt:         now,
	}

	return h.db.WithContext(ctx).
		Model(&models.UserCredential{}).
		Where("id = ?", userID).
		Updates(updates).Error
}
