package get

import (
	"context"
	"errors"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/sessions"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Session，实现 sessions.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*sessions.Session, error) {
	var m models.Session
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrSessionNotFound
		}
		return nil, err
	}
	return mapper.SessionModelToDomain(&m), nil
}
