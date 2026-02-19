package reqdto

import (
	userPhoneAppCommands "nfxid/modules/directory/application/user_phones/commands"

	"github.com/google/uuid"
)

type UserPhoneCreateRequestDTO struct {
	UserID                uuid.UUID `json:"user_id" validate:"required,uuid"`
	Phone                 string    `json:"phone" validate:"required"`
	CountryCode           *string   `json:"country_code,omitempty"`
	IsPrimary             bool      `json:"is_primary"`
	IsVerified            bool      `json:"is_verified"`
	VerificationCode      *string   `json:"verification_code,omitempty"`
	VerificationExpiresAt *string   `json:"verification_expires_at,omitempty"`
}

type UserPhoneSetPrimaryRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type UserPhoneVerifyRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type UserPhoneUpdateVerificationCodeRequestDTO struct {
	ID                    uuid.UUID `uri:"id" validate:"required,uuid"`
	VerificationCode      string    `json:"verification_code" validate:"required"`
	VerificationExpiresAt string    `json:"verification_expires_at" validate:"required"`
}

func (r *UserPhoneCreateRequestDTO) ToCreateCmd() userPhoneAppCommands.CreateUserPhoneCmd {
	return userPhoneAppCommands.CreateUserPhoneCmd{
		UserID:                r.UserID,
		Phone:                 r.Phone,
		CountryCode:           r.CountryCode,
		IsPrimary:             r.IsPrimary,
		IsVerified:            r.IsVerified,
		VerificationCode:      r.VerificationCode,
		VerificationExpiresAt: r.VerificationExpiresAt,
	}
}

func (r *UserPhoneSetPrimaryRequestDTO) ToSetPrimaryCmd() userPhoneAppCommands.SetPrimaryPhoneCmd {
	return userPhoneAppCommands.SetPrimaryPhoneCmd{
		UserPhoneID: r.ID,
	}
}

func (r *UserPhoneVerifyRequestDTO) ToVerifyCmd() userPhoneAppCommands.VerifyPhoneCmd {
	return userPhoneAppCommands.VerifyPhoneCmd{
		UserPhoneID: r.ID,
	}
}

func (r *UserPhoneUpdateVerificationCodeRequestDTO) ToUpdateVerificationCodeCmd() userPhoneAppCommands.UpdateVerificationCodeCmd {
	return userPhoneAppCommands.UpdateVerificationCodeCmd{
		UserPhoneID:           r.ID,
		VerificationCode:      r.VerificationCode,
		VerificationExpiresAt: r.VerificationExpiresAt,
	}
}
