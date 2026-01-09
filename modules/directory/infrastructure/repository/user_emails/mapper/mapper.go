package mapper

import (
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// UserEmailDomainToModel 将 Domain UserEmail 转换为 Model UserEmail
func UserEmailDomainToModel(ue *user_emails.UserEmail) *models.UserEmail {
	if ue == nil {
		return nil
	}

	return &models.UserEmail{
		ID:                ue.ID(),
		UserID:            ue.UserID(),
		Email:             ue.Email(),
		IsPrimary:         ue.IsPrimary(),
		IsVerified:        ue.IsVerified(),
		VerifiedAt:        ue.VerifiedAt(),
		VerificationToken: ue.VerificationToken(),
		CreatedAt:         ue.CreatedAt(),
		UpdatedAt:         ue.UpdatedAt(),
		DeletedAt:         timex.TimeToGormDeletedAt(ue.DeletedAt()),
	}
}

// UserEmailModelToDomain 将 Model UserEmail 转换为 Domain UserEmail
func UserEmailModelToDomain(m *models.UserEmail) *user_emails.UserEmail {
	if m == nil {
		return nil
	}

	state := user_emails.UserEmailState{
		ID:                m.ID,
		UserID:            m.UserID,
		Email:             m.Email,
		IsPrimary:         m.IsPrimary,
		IsVerified:        m.IsVerified,
		VerifiedAt:        m.VerifiedAt,
		VerificationToken: m.VerificationToken,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		DeletedAt:         timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_emails.NewUserEmailFromState(state)
}

// UserEmailModelToUpdates 将 Model UserEmail 转换为更新字段映射
func UserEmailModelToUpdates(m *models.UserEmail) map[string]any {
	return map[string]any{
		models.UserEmailCols.UserID:            m.UserID,
		models.UserEmailCols.Email:             m.Email,
		models.UserEmailCols.IsPrimary:         m.IsPrimary,
		models.UserEmailCols.IsVerified:        m.IsVerified,
		models.UserEmailCols.VerifiedAt:        m.VerifiedAt,
		models.UserEmailCols.VerificationToken: m.VerificationToken,
		models.UserEmailCols.UpdatedAt:         m.UpdatedAt,
		models.UserEmailCols.DeletedAt:         m.DeletedAt,
	}
}
