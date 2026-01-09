package mapper

import (
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// UserPhoneDomainToModel 将 Domain UserPhone 转换为 Model UserPhone
func UserPhoneDomainToModel(up *user_phones.UserPhone) *models.UserPhone {
	if up == nil {
		return nil
	}

	return &models.UserPhone{
		ID:                    up.ID(),
		UserID:                up.UserID(),
		Phone:                 up.Phone(),
		CountryCode:           up.CountryCode(),
		IsPrimary:             up.IsPrimary(),
		IsVerified:            up.IsVerified(),
		VerifiedAt:            up.VerifiedAt(),
		VerificationCode:      up.VerificationCode(),
		VerificationExpiresAt: up.VerificationExpiresAt(),
		CreatedAt:             up.CreatedAt(),
		UpdatedAt:             up.UpdatedAt(),
		DeletedAt:             timex.TimeToGormDeletedAt(up.DeletedAt()),
	}
}

// UserPhoneModelToDomain 将 Model UserPhone 转换为 Domain UserPhone
func UserPhoneModelToDomain(m *models.UserPhone) *user_phones.UserPhone {
	if m == nil {
		return nil
	}

	state := user_phones.UserPhoneState{
		ID:                    m.ID,
		UserID:                m.UserID,
		Phone:                 m.Phone,
		CountryCode:           m.CountryCode,
		IsPrimary:             m.IsPrimary,
		IsVerified:            m.IsVerified,
		VerifiedAt:            m.VerifiedAt,
		VerificationCode:      m.VerificationCode,
		VerificationExpiresAt: m.VerificationExpiresAt,
		CreatedAt:             m.CreatedAt,
		UpdatedAt:             m.UpdatedAt,
		DeletedAt:             timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_phones.NewUserPhoneFromState(state)
}

// UserPhoneModelToUpdates 将 Model UserPhone 转换为更新字段映射
func UserPhoneModelToUpdates(m *models.UserPhone) map[string]any {
	return map[string]any{
		models.UserPhoneCols.UserID:                m.UserID,
		models.UserPhoneCols.Phone:                 m.Phone,
		models.UserPhoneCols.CountryCode:           m.CountryCode,
		models.UserPhoneCols.IsPrimary:             m.IsPrimary,
		models.UserPhoneCols.IsVerified:            m.IsVerified,
		models.UserPhoneCols.VerifiedAt:            m.VerifiedAt,
		models.UserPhoneCols.VerificationCode:      m.VerificationCode,
		models.UserPhoneCols.VerificationExpiresAt: m.VerificationExpiresAt,
		models.UserPhoneCols.UpdatedAt:             m.UpdatedAt,
		models.UserPhoneCols.DeletedAt:             m.DeletedAt,
	}
}
