package reqdto

import (
	loginAttemptAppCommands "nfxid/modules/auth/application/login_attempts/commands"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"

	"github.com/google/uuid"
)

type LoginAttemptCreateRequestDTO struct {
	Identifier        string  `json:"identifier" validate:"required"`
	UserID            *string `json:"user_id,omitempty"`
	IP                *string `json:"ip,omitempty"`
	UAHash            *string `json:"ua_hash,omitempty"`
	DeviceFingerprint *string `json:"device_fingerprint,omitempty"`
	Success           bool    `json:"success"`
	FailureCode       *string `json:"failure_code,omitempty"`
	MFARequired       bool    `json:"mfa_required"`
	MFAVerified       bool    `json:"mfa_verified"`
}

type LoginAttemptByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *LoginAttemptCreateRequestDTO) ToCreateCmd() loginAttemptAppCommands.CreateLoginAttemptCmd {
	cmd := loginAttemptAppCommands.CreateLoginAttemptCmd{
		Identifier:        r.Identifier,
		IP:                r.IP,
		UAHash:            r.UAHash,
		DeviceFingerprint: r.DeviceFingerprint,
		Success:           r.Success,
		MFARequired:       r.MFARequired,
		MFAVerified:       r.MFAVerified,
	}

	if r.UserID != nil {
		uid, err := uuid.Parse(*r.UserID)
		if err == nil {
			cmd.UserID = &uid
		}
	}

	if r.FailureCode != nil {
		code := loginAttemptDomain.FailureCode(*r.FailureCode)
		cmd.FailureCode = &code
	}

	return cmd
}
