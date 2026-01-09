package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// APIKeyDomainToModel 将 Domain APIKey 转换为 Model ApiKey
func APIKeyDomainToModel(ak *api_keys.APIKey) *models.ApiKey {
	if ak == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if ak.Metadata() != nil {
		metadataBytes, _ := json.Marshal(ak.Metadata())
		metadata = (*datatypes.JSON)(&metadataBytes)
	}

	return &models.ApiKey{
		ID:           ak.ID(),
		KeyID:        ak.KeyID(),
		AppID:        ak.AppID(),
		KeyHash:      ak.KeyHash(),
		HashAlg:      ak.HashAlg(),
		Name:         ak.Name(),
		Status:       apiKeyStatusDomainToEnum(ak.Status()),
		ExpiresAt:    ak.ExpiresAt(),
		CreatedAt:    ak.CreatedAt(),
		RevokedAt:    ak.RevokedAt(),
		RevokedBy:    ak.RevokedBy(),
		RevokeReason: ak.RevokeReason(),
		LastUsedAt:   ak.LastUsedAt(),
		CreatedBy:    ak.CreatedBy(),
		Metadata:     metadata,
	}
}

// APIKeyModelToDomain 将 Model ApiKey 转换为 Domain APIKey
func APIKeyModelToDomain(m *models.ApiKey) *api_keys.APIKey {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := api_keys.APIKeyState{
		ID:           m.ID,
		KeyID:        m.KeyID,
		AppID:        m.AppID,
		KeyHash:      m.KeyHash,
		HashAlg:      m.HashAlg,
		Name:         m.Name,
		Status:       apiKeyStatusEnumToDomain(m.Status),
		ExpiresAt:    m.ExpiresAt,
		CreatedAt:    m.CreatedAt,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
		LastUsedAt:   m.LastUsedAt,
		CreatedBy:    m.CreatedBy,
		Metadata:     metadata,
	}

	return api_keys.NewAPIKeyFromState(state)
}

// APIKeyModelToUpdates 将 Model ApiKey 转换为更新字段映射
func APIKeyModelToUpdates(m *models.ApiKey) map[string]any {
	return map[string]any{
		models.ApiKeyCols.KeyHash:    m.KeyHash,
		models.ApiKeyCols.HashAlg:    m.HashAlg,
		models.ApiKeyCols.Name:       m.Name,
		models.ApiKeyCols.Status:     m.Status,
		models.ApiKeyCols.ExpiresAt:  m.ExpiresAt,
		models.ApiKeyCols.RevokedAt:  m.RevokedAt,
		models.ApiKeyCols.RevokedBy:  m.RevokedBy,
		models.ApiKeyCols.RevokeReason: m.RevokeReason,
		models.ApiKeyCols.LastUsedAt: m.LastUsedAt,
		models.ApiKeyCols.Metadata:   m.Metadata,
	}
}

// 枚举转换辅助函数

func apiKeyStatusDomainToEnum(aks api_keys.APIKeyStatus) enums.ClientsApiKeyStatus {
	switch aks {
	case api_keys.APIKeyStatusActive:
		return enums.ClientsApiKeyStatusActive
	case api_keys.APIKeyStatusRevoked:
		return enums.ClientsApiKeyStatusRevoked
	case api_keys.APIKeyStatusExpired:
		return enums.ClientsApiKeyStatusExpired
	default:
		return enums.ClientsApiKeyStatusActive
	}
}

func apiKeyStatusEnumToDomain(aks enums.ClientsApiKeyStatus) api_keys.APIKeyStatus {
	switch aks {
	case enums.ClientsApiKeyStatusActive:
		return api_keys.APIKeyStatusActive
	case enums.ClientsApiKeyStatusRevoked:
		return api_keys.APIKeyStatusRevoked
	case enums.ClientsApiKeyStatusExpired:
		return api_keys.APIKeyStatusExpired
	default:
		return api_keys.APIKeyStatusActive
	}
}
