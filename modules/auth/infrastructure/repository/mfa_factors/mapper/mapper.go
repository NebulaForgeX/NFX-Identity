package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// MFAFactorDomainToModel 将 Domain MFAFactor 转换为 Model MfaFactor
func MFAFactorDomainToModel(mf *mfa_factors.MFAFactor) *models.MfaFactor {
	if mf == nil {
		return nil
	}

	return &models.MfaFactor{
		ID:                mf.ID(),
		FactorID:          mf.FactorID(),
		UserID:            mf.UserID(),
		Type:              mfaTypeDomainToEnum(mf.Type()),
		SecretEncrypted:   mf.SecretEncrypted(),
		Phone:             mf.Phone(),
		Email:             mf.Email(),
		Name:              mf.Name(),
		Enabled:           mf.Enabled(),
		CreatedAt:         mf.CreatedAt(),
		LastUsedAt:        mf.LastUsedAt(),
		RecoveryCodesHash: mf.RecoveryCodesHash(),
		UpdatedAt:         mf.UpdatedAt(),
		DeletedAt:         timex.TimeToGormDeletedAt(mf.DeletedAt()),
	}
}

// MFAFactorModelToDomain 将 Model MfaFactor 转换为 Domain MFAFactor
func MFAFactorModelToDomain(m *models.MfaFactor) *mfa_factors.MFAFactor {
	if m == nil {
		return nil
	}

	state := mfa_factors.MFAFactorState{
		ID:                m.ID,
		FactorID:          m.FactorID,
		UserID:            m.UserID,
		Type:              mfaTypeEnumToDomain(m.Type),
		SecretEncrypted:   m.SecretEncrypted,
		Phone:             m.Phone,
		Email:             m.Email,
		Name:              m.Name,
		Enabled:           m.Enabled,
		CreatedAt:         m.CreatedAt,
		LastUsedAt:        m.LastUsedAt,
		RecoveryCodesHash: m.RecoveryCodesHash,
		UpdatedAt:         m.UpdatedAt,
		DeletedAt:         timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return mfa_factors.NewMFAFactorFromState(state)
}

// MFAFactorModelToUpdates 将 Model MfaFactor 转换为更新字段映射
func MFAFactorModelToUpdates(m *models.MfaFactor) map[string]any {
	return map[string]any{
		models.MfaFactorCols.UserID:            m.UserID,
		models.MfaFactorCols.Type:              m.Type,
		models.MfaFactorCols.SecretEncrypted:   m.SecretEncrypted,
		models.MfaFactorCols.Phone:             m.Phone,
		models.MfaFactorCols.Email:             m.Email,
		models.MfaFactorCols.Name:              m.Name,
		models.MfaFactorCols.Enabled:           m.Enabled,
		models.MfaFactorCols.LastUsedAt:        m.LastUsedAt,
		models.MfaFactorCols.RecoveryCodesHash: m.RecoveryCodesHash,
		models.MfaFactorCols.UpdatedAt:         m.UpdatedAt,
		models.MfaFactorCols.DeletedAt:         m.DeletedAt,
	}
}

// 枚举转换辅助函数

func mfaTypeDomainToEnum(mt mfa_factors.MFAType) enums.AuthMfaType {
	switch mt {
	case mfa_factors.MFATypeTOTP:
		return enums.AuthMfaTypeTotp
	case mfa_factors.MFATypeSMS:
		return enums.AuthMfaTypeSms
	case mfa_factors.MFATypeEmail:
		return enums.AuthMfaTypeEmail
	case mfa_factors.MFATypeWebAuthn:
		return enums.AuthMfaTypeWebauthn
	case mfa_factors.MFATypeBackupCode:
		return enums.AuthMfaTypeBackupCode
	default:
		return enums.AuthMfaTypeTotp
	}
}

func mfaTypeEnumToDomain(mt enums.AuthMfaType) mfa_factors.MFAType {
	switch mt {
	case enums.AuthMfaTypeTotp:
		return mfa_factors.MFATypeTOTP
	case enums.AuthMfaTypeSms:
		return mfa_factors.MFATypeSMS
	case enums.AuthMfaTypeEmail:
		return mfa_factors.MFATypeEmail
	case enums.AuthMfaTypeWebauthn:
		return mfa_factors.MFATypeWebAuthn
	case enums.AuthMfaTypeBackupCode:
		return mfa_factors.MFATypeBackupCode
	default:
		return mfa_factors.MFATypeTOTP
	}
}
