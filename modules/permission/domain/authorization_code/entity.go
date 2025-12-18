package authorization_code

import (
	"time"

	"github.com/google/uuid"
)

type AuthorizationCode struct {
	state AuthorizationCodeState
}

type AuthorizationCodeState struct {
	ID        uuid.UUID
	Editable  AuthorizationCodeEditable
	CreatedBy *uuid.UUID
	ExpiresAt *time.Time
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type AuthorizationCodeEditable struct {
	Code      string
	MaxUses   int
	UsedCount int
}

func (ac *AuthorizationCode) ID() uuid.UUID                       { return ac.state.ID }
func (ac *AuthorizationCode) Editable() AuthorizationCodeEditable { return ac.state.Editable }
func (ac *AuthorizationCode) CreatedBy() *uuid.UUID               { return ac.state.CreatedBy }
func (ac *AuthorizationCode) ExpiresAt() *time.Time               { return ac.state.ExpiresAt }
func (ac *AuthorizationCode) IsActive() bool                      { return ac.state.IsActive }
func (ac *AuthorizationCode) CreatedAt() time.Time                { return ac.state.CreatedAt }
func (ac *AuthorizationCode) UpdatedAt() time.Time                { return ac.state.UpdatedAt }
func (ac *AuthorizationCode) DeletedAt() *time.Time               { return ac.state.DeletedAt }

func (ac *AuthorizationCode) IsDeleted() bool {
	return ac.DeletedAt() != nil
}

func (ac *AuthorizationCode) IsAvailable() bool {
	if ac.IsDeleted() {
		return false
	}
	if !ac.IsActive() {
		return false
	}
	if ac.state.Editable.UsedCount >= ac.state.Editable.MaxUses {
		return false
	}
	if ac.ExpiresAt() != nil && time.Now().UTC().After(*ac.ExpiresAt()) {
		return false
	}
	return true
}
