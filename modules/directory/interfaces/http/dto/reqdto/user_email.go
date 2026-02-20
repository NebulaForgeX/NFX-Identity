package reqdto

import (
	userEmailAppCommands "nfxid/modules/directory/application/user_emails/commands"

	"github.com/google/uuid"
)

type UserEmailCreateRequestDTO struct {
	UserID            uuid.UUID `json:"user_id" validate:"required,uuid"`
	Email             string    `json:"email" validate:"required,email"`
	IsPrimary         bool      `json:"is_primary,omitempty"`
	IsVerified        bool      `json:"is_verified,omitempty"`
	VerificationToken *string   `json:"verification_token,omitempty"`
}

type UserEmailByIDRequestDTO struct {
	UserEmailID uuid.UUID `uri:"user_email_id" validate:"required,uuid"`
}

type UserEmailSetPrimaryRequestDTO struct {
	UserEmailID uuid.UUID `uri:"user_email_id" validate:"required,uuid"`
}

type UserEmailVerifyRequestDTO struct {
	UserEmailID uuid.UUID `uri:"user_email_id" validate:"required,uuid"`
}

func (r *UserEmailCreateRequestDTO) ToCreateCmd() userEmailAppCommands.CreateUserEmailCmd {
	return userEmailAppCommands.CreateUserEmailCmd{
		UserID:            r.UserID,
		Email:             r.Email,
		IsPrimary:         r.IsPrimary,
		IsVerified:        r.IsVerified,
		VerificationToken: r.VerificationToken,
	}
}

func (r *UserEmailSetPrimaryRequestDTO) ToSetPrimaryCmd() userEmailAppCommands.SetPrimaryEmailCmd {
	return userEmailAppCommands.SetPrimaryEmailCmd{
		UserEmailID: r.UserEmailID,
	}
}

func (r *UserEmailVerifyRequestDTO) ToVerifyCmd() userEmailAppCommands.VerifyEmailCmd {
	return userEmailAppCommands.VerifyEmailCmd{
		UserEmailID: r.UserEmailID,
	}
}
