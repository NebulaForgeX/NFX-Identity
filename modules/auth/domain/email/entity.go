package email

import (
	"time"

	"github.com/google/uuid"
)

type Email struct {
	state EmailState
}

type EmailState struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Email      string
	IsPrimary  bool
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (e *Email) ID() uuid.UUID          { return e.state.ID }
func (e *Email) AccountID() uuid.UUID   { return e.state.AccountID }
func (e *Email) Email() string          { return e.state.Email }
func (e *Email) IsPrimary() bool        { return e.state.IsPrimary }
func (e *Email) VerifiedAt() *time.Time { return e.state.VerifiedAt }
func (e *Email) CreatedAt() time.Time   { return e.state.CreatedAt }
func (e *Email) UpdatedAt() time.Time   { return e.state.UpdatedAt }
func (e *Email) DeletedAt() *time.Time  { return e.state.DeletedAt }
func (e *Email) IsVerified() bool       { return e.state.VerifiedAt != nil }

func NewEmailFromState(st EmailState) *Email {
	return &Email{state: st}
}
