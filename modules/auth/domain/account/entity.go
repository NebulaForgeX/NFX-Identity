package account

import (
	"time"

	"nfxidentity/enums"

	"github.com/google/uuid"
)

type Account struct {
	state AccountState
}

type AccountState struct {
	ID             uuid.UUID
	AccountStatus  enums.AuthAccountStatus
	SignupPlatform enums.AuthSignupPlatform
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func (u *Account) ID() uuid.UUID                            { return u.state.ID }
func (u *Account) AccountStatus() enums.AuthAccountStatus   { return u.state.AccountStatus }
func (u *Account) SignupPlatform() enums.AuthSignupPlatform { return u.state.SignupPlatform }
func (u *Account) CreatedAt() time.Time                     { return u.state.CreatedAt }
func (u *Account) UpdatedAt() time.Time                     { return u.state.UpdatedAt }
func (u *Account) DeletedAt() *time.Time                    { return u.state.DeletedAt }

func (u *Account) IsActive() bool {
	return u.AccountStatus() == enums.AuthAccountStatusActive && u.DeletedAt() == nil
}

func NewAccountFromState(st AccountState) *Account {
	return &Account{state: st}
}
