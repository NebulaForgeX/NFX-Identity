package mapper

import (
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ClientScopeDomainToModel 将 Domain ClientScope 转换为 Model ClientScope
func ClientScopeDomainToModel(cs *client_scopes.ClientScope) *models.ClientScope {
	if cs == nil {
		return nil
	}

	return &models.ClientScope{
		ID:             cs.ID(),
		ApplicationID:  cs.AppID(),
		Scope:        cs.Scope(),
		GrantedBy:    cs.GrantedBy(),
		GrantedAt:    cs.GrantedAt(),
		ExpiresAt:    cs.ExpiresAt(),
		CreatedAt:    cs.CreatedAt(),
		RevokedAt:    cs.RevokedAt(),
		RevokedBy:    cs.RevokedBy(),
		RevokeReason: cs.RevokeReason(),
	}
}

// ClientScopeModelToDomain 将 Model ClientScope 转换为 Domain ClientScope
func ClientScopeModelToDomain(m *models.ClientScope) *client_scopes.ClientScope {
	if m == nil {
		return nil
	}

	state := client_scopes.ClientScopeState{
		ID:           m.ID,
		AppID:        m.ApplicationID,
		Scope:        m.Scope,
		GrantedBy:    m.GrantedBy,
		GrantedAt:    m.GrantedAt,
		ExpiresAt:    m.ExpiresAt,
		CreatedAt:    m.CreatedAt,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
	}

	return client_scopes.NewClientScopeFromState(state)
}

// ClientScopeModelToUpdates 将 Model ClientScope 转换为更新字段映射
func ClientScopeModelToUpdates(m *models.ClientScope) map[string]any {
	return map[string]any{
		models.ClientScopeCols.ExpiresAt:    m.ExpiresAt,
		models.ClientScopeCols.RevokedAt:    m.RevokedAt,
		models.ClientScopeCols.RevokedBy:    m.RevokedBy,
		models.ClientScopeCols.RevokeReason: m.RevokeReason,
	}
}
