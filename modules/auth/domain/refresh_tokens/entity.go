package refresh_tokens

import (
	"time"

	"github.com/google/uuid"
)

type RevokeReason string

const (
	RevokeReasonUserLogout      RevokeReason = "user_logout"
	RevokeReasonAdminRevoke     RevokeReason = "admin_revoke"
	RevokeReasonPasswordChanged RevokeReason = "password_changed"
	RevokeReasonRotation        RevokeReason = "rotation"
	RevokeReasonAccountLocked   RevokeReason = "account_locked"
	RevokeReasonDeviceChanged   RevokeReason = "device_changed"
	RevokeReasonSuspiciousActivity RevokeReason = "suspicious_activity"
	RevokeReasonOther           RevokeReason = "other"
)

type RefreshToken struct {
	state RefreshTokenState
}

type RefreshTokenState struct {
	ID          uuid.UUID
	TokenID     string
	UserID      uuid.UUID
	AppID       *uuid.UUID
	ClientID    *string
	SessionID   *uuid.UUID
	IssuedAt    time.Time
	ExpiresAt   time.Time
	RevokedAt   *time.Time
	RevokeReason *RevokeReason
	RotatedFrom *uuid.UUID
	DeviceID    *string
	IP          *string
	UAHash      *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (rt *RefreshToken) ID() uuid.UUID            { return rt.state.ID }
func (rt *RefreshToken) TokenID() string          { return rt.state.TokenID }
func (rt *RefreshToken) UserID() uuid.UUID        { return rt.state.UserID }
func (rt *RefreshToken) AppID() *uuid.UUID        { return rt.state.AppID }
func (rt *RefreshToken) ClientID() *string        { return rt.state.ClientID }
func (rt *RefreshToken) SessionID() *uuid.UUID    { return rt.state.SessionID }
func (rt *RefreshToken) IssuedAt() time.Time      { return rt.state.IssuedAt }
func (rt *RefreshToken) ExpiresAt() time.Time     { return rt.state.ExpiresAt }
func (rt *RefreshToken) RevokedAt() *time.Time    { return rt.state.RevokedAt }
func (rt *RefreshToken) RevokeReason() *RevokeReason { return rt.state.RevokeReason }
func (rt *RefreshToken) RotatedFrom() *uuid.UUID  { return rt.state.RotatedFrom }
func (rt *RefreshToken) DeviceID() *string        { return rt.state.DeviceID }
func (rt *RefreshToken) IP() *string              { return rt.state.IP }
func (rt *RefreshToken) UAHash() *string          { return rt.state.UAHash }
func (rt *RefreshToken) CreatedAt() time.Time     { return rt.state.CreatedAt }
func (rt *RefreshToken) UpdatedAt() time.Time     { return rt.state.UpdatedAt }
