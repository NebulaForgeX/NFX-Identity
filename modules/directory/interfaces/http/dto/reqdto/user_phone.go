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

type UserPhoneByIDRequestDTO struct {
	UserPhoneID uuid.UUID `uri:"user_phone_id" validate:"required,uuid"`
}

type UserPhoneSetPrimaryRequestDTO struct {
	UserPhoneID uuid.UUID `uri:"user_phone_id" validate:"required,uuid"`
}

type UserPhoneVerifyRequestDTO struct {
	UserPhoneID uuid.UUID `uri:"user_phone_id" validate:"required,uuid"`
}

type UserPhoneUpdateVerificationCodeRequestDTO struct {
	UserPhoneID           uuid.UUID `uri:"user_phone_id" validate:"required,uuid"`
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
		UserPhoneID: r.UserPhoneID,
	}
}

func (r *UserPhoneVerifyRequestDTO) ToVerifyCmd() userPhoneAppCommands.VerifyPhoneCmd {
	return userPhoneAppCommands.VerifyPhoneCmd{
		UserPhoneID: r.UserPhoneID,
	}
}

func (r *UserPhoneUpdateVerificationCodeRequestDTO) ToUpdateVerificationCodeCmd() userPhoneAppCommands.UpdateVerificationCodeCmd {
	return userPhoneAppCommands.UpdateVerificationCodeCmd{
		UserPhoneID:           r.UserPhoneID,
		VerificationCode:      r.VerificationCode,
		VerificationExpiresAt: r.VerificationExpiresAt,
	}
}
