package reqdto

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

type CreateEmail struct {
	Email string `json:"email"`
}

func (r CreateEmail) Normalize() CreateEmail {
	return CreateEmail{Email: strings.ToLower(strings.TrimSpace(r.Email))}
}

func (r CreateEmail) Validate() error {
	if r.Email == "" {
		return authErr.ErrEmailAddressRequired
	}
	return nil
}

type UpdateEmail struct {
	Email string `json:"email"`
}

func (r UpdateEmail) Normalize() UpdateEmail {
	return UpdateEmail{Email: strings.ToLower(strings.TrimSpace(r.Email))}
}

func (r UpdateEmail) Validate() error {
	if r.Email == "" {
		return authErr.ErrEmailAddressRequired
	}
	return nil
}

type VerifyEmail struct {
	VerificationCode string `json:"verification_code"`
}

func (r VerifyEmail) Normalize() VerifyEmail {
	return VerifyEmail{VerificationCode: strings.TrimSpace(r.VerificationCode)}
}

func (r VerifyEmail) Validate() error {
	if r.VerificationCode == "" {
		return authErr.ErrVerificationCodeWrong
	}
	return nil
}

type SendEmailVerificationCode struct {
	Lang string `json:"lang"`
}

type EmailURI struct {
	EmailID uuid.UUID `uri:"emailId"`
}

type SendChangePasswordVerificationCode struct {
	Lang string `json:"lang"`
}

type ChangePassword struct {
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	VerificationCode string `json:"verification_code"`
}

func (r ChangePassword) Normalize() ChangePassword {
	return ChangePassword{
		CurrentPassword:  r.CurrentPassword,
		NewPassword:      r.NewPassword,
		VerificationCode: strings.TrimSpace(r.VerificationCode),
	}
}

func (r ChangePassword) Validate() error {
	if strings.TrimSpace(r.CurrentPassword) == "" || strings.TrimSpace(r.NewPassword) == "" {
		return authErr.ErrAccountPasswordRequired
	}
	if len(r.NewPassword) < 8 || len(r.NewPassword) > 128 {
		return authErr.ErrAccountPasswordRequired
	}
	if strings.TrimSpace(r.VerificationCode) == "" {
		return authErr.ErrVerificationCodeWrong
	}
	return nil
}
