package phone

import (
	"time"

	"github.com/google/uuid"
)

type Phone struct {
	state PhoneState
}

type PhoneState struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Phone      string
	IsPrimary  bool
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (p *Phone) ID() uuid.UUID          { return p.state.ID }
func (p *Phone) AccountID() uuid.UUID   { return p.state.AccountID }
func (p *Phone) Phone() string          { return p.state.Phone }
func (p *Phone) IsPrimary() bool        { return p.state.IsPrimary }
func (p *Phone) VerifiedAt() *time.Time { return p.state.VerifiedAt }
func (p *Phone) CreatedAt() time.Time   { return p.state.CreatedAt }
func (p *Phone) UpdatedAt() time.Time   { return p.state.UpdatedAt }
func (p *Phone) DeletedAt() *time.Time  { return p.state.DeletedAt }
func (p *Phone) IsVerified() bool       { return p.state.VerifiedAt != nil }

func NewPhoneFromState(st PhoneState) *Phone {
	return &Phone{state: st}
}
