package mapper

import (
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// PasswordHistoryDomainToModel 将 Domain PasswordHistory 转换为 Model PasswordHistory
func PasswordHistoryDomainToModel(ph *password_history.PasswordHistory) *models.PasswordHistory {
	if ph == nil {
		return nil
	}

	return &models.PasswordHistory{
		ID:           ph.ID(),
		UserID:       ph.UserID(),
		TenantID:     ph.TenantID(),
		PasswordHash: ph.PasswordHash(),
		HashAlg:      ph.HashAlg(),
		CreatedAt:    ph.CreatedAt(),
	}
}

// PasswordHistoryModelToDomain 将 Model PasswordHistory 转换为 Domain PasswordHistory
func PasswordHistoryModelToDomain(m *models.PasswordHistory) *password_history.PasswordHistory {
	if m == nil {
		return nil
	}

	state := password_history.PasswordHistoryState{
		ID:           m.ID,
		UserID:       m.UserID,
		TenantID:     m.TenantID,
		PasswordHash: m.PasswordHash,
		HashAlg:      m.HashAlg,
		CreatedAt:    m.CreatedAt,
	}

	return password_history.NewPasswordHistoryFromState(state)
}

// PasswordHistoryModelToUpdates 将 Model PasswordHistory 转换为更新字段映射
func PasswordHistoryModelToUpdates(m *models.PasswordHistory) map[string]any {
	return map[string]any{
		models.PasswordHistoryCols.UserID:       m.UserID,
		models.PasswordHistoryCols.TenantID:     m.TenantID,
		models.PasswordHistoryCols.PasswordHash: m.PasswordHash,
		models.PasswordHistoryCols.HashAlg:      m.HashAlg,
	}
}
