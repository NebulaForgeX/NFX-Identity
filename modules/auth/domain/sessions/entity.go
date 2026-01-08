package sessions

import (
	"time"

	"github.com/google/uuid"
)

type SessionRevokeReason string

const (
	SessionRevokeReasonUserLogout      SessionRevokeReason = "user_logout"
	SessionRevokeReasonAdminRevoke     SessionRevokeReason = "admin_revoke"
	SessionRevokeReasonPasswordChanged SessionRevokeReason = "password_changed"
	SessionRevokeReasonDeviceChanged   SessionRevokeReason = "device_changed"
	SessionRevokeReasonAccountLocked   SessionRevokeReason = "account_locked"
	SessionRevokeReasonSuspiciousActivity SessionRevokeReason = "suspicious_activity"
	SessionRevokeReasonSessionExpired  SessionRevokeReason = "session_expired"
	SessionRevokeReasonOther           SessionRevokeReason = "other"
)

type Session struct {
	state SessionState
}

type SessionState struct {
	ID              uuid.UUID
	SessionID       string
	TenantID        uuid.UUID
	UserID          uuid.UUID
	AppID           *uuid.UUID
	ClientID        *string
	CreatedAt       time.Time
	LastSeenAt      time.Time
	ExpiresAt       time.Time
	IP              *string
	UAHash          *string
	DeviceID        *string
	DeviceFingerprint *string
	DeviceName      *string
	RevokedAt       *time.Time
	RevokeReason    *SessionRevokeReason
	RevokedBy       *string
	UpdatedAt       time.Time
}

func (s *Session) ID() uuid.UUID                  { return s.state.ID }
func (s *Session) SessionID() string              { return s.state.SessionID }
func (s *Session) TenantID() uuid.UUID            { return s.state.TenantID }
func (s *Session) UserID() uuid.UUID              { return s.state.UserID }
func (s *Session) AppID() *uuid.UUID              { return s.state.AppID }
func (s *Session) ClientID() *string              { return s.state.ClientID }
func (s *Session) CreatedAt() time.Time           { return s.state.CreatedAt }
func (s *Session) LastSeenAt() time.Time          { return s.state.LastSeenAt }
func (s *Session) ExpiresAt() time.Time           { return s.state.ExpiresAt }
func (s *Session) IP() *string                    { return s.state.IP }
func (s *Session) UAHash() *string                { return s.state.UAHash }
func (s *Session) DeviceID() *string              { return s.state.DeviceID }
func (s *Session) DeviceFingerprint() *string     { return s.state.DeviceFingerprint }
func (s *Session) DeviceName() *string            { return s.state.DeviceName }
func (s *Session) RevokedAt() *time.Time          { return s.state.RevokedAt }
func (s *Session) RevokeReason() *SessionRevokeReason { return s.state.RevokeReason }
func (s *Session) RevokedBy() *string             { return s.state.RevokedBy }
func (s *Session) UpdatedAt() time.Time           { return s.state.UpdatedAt }
