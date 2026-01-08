package password_history

import (
	"time"

	"github.com/google/uuid"
)

type PasswordHistory struct {
	state PasswordHistoryState
}

type PasswordHistoryState struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	TenantID    uuid.UUID
	PasswordHash string
	HashAlg     *string
	CreatedAt   time.Time
}

func (ph *PasswordHistory) ID() uuid.UUID          { return ph.state.ID }
func (ph *PasswordHistory) UserID() uuid.UUID      { return ph.state.UserID }
func (ph *PasswordHistory) TenantID() uuid.UUID    { return ph.state.TenantID }
func (ph *PasswordHistory) PasswordHash() string   { return ph.state.PasswordHash }
func (ph *PasswordHistory) HashAlg() *string       { return ph.state.HashAlg }
func (ph *PasswordHistory) CreatedAt() time.Time   { return ph.state.CreatedAt }
