package email

import (
	"time"

	"github.com/google/uuid"
)

type Email struct{ state EmailState }

type EmailState struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Address    string
	IsPrimary  bool
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewFromState(st EmailState) *Email { return &Email{state: st} }
func (e *Email) ID() uuid.UUID          { return e.state.ID }
func (e *Email) AccountID() uuid.UUID   { return e.state.AccountID }
func (e *Email) Address() string        { return e.state.Address }
func (e *Email) IsPrimary() bool        { return e.state.IsPrimary }
func (e *Email) VerifiedAt() *time.Time { return e.state.VerifiedAt }
func (e *Email) CreatedAt() time.Time   { return e.state.CreatedAt }
func (e *Email) UpdatedAt() time.Time   { return e.state.UpdatedAt }
func (e *Email) DeletedAt() *time.Time  { return e.state.DeletedAt }
func (e *Email) State() EmailState      { return e.state }
func (e *Email) IsVerified() bool       { return e.state.VerifiedAt != nil && e.state.DeletedAt == nil }

func (e *Email) SetPrimary(v bool) {
	e.state.IsPrimary = v
	e.state.UpdatedAt = time.Now()
}
func (e *Email) MarkVerified(at time.Time) {
	e.state.VerifiedAt = &at
	e.state.UpdatedAt = at
}
func (e *Email) ChangeAddress(address string) {
	e.state.Address = address
	e.state.VerifiedAt = nil
	e.state.UpdatedAt = time.Now()
}
func (e *Email) SoftDelete(at time.Time) {
	e.state.DeletedAt = &at
	e.state.UpdatedAt = at
}
