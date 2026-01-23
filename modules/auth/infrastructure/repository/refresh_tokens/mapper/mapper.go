package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// RefreshTokenDomainToModel 将 Domain RefreshToken 转换为 Model RefreshToken
func RefreshTokenDomainToModel(rt *refresh_tokens.RefreshToken) *models.RefreshToken {
	if rt == nil {
		return nil
	}

	var revokeReason *enums.AuthRevokeReason
	if rt.RevokeReason() != nil {
		reason := revokeReasonDomainToEnum(*rt.RevokeReason())
		revokeReason = &reason
	}

	return &models.RefreshToken{
		ID:           rt.ID(),
		TokenID:      rt.TokenID(),
		UserID:       rt.UserID(),
		AppID:        rt.AppID(),
		ClientID:     rt.ClientID(),
		SessionID:    rt.SessionID(),
		IssuedAt:     rt.IssuedAt(),
		ExpiresAt:    rt.ExpiresAt(),
		RevokedAt:    rt.RevokedAt(),
		RevokeReason: revokeReason,
		RotatedFrom:  rt.RotatedFrom(),
		DeviceID:     rt.DeviceID(),
		IP:           rt.IP(),
		UaHash:       rt.UAHash(), // Model 使用 UaHash，Domain 使用 UAHash
		CreatedAt:    rt.CreatedAt(),
		UpdatedAt:    rt.UpdatedAt(),
	}
}

// RefreshTokenModelToDomain 将 Model RefreshToken 转换为 Domain RefreshToken
func RefreshTokenModelToDomain(m *models.RefreshToken) *refresh_tokens.RefreshToken {
	if m == nil {
		return nil
	}

	var revokeReason *refresh_tokens.RevokeReason
	if m.RevokeReason != nil {
		reason := revokeReasonEnumToDomain(*m.RevokeReason)
		revokeReason = &reason
	}

	state := refresh_tokens.RefreshTokenState{
		ID:           m.ID,
		TokenID:      m.TokenID,
		UserID:       m.UserID,
		AppID:        m.AppID,
		ClientID:     m.ClientID,
		SessionID:    m.SessionID,
		IssuedAt:     m.IssuedAt,
		ExpiresAt:    m.ExpiresAt,
		RevokedAt:    m.RevokedAt,
		RevokeReason: revokeReason,
		RotatedFrom:  m.RotatedFrom,
		DeviceID:     m.DeviceID,
		IP:           m.IP,
		UAHash:       m.UaHash, // Model 使用 UaHash，Domain 使用 UAHash
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}

	return refresh_tokens.NewRefreshTokenFromState(state)
}

// RefreshTokenModelToUpdates 将 Model RefreshToken 转换为更新字段映射
func RefreshTokenModelToUpdates(m *models.RefreshToken) map[string]any {
	return map[string]any{
		models.RefreshTokenCols.UserID:       m.UserID,
		models.RefreshTokenCols.AppID:       m.AppID,
		models.RefreshTokenCols.ClientID:    m.ClientID,
		models.RefreshTokenCols.SessionID:   m.SessionID,
		models.RefreshTokenCols.IssuedAt:    m.IssuedAt,
		models.RefreshTokenCols.ExpiresAt:   m.ExpiresAt,
		models.RefreshTokenCols.RevokedAt:   m.RevokedAt,
		models.RefreshTokenCols.RevokeReason: m.RevokeReason,
		models.RefreshTokenCols.RotatedFrom: m.RotatedFrom,
		models.RefreshTokenCols.DeviceID:    m.DeviceID,
		models.RefreshTokenCols.IP:          m.IP,
		models.RefreshTokenCols.UaHash:      m.UaHash,
		models.RefreshTokenCols.UpdatedAt:   m.UpdatedAt,
	}
}

// 枚举转换辅助函数

// RevokeReasonDomainToEnum 将 Domain RevokeReason 转换为 Enum RevokeReason
func RevokeReasonDomainToEnum(rr refresh_tokens.RevokeReason) enums.AuthRevokeReason {
	return revokeReasonDomainToEnum(rr)
}

func revokeReasonDomainToEnum(rr refresh_tokens.RevokeReason) enums.AuthRevokeReason {
	switch rr {
	case refresh_tokens.RevokeReasonUserLogout:
		return enums.AuthRevokeReasonUserLogout
	case refresh_tokens.RevokeReasonAdminRevoke:
		return enums.AuthRevokeReasonAdminRevoke
	case refresh_tokens.RevokeReasonPasswordChanged:
		return enums.AuthRevokeReasonPasswordChanged
	case refresh_tokens.RevokeReasonRotation:
		return enums.AuthRevokeReasonRotation
	case refresh_tokens.RevokeReasonAccountLocked:
		return enums.AuthRevokeReasonAccountLocked
	case refresh_tokens.RevokeReasonDeviceChanged:
		return enums.AuthRevokeReasonDeviceChanged
	case refresh_tokens.RevokeReasonSuspiciousActivity:
		return enums.AuthRevokeReasonSuspiciousActivity
	case refresh_tokens.RevokeReasonOther:
		return enums.AuthRevokeReasonOther
	default:
		return enums.AuthRevokeReasonOther
	}
}

func revokeReasonEnumToDomain(rr enums.AuthRevokeReason) refresh_tokens.RevokeReason {
	switch rr {
	case enums.AuthRevokeReasonUserLogout:
		return refresh_tokens.RevokeReasonUserLogout
	case enums.AuthRevokeReasonAdminRevoke:
		return refresh_tokens.RevokeReasonAdminRevoke
	case enums.AuthRevokeReasonPasswordChanged:
		return refresh_tokens.RevokeReasonPasswordChanged
	case enums.AuthRevokeReasonRotation:
		return refresh_tokens.RevokeReasonRotation
	case enums.AuthRevokeReasonAccountLocked:
		return refresh_tokens.RevokeReasonAccountLocked
	case enums.AuthRevokeReasonDeviceChanged:
		return refresh_tokens.RevokeReasonDeviceChanged
	case enums.AuthRevokeReasonSuspiciousActivity:
		return refresh_tokens.RevokeReasonSuspiciousActivity
	case enums.AuthRevokeReasonOther:
		return refresh_tokens.RevokeReasonOther
	default:
		return refresh_tokens.RevokeReasonOther
	}
}
