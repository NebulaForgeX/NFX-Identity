package reqdto

import (
	"nfxidentity/enums"

	"github.com/google/uuid"
)

// LoginWithEmailRequestDTO is the body for POST /auth/login/with-email.
type LoginWithEmailRequestDTO struct {
	Email    string `json:"email"     validate:"required,email"`
	Password string `json:"password"  validate:"required,min=8,max=128"`
	DeviceID string `json:"device_id" validate:"required,min=1,max=128"`
}

// LoginWithPhoneRequestDTO is the body for POST /auth/login/with-phone.
type LoginWithPhoneRequestDTO struct {
	Phone    string `json:"phone"     validate:"required,min=1,max=32"`
	Password string `json:"password"  validate:"required,min=8,max=128"`
	DeviceID string `json:"device_id" validate:"required,min=1,max=128"`
}

// RefreshTokensRequestDTO is the body for POST /auth/refresh.
type RefreshTokensRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
	DeviceID     string `json:"device_id"     validate:"required,min=1,max=128"`
}

// LogoutRequestDTO is the body for POST /auth/logout.
type LogoutRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// SelectProfileRequestDTO is the body for POST /auth/me/select-profile.
type SelectProfileRequestDTO struct {
	ProfileID uuid.UUID              `json:"profile_id" validate:"required,uuid"`
	Kind      enums.AuthProfileScope `json:"kind"       validate:"required,oneof=community authority"`
	DeviceID  string                 `json:"device_id"  validate:"required,min=1,max=128"`
}
