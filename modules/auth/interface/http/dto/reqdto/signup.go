package reqdto

import "nfxidentity/enums"

// SignupRequestDTO is the body for POST /auth/signup/with-email.
type SignupRequestDTO struct {
	Email            string                    `json:"email"             validate:"required,email"`
	Password         string                    `json:"password"          validate:"required,min=8,max=128"`
	VerificationCode string                    `json:"verification_code" validate:"required,alphanum,len=6"`
	Lang             enums.AuthProfileLanguage `json:"lang"              validate:"omitempty"`
	DeviceID         string                    `json:"device_id"         validate:"required,min=1,max=128"`
	SignupPlatform   enums.AuthSignupPlatform  `json:"signup_platform"   validate:"required,oneof=nfxidentity nfxnews nfxstorages nfxvault"`
}
