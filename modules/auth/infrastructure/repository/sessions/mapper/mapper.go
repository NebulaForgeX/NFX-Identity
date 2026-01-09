package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// SessionDomainToModel 将 Domain Session 转换为 Model Session
func SessionDomainToModel(s *sessions.Session) *models.Session {
	if s == nil {
		return nil
	}

	var revokeReason *enums.AuthSessionRevokeReason
	if s.RevokeReason() != nil {
		reason := sessionRevokeReasonDomainToEnum(*s.RevokeReason())
		revokeReason = &reason
	}

	return &models.Session{
		ID:                s.ID(),
		SessionID:         s.SessionID(),
		TenantID:          s.TenantID(),
		UserID:            s.UserID(),
		AppID:             s.AppID(),
		ClientID:          s.ClientID(),
		CreatedAt:         s.CreatedAt(),
		LastSeenAt:        s.LastSeenAt(),
		ExpiresAt:         s.ExpiresAt(),
		IP:                s.IP(),
		UaHash:            s.UAHash(), // Model 使用 UaHash，Domain 使用 UAHash
		DeviceID:          s.DeviceID(),
		DeviceFingerprint: s.DeviceFingerprint(),
		DeviceName:        s.DeviceName(),
		RevokedAt:         s.RevokedAt(),
		RevokeReason:      revokeReason,
		RevokedBy:         s.RevokedBy(),
		UpdatedAt:         s.UpdatedAt(),
	}
}

// SessionModelToDomain 将 Model Session 转换为 Domain Session
func SessionModelToDomain(m *models.Session) *sessions.Session {
	if m == nil {
		return nil
	}

	var revokeReason *sessions.SessionRevokeReason
	if m.RevokeReason != nil {
		reason := sessionRevokeReasonEnumToDomain(*m.RevokeReason)
		revokeReason = &reason
	}

	state := sessions.SessionState{
		ID:                m.ID,
		SessionID:         m.SessionID,
		TenantID:          m.TenantID,
		UserID:            m.UserID,
		AppID:             m.AppID,
		ClientID:          m.ClientID,
		CreatedAt:         m.CreatedAt,
		LastSeenAt:        m.LastSeenAt,
		ExpiresAt:         m.ExpiresAt,
		IP:                m.IP,
		UAHash:            m.UaHash, // Model 使用 UaHash，Domain 使用 UAHash
		DeviceID:          m.DeviceID,
		DeviceFingerprint: m.DeviceFingerprint,
		DeviceName:        m.DeviceName,
		RevokedAt:         m.RevokedAt,
		RevokeReason:      revokeReason,
		RevokedBy:         m.RevokedBy,
		UpdatedAt:         m.UpdatedAt,
	}

	return sessions.NewSessionFromState(state)
}

// SessionModelToUpdates 将 Model Session 转换为更新字段映射
func SessionModelToUpdates(m *models.Session) map[string]any {
	return map[string]any{
		models.SessionCols.TenantID:          m.TenantID,
		models.SessionCols.UserID:            m.UserID,
		models.SessionCols.AppID:             m.AppID,
		models.SessionCols.ClientID:          m.ClientID,
		models.SessionCols.LastSeenAt:        m.LastSeenAt,
		models.SessionCols.ExpiresAt:         m.ExpiresAt,
		models.SessionCols.IP:                m.IP,
		models.SessionCols.UaHash:            m.UaHash,
		models.SessionCols.DeviceID:           m.DeviceID,
		models.SessionCols.DeviceFingerprint: m.DeviceFingerprint,
		models.SessionCols.DeviceName:        m.DeviceName,
		models.SessionCols.RevokedAt:         m.RevokedAt,
		models.SessionCols.RevokeReason:     m.RevokeReason,
		models.SessionCols.RevokedBy:         m.RevokedBy,
		models.SessionCols.UpdatedAt:         m.UpdatedAt,
	}
}

// 枚举转换辅助函数

// SessionRevokeReasonDomainToEnum 将 Domain SessionRevokeReason 转换为 Enum SessionRevokeReason
func SessionRevokeReasonDomainToEnum(srr sessions.SessionRevokeReason) enums.AuthSessionRevokeReason {
	return sessionRevokeReasonDomainToEnum(srr)
}

func sessionRevokeReasonDomainToEnum(srr sessions.SessionRevokeReason) enums.AuthSessionRevokeReason {
	switch srr {
	case sessions.SessionRevokeReasonUserLogout:
		return enums.AuthSessionRevokeReasonUserLogout
	case sessions.SessionRevokeReasonAdminRevoke:
		return enums.AuthSessionRevokeReasonAdminRevoke
	case sessions.SessionRevokeReasonPasswordChanged:
		return enums.AuthSessionRevokeReasonPasswordChanged
	case sessions.SessionRevokeReasonDeviceChanged:
		return enums.AuthSessionRevokeReasonDeviceChanged
	case sessions.SessionRevokeReasonAccountLocked:
		return enums.AuthSessionRevokeReasonAccountLocked
	case sessions.SessionRevokeReasonSuspiciousActivity:
		return enums.AuthSessionRevokeReasonSuspiciousActivity
	case sessions.SessionRevokeReasonSessionExpired:
		return enums.AuthSessionRevokeReasonSessionExpired
	case sessions.SessionRevokeReasonOther:
		return enums.AuthSessionRevokeReasonOther
	default:
		return enums.AuthSessionRevokeReasonOther
	}
}

func sessionRevokeReasonEnumToDomain(srr enums.AuthSessionRevokeReason) sessions.SessionRevokeReason {
	switch srr {
	case enums.AuthSessionRevokeReasonUserLogout:
		return sessions.SessionRevokeReasonUserLogout
	case enums.AuthSessionRevokeReasonAdminRevoke:
		return sessions.SessionRevokeReasonAdminRevoke
	case enums.AuthSessionRevokeReasonPasswordChanged:
		return sessions.SessionRevokeReasonPasswordChanged
	case enums.AuthSessionRevokeReasonDeviceChanged:
		return sessions.SessionRevokeReasonDeviceChanged
	case enums.AuthSessionRevokeReasonAccountLocked:
		return sessions.SessionRevokeReasonAccountLocked
	case enums.AuthSessionRevokeReasonSuspiciousActivity:
		return sessions.SessionRevokeReasonSuspiciousActivity
	case enums.AuthSessionRevokeReasonSessionExpired:
		return sessions.SessionRevokeReasonSessionExpired
	case enums.AuthSessionRevokeReasonOther:
		return sessions.SessionRevokeReasonOther
	default:
		return sessions.SessionRevokeReasonOther
	}
}
