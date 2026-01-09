package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// PasswordResetDomainToModel 将 Domain PasswordReset 转换为 Model PasswordReset
func PasswordResetDomainToModel(pr *password_resets.PasswordReset) *models.PasswordReset {
	if pr == nil {
		return nil
	}

	return &models.PasswordReset{
		ID:           pr.ID(),
		ResetID:      pr.ResetID(),
		TenantID:     pr.TenantID(),
		UserID:       pr.UserID(),
		Delivery:     resetDeliveryDomainToEnum(pr.Delivery()),
		CodeHash:     pr.CodeHash(),
		ExpiresAt:    pr.ExpiresAt(),
		UsedAt:       pr.UsedAt(),
		RequestedIP:  pr.RequestedIP(),
		UaHash:       pr.UAHash(), // Model 使用 UaHash，Domain 使用 UAHash
		AttemptCount: pr.AttemptCount(),
		Status:       resetStatusDomainToEnum(pr.Status()),
		CreatedAt:    pr.CreatedAt(),
		UpdatedAt:    pr.UpdatedAt(),
	}
}

// PasswordResetModelToDomain 将 Model PasswordReset 转换为 Domain PasswordReset
func PasswordResetModelToDomain(m *models.PasswordReset) *password_resets.PasswordReset {
	if m == nil {
		return nil
	}

	state := password_resets.PasswordResetState{
		ID:           m.ID,
		ResetID:      m.ResetID,
		TenantID:     m.TenantID,
		UserID:       m.UserID,
		Delivery:     resetDeliveryEnumToDomain(m.Delivery),
		CodeHash:     m.CodeHash,
		ExpiresAt:    m.ExpiresAt,
		UsedAt:       m.UsedAt,
		RequestedIP:  m.RequestedIP,
		UAHash:       m.UaHash, // Model 使用 UaHash，Domain 使用 UAHash
		AttemptCount: m.AttemptCount,
		Status:       resetStatusEnumToDomain(m.Status),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}

	return password_resets.NewPasswordResetFromState(state)
}

// PasswordResetModelToUpdates 将 Model PasswordReset 转换为更新字段映射
func PasswordResetModelToUpdates(m *models.PasswordReset) map[string]any {
	return map[string]any{
		models.PasswordResetCols.TenantID:     m.TenantID,
		models.PasswordResetCols.UserID:      m.UserID,
		models.PasswordResetCols.Delivery:    m.Delivery,
		models.PasswordResetCols.CodeHash:    m.CodeHash,
		models.PasswordResetCols.ExpiresAt:   m.ExpiresAt,
		models.PasswordResetCols.UsedAt:     m.UsedAt,
		models.PasswordResetCols.RequestedIP: m.RequestedIP,
		models.PasswordResetCols.UaHash:     m.UaHash,
		models.PasswordResetCols.AttemptCount: m.AttemptCount,
		models.PasswordResetCols.Status:      m.Status,
		models.PasswordResetCols.UpdatedAt:   m.UpdatedAt,
	}
}

// 枚举转换辅助函数

func resetDeliveryDomainToEnum(rd password_resets.ResetDelivery) enums.AuthResetDelivery {
	switch rd {
	case password_resets.ResetDeliveryEmail:
		return enums.AuthResetDeliveryEmail
	case password_resets.ResetDeliverySMS:
		return enums.AuthResetDeliverySms
	default:
		return enums.AuthResetDeliveryEmail
	}
}

func resetDeliveryEnumToDomain(rd enums.AuthResetDelivery) password_resets.ResetDelivery {
	switch rd {
	case enums.AuthResetDeliveryEmail:
		return password_resets.ResetDeliveryEmail
	case enums.AuthResetDeliverySms:
		return password_resets.ResetDeliverySMS
	default:
		return password_resets.ResetDeliveryEmail
	}
}

// ResetStatusDomainToEnum 将 Domain ResetStatus 转换为 Enum ResetStatus
func ResetStatusDomainToEnum(rs password_resets.ResetStatus) enums.AuthResetStatus {
	return resetStatusDomainToEnum(rs)
}

func resetStatusDomainToEnum(rs password_resets.ResetStatus) enums.AuthResetStatus {
	switch rs {
	case password_resets.ResetStatusIssued:
		return enums.AuthResetStatusIssued
	case password_resets.ResetStatusUsed:
		return enums.AuthResetStatusUsed
	case password_resets.ResetStatusExpired:
		return enums.AuthResetStatusExpired
	case password_resets.ResetStatusRevoked:
		return enums.AuthResetStatusRevoked
	default:
		return enums.AuthResetStatusIssued
	}
}

func resetStatusEnumToDomain(rs enums.AuthResetStatus) password_resets.ResetStatus {
	switch rs {
	case enums.AuthResetStatusIssued:
		return password_resets.ResetStatusIssued
	case enums.AuthResetStatusUsed:
		return password_resets.ResetStatusUsed
	case enums.AuthResetStatusExpired:
		return password_resets.ResetStatusExpired
	case enums.AuthResetStatusRevoked:
		return password_resets.ResetStatusRevoked
	default:
		return password_resets.ResetStatusIssued
	}
}
