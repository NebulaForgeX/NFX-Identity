package reqdto

import "nfxidentity/enums"

// SendVerificationCodeRequestDTO is the body for POST /auth/signup/send-code.
type SendVerificationCodeRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
	Lang  enums.AuthProfileLanguage `json:"lang"  validate:"omitempty"`
}
