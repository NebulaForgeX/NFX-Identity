package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct{ state AccountState }

type AccountState struct {
	ID             uuid.UUID
	AccountStatus  string
	SignupPlatform string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewFromState(st AccountState) *Account { return &Account{state: st} }
func (a *Account) ID() uuid.UUID            { return a.state.ID }
func (a *Account) AccountStatus() string    { return a.state.AccountStatus }
func (a *Account) SignupPlatform() string   { return a.state.SignupPlatform }
func (a *Account) CreatedAt() time.Time     { return a.state.CreatedAt }
func (a *Account) UpdatedAt() time.Time     { return a.state.UpdatedAt }
func (a *Account) DeletedAt() *time.Time    { return a.state.DeletedAt }
func (a *Account) IsActive() bool {
	return a.state.AccountStatus == "active" && a.state.DeletedAt == nil
}
func (a *Account) State() AccountState { return a.state }
