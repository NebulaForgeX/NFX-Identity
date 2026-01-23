package login_attempts

import (
	"time"

	"github.com/google/uuid"
)

type NewLoginAttemptParams struct {
	Identifier        string
	UserID            *uuid.UUID
	IP                *string
	UAHash            *string
	DeviceFingerprint *string
	Success           bool
	FailureCode       *FailureCode
	MFARequired       bool
	MFAVerified       bool
}

func NewLoginAttempt(p NewLoginAttemptParams) (*LoginAttempt, error) {
	if err := validateLoginAttemptParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewLoginAttemptFromState(LoginAttemptState{
		ID:                id,
		Identifier:        p.Identifier,
		UserID:            p.UserID,
		IP:                p.IP,
		UAHash:            p.UAHash,
		DeviceFingerprint: p.DeviceFingerprint,
		Success:           p.Success,
		FailureCode:       p.FailureCode,
		MFARequired:       p.MFARequired,
		MFAVerified:       p.MFAVerified,
		CreatedAt:         now,
	}), nil
}

func NewLoginAttemptFromState(st LoginAttemptState) *LoginAttempt {
	return &LoginAttempt{state: st}
}

func validateLoginAttemptParams(p NewLoginAttemptParams) error {
	if p.Identifier == "" {
		return ErrIdentifierRequired
	}
	if p.FailureCode != nil {
		validCodes := map[FailureCode]struct{}{
			FailureCodeBadPassword:      {},
			FailureCodeUserNotFound:     {},
			FailureCodeLocked:           {},
			FailureCodeMFARequired:      {},
			FailureCodeMFAFailed:        {},
			FailureCodeAccountDisabled:  {},
			FailureCodeCredentialExpired: {},
			FailureCodeRateLimited:      {},
			FailureCodeIPBlocked:        {},
			FailureCodeDeviceNotTrusted: {},
			FailureCodeOther:            {},
		}
		if _, ok := validCodes[*p.FailureCode]; !ok {
			return ErrInvalidFailureCode
		}
	}
	return nil
}
