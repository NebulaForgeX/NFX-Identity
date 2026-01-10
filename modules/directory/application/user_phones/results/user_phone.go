package results

import (
	"time"

	"nfxid/modules/directory/domain/user_phones"

	"github.com/google/uuid"
)

type UserPhoneRO struct {
	ID                    uuid.UUID
	UserID                uuid.UUID
	Phone                 string
	CountryCode           *string
	IsPrimary             bool
	IsVerified            bool
	VerifiedAt            *time.Time
	VerificationCode      *string
	VerificationExpiresAt *time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
}

// UserPhoneMapper 将 Domain UserPhone 转换为 Application UserPhoneRO
func UserPhoneMapper(up *user_phones.UserPhone) UserPhoneRO {
	if up == nil {
		return UserPhoneRO{}
	}

	return UserPhoneRO{
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
		DeletedAt:             up.DeletedAt(),
	}
}
