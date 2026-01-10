package results

import (
	"time"

	"nfxid/modules/directory/domain/user_emails"

	"github.com/google/uuid"
)

type UserEmailRO struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	Email             string
	IsPrimary         bool
	IsVerified        bool
	VerifiedAt        *time.Time
	VerificationToken *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

// UserEmailMapper 将 Domain UserEmail 转换为 Application UserEmailRO
func UserEmailMapper(ue *user_emails.UserEmail) UserEmailRO {
	if ue == nil {
		return UserEmailRO{}
	}

	return UserEmailRO{
		ID:                ue.ID(),
		UserID:            ue.UserID(),
		Email:             ue.Email(),
		IsPrimary:         ue.IsPrimary(),
		IsVerified:        ue.IsVerified(),
		VerifiedAt:        ue.VerifiedAt(),
		VerificationToken: ue.VerificationToken(),
		CreatedAt:         ue.CreatedAt(),
		UpdatedAt:         ue.UpdatedAt(),
		DeletedAt:         ue.DeletedAt(),
	}
}
