package identity

import (
	"time"

	"nfxidentity/enums"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Identity struct {
	state IdentityState
}

type IdentityState struct {
	ID               uuid.UUID
	AccountID        uuid.UUID
	IdentityProvider enums.AuthIdentityProvider
	ProviderSubject  string
	PasswordHash     *string
	Metadata         *datatypes.JSON
	LastLoginAt      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func (i *Identity) ID() uuid.UUID        { return i.state.ID }
func (i *Identity) AccountID() uuid.UUID { return i.state.AccountID }
func (i *Identity) IdentityProvider() enums.AuthIdentityProvider {
	return i.state.IdentityProvider
}
func (i *Identity) ProviderSubject() string   { return i.state.ProviderSubject }
func (i *Identity) PasswordHash() *string     { return i.state.PasswordHash }
func (i *Identity) Metadata() *datatypes.JSON { return i.state.Metadata }
func (i *Identity) LastLoginAt() *time.Time   { return i.state.LastLoginAt }
func (i *Identity) CreatedAt() time.Time      { return i.state.CreatedAt }
func (i *Identity) UpdatedAt() time.Time      { return i.state.UpdatedAt }
func (i *Identity) DeletedAt() *time.Time     { return i.state.DeletedAt }

func NewIdentityFromState(st IdentityState) *Identity {
	return &Identity{state: st}
}
