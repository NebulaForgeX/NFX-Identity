package email

import (
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"
)

type EmailEditable struct {
	Email     string
	IsPrimary bool
}

func (e *Email) EnsureNotDeleted() error {
	if e.DeletedAt() != nil {
		return authErr.ErrEmailBindingNotFound
	}
	return nil
}

func (e *Email) Update(ed EmailEditable) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := ed.Validate(); err != nil {
		return err
	}
	e.state.Email = strings.TrimSpace(ed.Email)
	e.state.IsPrimary = ed.IsPrimary
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

// MarkVerified sets verified_at when not already verified.
func (e *Email) MarkVerified(at time.Time) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if e.state.VerifiedAt != nil {
		return nil
	}
	t := at.UTC()
	e.state.VerifiedAt = &t
	e.state.UpdatedAt = t
	return nil
}

func (e *Email) Delete() error {
	if e.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	e.state.DeletedAt = &now
	e.state.UpdatedAt = now
	return nil
}

func (e *Email) UpdateAddress(addr string) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := validateEmailAddress(addr); err != nil {
		return err
	}
	e.state.Email = strings.TrimSpace(addr)
	e.state.VerifiedAt = nil
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Email) SetPrimaryFlag(primary bool) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	e.state.IsPrimary = primary
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Email) SetPrimary(v bool) {
	_ = e.SetPrimaryFlag(v)
}

func (e *Email) SoftDelete(at time.Time) {
	e.state.DeletedAt = &at
	e.state.UpdatedAt = at
}

func (e *Email) ChangeAddress(address string) {
	_ = e.UpdateAddress(address)
}
