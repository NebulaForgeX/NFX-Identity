package respdto

import (
	"time"

	loginAttemptAppResult "nfxid/modules/auth/application/login_attempts/results"

	"github.com/google/uuid"
)

type LoginAttemptDTO struct {
	ID                uuid.UUID  `json:"id"`
	Identifier        string     `json:"identifier"`
	UserID            *uuid.UUID `json:"user_id,omitempty"`
	IP                *string    `json:"ip,omitempty"`
	UAHash            *string    `json:"ua_hash,omitempty"`
	DeviceFingerprint *string    `json:"device_fingerprint,omitempty"`
	Success           bool       `json:"success"`
	FailureCode       *string    `json:"failure_code,omitempty"`
	MFARequired       bool       `json:"mfa_required"`
	MFAVerified       bool       `json:"mfa_verified"`
	CreatedAt         time.Time  `json:"created_at"`
}

// LoginAttemptROToDTO converts application LoginAttemptRO to response DTO
func LoginAttemptROToDTO(v *loginAttemptAppResult.LoginAttemptRO) *LoginAttemptDTO {
	if v == nil {
		return nil
	}

	var failureCode *string
	if v.FailureCode != nil {
		code := string(*v.FailureCode)
		failureCode = &code
	}

	return &LoginAttemptDTO{
		ID:                v.ID,
		Identifier:        v.Identifier,
		UserID:            v.UserID,
		IP:                v.IP,
		UAHash:            v.UAHash,
		DeviceFingerprint: v.DeviceFingerprint,
		Success:           v.Success,
		FailureCode:       failureCode,
		MFARequired:       v.MFARequired,
		MFAVerified:       v.MFAVerified,
		CreatedAt:         v.CreatedAt,
	}
}
