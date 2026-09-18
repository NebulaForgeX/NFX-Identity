package phone

import (
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"
)

type PhoneEditable struct {
	Phone     string
	IsPrimary bool
}

func (p *Phone) EnsureNotDeleted() error {
	if p.DeletedAt() != nil {
		return authErr.ErrPhoneBindingNotFound
	}
	return nil
}

func (p *Phone) Update(ed PhoneEditable) error {
	if err := p.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := ed.Validate(); err != nil {
		return err
	}
	p.state.Phone = strings.TrimSpace(ed.Phone)
	p.state.IsPrimary = ed.IsPrimary
	p.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (p *Phone) MarkVerified(at time.Time) error {
	if err := p.EnsureNotDeleted(); err != nil {
		return err
	}
	if p.state.VerifiedAt != nil {
		return nil
	}
	t := at.UTC()
	p.state.VerifiedAt = &t
	p.state.UpdatedAt = t
	return nil
}

func (p *Phone) SetPrimary(v bool) {
	p.state.IsPrimary = v
	p.state.UpdatedAt = time.Now().UTC()
}

func (p *Phone) ChangeNumber(n string) {
	p.state.Phone = strings.TrimSpace(n)
	p.state.VerifiedAt = nil
	p.state.UpdatedAt = time.Now().UTC()
}

func (p *Phone) SoftDelete(at time.Time) {
	p.state.DeletedAt = &at
	p.state.UpdatedAt = at
}

func (p *Phone) Delete() error {
	if p.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	p.state.DeletedAt = &now
	p.state.UpdatedAt = now
	return nil
}
