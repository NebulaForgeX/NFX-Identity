package login_attempts

import (
	"time"

	"github.com/google/uuid"
)

type FailureCode string

const (
	FailureCodeBadPassword      FailureCode = "bad_password"
	FailureCodeUserNotFound     FailureCode = "user_not_found"
	FailureCodeLocked           FailureCode = "locked"
	FailureCodeMFARequired      FailureCode = "mfa_required"
	FailureCodeMFAFailed        FailureCode = "mfa_failed"
	FailureCodeAccountDisabled  FailureCode = "account_disabled"
	FailureCodeCredentialExpired FailureCode = "credential_expired"
	FailureCodeRateLimited      FailureCode = "rate_limited"
	FailureCodeIPBlocked        FailureCode = "ip_blocked"
	FailureCodeDeviceNotTrusted FailureCode = "device_not_trusted"
	FailureCodeOther            FailureCode = "other"
)

type LoginAttempt struct {
	state LoginAttemptState
}

type LoginAttemptState struct {
	ID                uuid.UUID
	Identifier        string
	UserID            *uuid.UUID
	IP                *string
	UAHash            *string
	DeviceFingerprint *string
	Success           bool
	FailureCode       *FailureCode
	MFARequired       bool
	MFAVerified       bool
	CreatedAt         time.Time
}

func (la *LoginAttempt) ID() uuid.UUID              { return la.state.ID }
func (la *LoginAttempt) Identifier() string         { return la.state.Identifier }
func (la *LoginAttempt) UserID() *uuid.UUID         { return la.state.UserID }
func (la *LoginAttempt) IP() *string                 { return la.state.IP }
func (la *LoginAttempt) UAHash() *string             { return la.state.UAHash }
func (la *LoginAttempt) DeviceFingerprint() *string { return la.state.DeviceFingerprint }
func (la *LoginAttempt) Success() bool               { return la.state.Success }
func (la *LoginAttempt) FailureCode() *FailureCode   { return la.state.FailureCode }
func (la *LoginAttempt) MFARequired() bool           { return la.state.MFARequired }
func (la *LoginAttempt) MFAVerified() bool           { return la.state.MFAVerified }
func (la *LoginAttempt) CreatedAt() time.Time        { return la.state.CreatedAt }
