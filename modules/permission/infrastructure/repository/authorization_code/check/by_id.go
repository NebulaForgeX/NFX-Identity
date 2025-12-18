package check

import (
	"context"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 检查 AuthorizationCode 是否存在，实现 authorizationCodeDomain.Check 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.AuthorizationCode{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
