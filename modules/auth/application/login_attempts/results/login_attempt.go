package results

import (
	"time"

	"nfxid/modules/auth/domain/login_attempts"

	"github.com/google/uuid"
)

type LoginAttemptRO struct {
	ID                uuid.UUID
	TenantID          uuid.UUID
	Identifier        string
	UserID            *uuid.UUID
	IP                *string
	UAHash            *string
	DeviceFingerprint *string
	Success           bool
	FailureCode       *login_attempts.FailureCode
	MFARequired       bool
	MFAVerified       bool
	CreatedAt         time.Time
}

// LoginAttemptMapper 将 Domain LoginAttempt 转换为 Application LoginAttemptRO
func LoginAttemptMapper(la *login_attempts.LoginAttempt) LoginAttemptRO {
	if la == nil {
		return LoginAttemptRO{}
	}

	return LoginAttemptRO{
		ID:                la.ID(),
		TenantID:          la.TenantID(),
		Identifier:        la.Identifier(),
		UserID:            la.UserID(),
		IP:                la.IP(),
		UAHash:            la.UAHash(),
		DeviceFingerprint: la.DeviceFingerprint(),
		Success:           la.Success(),
		FailureCode:       la.FailureCode(),
		MFARequired:       la.MFARequired(),
		MFAVerified:       la.MFAVerified(),
		CreatedAt:         la.CreatedAt(),
	}
}
