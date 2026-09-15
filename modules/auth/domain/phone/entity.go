package phone

import (
	"time"

	"github.com/google/uuid"
)

type Phone struct{ state PhoneState }

type PhoneState struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Number     string
	IsPrimary  bool
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewFromState(st PhoneState) *Phone { return &Phone{state: st} }
func (p *Phone) ID() uuid.UUID          { return p.state.ID }
func (p *Phone) AccountID() uuid.UUID   { return p.state.AccountID }
func (p *Phone) Number() string         { return p.state.Number }
func (p *Phone) IsPrimary() bool        { return p.state.IsPrimary }
func (p *Phone) VerifiedAt() *time.Time { return p.state.VerifiedAt }
func (p *Phone) CreatedAt() time.Time   { return p.state.CreatedAt }
func (p *Phone) UpdatedAt() time.Time   { return p.state.UpdatedAt }
func (p *Phone) DeletedAt() *time.Time  { return p.state.DeletedAt }
func (p *Phone) State() PhoneState      { return p.state }

func (p *Phone) SetPrimary(v bool) {
	p.state.IsPrimary = v
	p.state.UpdatedAt = time.Now()
}
func (p *Phone) MarkVerified(at time.Time) {
	p.state.VerifiedAt = &at
	p.state.UpdatedAt = at
}
func (p *Phone) ChangeNumber(n string) {
	p.state.Number = n
	p.state.VerifiedAt = nil
	p.state.UpdatedAt = time.Now()
}
func (p *Phone) SoftDelete(at time.Time) {
	p.state.DeletedAt = &at
	p.state.UpdatedAt = at
}
