package user_phones

import (
	"time"

	"github.com/google/uuid"
)

type NewUserPhoneParams struct {
	UserID                uuid.UUID
	Phone                 string
	CountryCode           *string
	IsPrimary             bool
	IsVerified            bool
	VerificationCode      *string
	VerificationExpiresAt *time.Time
}

func NewUserPhone(p NewUserPhoneParams) (*UserPhone, error) {
	if err := validateUserPhoneParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserPhoneFromState(UserPhoneState{
		ID:                    id,
		UserID:                p.UserID,
		Phone:                 p.Phone,
		CountryCode:           p.CountryCode,
		IsPrimary:             p.IsPrimary,
		IsVerified:            p.IsVerified,
		VerifiedAt:            nil,
		VerificationCode:      p.VerificationCode,
		VerificationExpiresAt: p.VerificationExpiresAt,
		CreatedAt:             now,
		UpdatedAt:             now,
	}), nil
}

func NewUserPhoneFromState(st UserPhoneState) *UserPhone {
	return &UserPhone{state: st}
}

func validateUserPhoneParams(p NewUserPhoneParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.Phone == "" {
		return ErrPhoneRequired
	}
	return nil
}
