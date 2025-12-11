package profile

import (
	"time"

	profileErrors "nfxid/modules/auth/domain/profile/errors"
)

func (p *Profile) EnsureEditable(e ProfileEditable) error {
	if err := e.Validate(); err != nil {
		return err
	}
	if p.DeletedAt() != nil {
		return profileErrors.ErrProfileNotFound
	}
	return nil
}

func (p *Profile) Update(e ProfileEditable) error {
	if err := p.EnsureEditable(e); err != nil {
		return err
	}

	p.state.Editable = e
	p.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (p *Profile) Delete() error {
	if p.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	p.state.DeletedAt = &now
	p.state.UpdatedAt = now
	return nil
}
