package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/scopes"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/scopes/mapper"

	"gorm.io/gorm"
)

// ByScope 根据 Scope 获取 Scope，实现 scopes.Get 接口
func (h *Handler) ByScope(ctx context.Context, scope string) (*scopes.Scope, error) {
	var m models.Scope
	if err := h.db.WithContext(ctx).Where("scope = ?", scope).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, scopes.ErrScopeNotFound
		}
		return nil, err
	}
	return mapper.ScopeModelToDomain(&m), nil
}
