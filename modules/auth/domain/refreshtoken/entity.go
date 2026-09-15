package refreshtoken

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct{ state RefreshTokenState }

type RefreshTokenState struct {
	ID           uuid.UUID
	AccountID    uuid.UUID
	IdentityID   *uuid.UUID
	ProfileID    *uuid.UUID
	ProfileScope *string
	DeviceID     *string
	TokenHash    string
	ExpiresAt    time.Time
	RevokedAt    *time.Time
	CreatedAt    time.Time
	DeletedAt    *time.Time
}

func NewFromState(st RefreshTokenState) *RefreshToken { return &RefreshToken{state: st} }
func (t *RefreshToken) ID() uuid.UUID                 { return t.state.ID }
func (t *RefreshToken) AccountID() uuid.UUID          { return t.state.AccountID }
func (t *RefreshToken) IdentityID() *uuid.UUID        { return t.state.IdentityID }
func (t *RefreshToken) ProfileID() *uuid.UUID         { return t.state.ProfileID }
func (t *RefreshToken) ProfileScope() *string         { return t.state.ProfileScope }
func (t *RefreshToken) DeviceID() *string             { return t.state.DeviceID }
func (t *RefreshToken) TokenHash() string             { return t.state.TokenHash }
func (t *RefreshToken) ExpiresAt() time.Time          { return t.state.ExpiresAt }
func (t *RefreshToken) RevokedAt() *time.Time         { return t.state.RevokedAt }
func (t *RefreshToken) CreatedAt() time.Time          { return t.state.CreatedAt }
func (t *RefreshToken) DeletedAt() *time.Time         { return t.state.DeletedAt }
func (t *RefreshToken) State() RefreshTokenState      { return t.state }

func (t *RefreshToken) Revoke(at time.Time) { t.state.RevokedAt = &at }
