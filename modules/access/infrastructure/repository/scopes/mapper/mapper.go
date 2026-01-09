package mapper

import (
	"nfxid/modules/access/domain/scopes"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// ScopeDomainToModel 将 Domain Scope 转换为 Model Scope
func ScopeDomainToModel(s *scopes.Scope) *models.Scope {
	if s == nil {
		return nil
	}

	return &models.Scope{
		Scope:       s.ScopeKey(),
		Description: s.Description(),
		IsSystem:    s.IsSystem(),
		CreatedAt:   s.CreatedAt(),
		UpdatedAt:   s.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(s.DeletedAt()),
	}
}

// ScopeModelToDomain 将 Model Scope 转换为 Domain Scope
func ScopeModelToDomain(m *models.Scope) *scopes.Scope {
	if m == nil {
		return nil
	}

	state := scopes.ScopeState{
		Scope:       m.Scope,
		Description: m.Description,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return scopes.NewScopeFromState(state)
}

// ScopeModelToUpdates 将 Model Scope 转换为更新字段映射
func ScopeModelToUpdates(m *models.Scope) map[string]any {
	return map[string]any{
		models.ScopeCols.Description: m.Description,
		models.ScopeCols.IsSystem:    m.IsSystem,
		models.ScopeCols.UpdatedAt:   m.UpdatedAt,
		models.ScopeCols.DeletedAt:   m.DeletedAt,
	}
}
