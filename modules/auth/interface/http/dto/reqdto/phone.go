package reqdto

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

type CreatePhone struct {
	Phone string `json:"phone"`
}

func (r CreatePhone) Normalize() CreatePhone {
	return CreatePhone{Phone: strings.TrimSpace(r.Phone)}
}

func (r CreatePhone) Validate() error {
	if r.Phone == "" {
		return authErr.ErrInvalidPhone
	}
	return nil
}

type UpdatePhone struct {
	Phone string `json:"phone"`
}

func (r UpdatePhone) Normalize() UpdatePhone {
	return UpdatePhone{Phone: strings.TrimSpace(r.Phone)}
}

func (r UpdatePhone) Validate() error {
	if r.Phone == "" {
		return authErr.ErrInvalidPhone
	}
	return nil
}

type VerifyPhone struct {
	VerificationCode string `json:"verification_code"`
}

func (r VerifyPhone) Normalize() VerifyPhone {
	return VerifyPhone{VerificationCode: strings.TrimSpace(r.VerificationCode)}
}

func (r VerifyPhone) Validate() error {
	if r.VerificationCode == "" {
		return authErr.ErrVerificationCodeWrong
	}
	return nil
}

type PhoneURI struct {
	PhoneID uuid.UUID `uri:"phoneId"`
}
