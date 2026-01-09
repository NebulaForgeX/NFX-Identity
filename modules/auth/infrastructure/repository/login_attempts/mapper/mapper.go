package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// LoginAttemptDomainToModel 将 Domain LoginAttempt 转换为 Model LoginAttempt
func LoginAttemptDomainToModel(la *login_attempts.LoginAttempt) *models.LoginAttempt {
	if la == nil {
		return nil
	}

	var failureCode *enums.AuthFailureCode
	if la.FailureCode() != nil {
		code := failureCodeDomainToEnum(*la.FailureCode())
		failureCode = &code
	}

	return &models.LoginAttempt{
		ID:                la.ID(),
		TenantID:          la.TenantID(),
		Identifier:        la.Identifier(),
		UserID:            la.UserID(),
		IP:                la.IP(),
		UaHash:            la.UAHash(), // Model 使用 UaHash，Domain 使用 UAHash
		DeviceFingerprint: la.DeviceFingerprint(),
		Success:           la.Success(),
		FailureCode:       failureCode,
		MfaRequired:       la.MFARequired(), // Model 使用 MfaRequired，Domain 使用 MFARequired
		MfaVerified:       la.MFAVerified(), // Model 使用 MfaVerified，Domain 使用 MFAVerified
		CreatedAt:         la.CreatedAt(),
	}
}

// LoginAttemptModelToDomain 将 Model LoginAttempt 转换为 Domain LoginAttempt
func LoginAttemptModelToDomain(m *models.LoginAttempt) *login_attempts.LoginAttempt {
	if m == nil {
		return nil
	}

	var failureCode *login_attempts.FailureCode
	if m.FailureCode != nil {
		code := failureCodeEnumToDomain(*m.FailureCode)
		failureCode = &code
	}

	state := login_attempts.LoginAttemptState{
		ID:                m.ID,
		TenantID:          m.TenantID,
		Identifier:        m.Identifier,
		UserID:            m.UserID,
		IP:                m.IP,
		UAHash:            m.UaHash, // Model 使用 UaHash，Domain 使用 UAHash
		DeviceFingerprint: m.DeviceFingerprint,
		Success:           m.Success,
		FailureCode:       failureCode,
		MFARequired:       m.MfaRequired, // Model 使用 MfaRequired，Domain 使用 MFARequired
		MFAVerified:       m.MfaVerified, // Model 使用 MfaVerified，Domain 使用 MFAVerified
		CreatedAt:         m.CreatedAt,
	}

	return login_attempts.NewLoginAttemptFromState(state)
}

// LoginAttemptModelToUpdates 将 Model LoginAttempt 转换为更新字段映射
func LoginAttemptModelToUpdates(m *models.LoginAttempt) map[string]any {
	return map[string]any{
		models.LoginAttemptCols.TenantID:          m.TenantID,
		models.LoginAttemptCols.Identifier:        m.Identifier,
		models.LoginAttemptCols.UserID:            m.UserID,
		models.LoginAttemptCols.IP:                m.IP,
		models.LoginAttemptCols.UaHash:            m.UaHash,
		models.LoginAttemptCols.DeviceFingerprint: m.DeviceFingerprint,
		models.LoginAttemptCols.Success:          m.Success,
		models.LoginAttemptCols.FailureCode:      m.FailureCode,
		models.LoginAttemptCols.MfaRequired:      m.MfaRequired,
		models.LoginAttemptCols.MfaVerified:      m.MfaVerified,
	}
}

// 枚举转换辅助函数

func failureCodeDomainToEnum(fc login_attempts.FailureCode) enums.AuthFailureCode {
	switch fc {
	case login_attempts.FailureCodeBadPassword:
		return enums.AuthFailureCodeBadPassword
	case login_attempts.FailureCodeUserNotFound:
		return enums.AuthFailureCodeUserNotFound
	case login_attempts.FailureCodeLocked:
		return enums.AuthFailureCodeLocked
	case login_attempts.FailureCodeMFARequired:
		return enums.AuthFailureCodeMfaRequired
	case login_attempts.FailureCodeMFAFailed:
		return enums.AuthFailureCodeMfaFailed
	case login_attempts.FailureCodeAccountDisabled:
		return enums.AuthFailureCodeAccountDisabled
	case login_attempts.FailureCodeCredentialExpired:
		return enums.AuthFailureCodeCredentialExpired
	case login_attempts.FailureCodeRateLimited:
		return enums.AuthFailureCodeRateLimited
	case login_attempts.FailureCodeIPBlocked:
		return enums.AuthFailureCodeIpBlocked
	case login_attempts.FailureCodeDeviceNotTrusted:
		return enums.AuthFailureCodeDeviceNotTrusted
	case login_attempts.FailureCodeOther:
		return enums.AuthFailureCodeOther
	default:
		return enums.AuthFailureCodeOther
	}
}

func failureCodeEnumToDomain(fc enums.AuthFailureCode) login_attempts.FailureCode {
	switch fc {
	case enums.AuthFailureCodeBadPassword:
		return login_attempts.FailureCodeBadPassword
	case enums.AuthFailureCodeUserNotFound:
		return login_attempts.FailureCodeUserNotFound
	case enums.AuthFailureCodeLocked:
		return login_attempts.FailureCodeLocked
	case enums.AuthFailureCodeMfaRequired:
		return login_attempts.FailureCodeMFARequired
	case enums.AuthFailureCodeMfaFailed:
		return login_attempts.FailureCodeMFAFailed
	case enums.AuthFailureCodeAccountDisabled:
		return login_attempts.FailureCodeAccountDisabled
	case enums.AuthFailureCodeCredentialExpired:
		return login_attempts.FailureCodeCredentialExpired
	case enums.AuthFailureCodeRateLimited:
		return login_attempts.FailureCodeRateLimited
	case enums.AuthFailureCodeIpBlocked:
		return login_attempts.FailureCodeIPBlocked
	case enums.AuthFailureCodeDeviceNotTrusted:
		return login_attempts.FailureCodeDeviceNotTrusted
	case enums.AuthFailureCodeOther:
		return login_attempts.FailureCodeOther
	default:
		return login_attempts.FailureCodeOther
	}
}
