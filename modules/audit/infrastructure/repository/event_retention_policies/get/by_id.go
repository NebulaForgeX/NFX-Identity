package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 EventRetentionPolicy，实现 event_retention_policies.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*event_retention_policies.EventRetentionPolicy, error) {
	var m models.EventRetentionPolicy
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, event_retention_policies.ErrEventRetentionPolicyNotFound
		}
		return nil, err
	}
	return mapper.EventRetentionPolicyModelToDomain(&m), nil
}
