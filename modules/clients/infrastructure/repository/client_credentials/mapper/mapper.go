package mapper

import (
	"nfxid/enums"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ClientCredentialDomainToModel 将 Domain ClientCredential 转换为 Model ClientCredential
func ClientCredentialDomainToModel(cc *client_credentials.ClientCredential) *models.ClientCredential {
	if cc == nil {
		return nil
	}

	return &models.ClientCredential{
		ID:           cc.ID(),
		ApplicationID: cc.AppID(),
		ClientID:     cc.ClientID(),
		SecretHash:   cc.SecretHash(),
		HashAlg:      cc.HashAlg(),
		Status:       credentialStatusDomainToEnum(cc.Status()),
		CreatedAt:    cc.CreatedAt(),
		RotatedAt:    cc.RotatedAt(),
		ExpiresAt:    cc.ExpiresAt(),
		LastUsedAt:   cc.LastUsedAt(),
		CreatedBy:    cc.CreatedBy(),
		RevokedAt:    cc.RevokedAt(),
		RevokedBy:    cc.RevokedBy(),
		RevokeReason: cc.RevokeReason(),
	}
}

// ClientCredentialModelToDomain 将 Model ClientCredential 转换为 Domain ClientCredential
func ClientCredentialModelToDomain(m *models.ClientCredential) *client_credentials.ClientCredential {
	if m == nil {
		return nil
	}

	state := client_credentials.ClientCredentialState{
		ID:           m.ID,
		AppID:        m.ApplicationID,
		ClientID:     m.ClientID,
		SecretHash:   m.SecretHash,
		HashAlg:      m.HashAlg,
		Status:       credentialStatusEnumToDomain(m.Status),
		CreatedAt:    m.CreatedAt,
		RotatedAt:    m.RotatedAt,
		ExpiresAt:    m.ExpiresAt,
		LastUsedAt:   m.LastUsedAt,
		CreatedBy:    m.CreatedBy,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
	}

	return client_credentials.NewClientCredentialFromState(state)
}

// ClientCredentialModelToUpdates 将 Model ClientCredential 转换为更新字段映射
func ClientCredentialModelToUpdates(m *models.ClientCredential) map[string]any {
	return map[string]any{
		models.ClientCredentialCols.SecretHash:   m.SecretHash,
		models.ClientCredentialCols.HashAlg:      m.HashAlg,
		models.ClientCredentialCols.Status:      m.Status,
		models.ClientCredentialCols.RotatedAt:    m.RotatedAt,
		models.ClientCredentialCols.ExpiresAt:    m.ExpiresAt,
		models.ClientCredentialCols.LastUsedAt:   m.LastUsedAt,
		models.ClientCredentialCols.RevokedAt:     m.RevokedAt,
		models.ClientCredentialCols.RevokedBy:     m.RevokedBy,
		models.ClientCredentialCols.RevokeReason: m.RevokeReason,
	}
}

// 枚举转换辅助函数

func credentialStatusDomainToEnum(cs client_credentials.CredentialStatus) enums.ClientsCredentialStatus {
	switch cs {
	case client_credentials.CredentialStatusActive:
		return enums.ClientsCredentialStatusActive
	case client_credentials.CredentialStatusExpired:
		return enums.ClientsCredentialStatusExpired
	case client_credentials.CredentialStatusRevoked:
		return enums.ClientsCredentialStatusRevoked
	case client_credentials.CredentialStatusRotating:
		return enums.ClientsCredentialStatusRotating
	default:
		return enums.ClientsCredentialStatusActive
	}
}

func credentialStatusEnumToDomain(cs enums.ClientsCredentialStatus) client_credentials.CredentialStatus {
	switch cs {
	case enums.ClientsCredentialStatusActive:
		return client_credentials.CredentialStatusActive
	case enums.ClientsCredentialStatusExpired:
		return client_credentials.CredentialStatusExpired
	case enums.ClientsCredentialStatusRevoked:
		return client_credentials.CredentialStatusRevoked
	case enums.ClientsCredentialStatusRotating:
		return client_credentials.CredentialStatusRotating
	default:
		return client_credentials.CredentialStatusActive
	}
}
