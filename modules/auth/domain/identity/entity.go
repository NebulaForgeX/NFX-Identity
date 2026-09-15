package identity

import (
	"time"

	"github.com/google/uuid"
)

type Identity struct{ state IdentityState }

type IdentityState struct {
	ID               uuid.UUID
	AccountID        uuid.UUID
	IdentityProvider string
	ProviderSubject  string
	PasswordHash     *string
	LastLoginAt      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func NewFromState(st IdentityState) *Identity { return &Identity{state: st} }
func (i *Identity) ID() uuid.UUID               { return i.state.ID }
func (i *Identity) AccountID() uuid.UUID        { return i.state.AccountID }
func (i *Identity) IdentityProvider() string    { return i.state.IdentityProvider }
func (i *Identity) ProviderSubject() string     { return i.state.ProviderSubject }
func (i *Identity) PasswordHash() *string       { return i.state.PasswordHash }
func (i *Identity) LastLoginAt() *time.Time     { return i.state.LastLoginAt }
func (i *Identity) CreatedAt() time.Time        { return i.state.CreatedAt }
func (i *Identity) UpdatedAt() time.Time        { return i.state.UpdatedAt }
func (i *Identity) DeletedAt() *time.Time       { return i.state.DeletedAt }
func (i *Identity) State() IdentityState        { return i.state }

func (i *Identity) TouchLogin(at time.Time) {
	i.state.LastLoginAt = &at
	i.state.UpdatedAt = at
}
func (i *Identity) SetPasswordHash(hash string) {
	i.state.PasswordHash = &hash
	i.state.UpdatedAt = time.Now()
}
func (i *Identity) SoftDelete(at time.Time) {
	i.state.DeletedAt = &at
	i.state.UpdatedAt = at
}
