package identity

import (
	"strings"
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"

	"gorm.io/datatypes"
)

type IdentityEditable struct {
	IdentityProvider enums.AuthIdentityProvider
	ProviderSubject  string
	PasswordHash     *string
	Metadata         *datatypes.JSON
}

func (i *Identity) EnsureNotDeleted() error {
	if i.DeletedAt() != nil {
		return authErr.ErrIdentityNotFound
	}
	return nil
}

func (i *Identity) UpdateProviderSubject(subject string) error {
	if err := i.EnsureNotDeleted(); err != nil {
		return err
	}
	subject = strings.TrimSpace(subject)
	if err := validateProviderSubject(subject); err != nil {
		return err
	}
	if i.ProviderSubject() == subject {
		return nil
	}
	i.state.ProviderSubject = subject
	i.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (i *Identity) Update(ed IdentityEditable) error {
	if err := i.EnsureNotDeleted(); err != nil {
		return err
	}
	ed.ProviderSubject = strings.TrimSpace(ed.ProviderSubject)
	if err := ed.Validate(); err != nil {
		return err
	}
	i.state.IdentityProvider = ed.IdentityProvider
	i.state.ProviderSubject = ed.ProviderSubject
	if ed.PasswordHash != nil {
		h := strings.TrimSpace(*ed.PasswordHash)
		i.state.PasswordHash = &h
	} else {
		i.state.PasswordHash = nil
	}
	i.state.Metadata = ed.Metadata
	i.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (i *Identity) UpdatePasswordHash(hashed string) error {
	if err := i.EnsureNotDeleted(); err != nil {
		return err
	}
	if i.IdentityProvider() != enums.AuthIdentityProviderPassword {
		return authErr.ErrIdentityProviderInvalid
	}
	if strings.TrimSpace(hashed) == "" {
		return authErr.ErrIdentityPasswordHashRequired
	}
	h := strings.TrimSpace(hashed)
	i.state.PasswordHash = &h
	i.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (i *Identity) RecordLogin(at time.Time) error {
	if err := i.EnsureNotDeleted(); err != nil {
		return err
	}
	t := at.UTC()
	i.state.LastLoginAt = &t
	i.state.UpdatedAt = t
	return nil
}

func (i *Identity) TouchLogin(at time.Time) {
	_ = i.RecordLogin(at)
}

func (i *Identity) SetPasswordHash(hash string) {
	_ = i.UpdatePasswordHash(hash)
}

func (i *Identity) SoftDelete(at time.Time) {
	i.state.DeletedAt = &at
	i.state.UpdatedAt = at
}

func (i *Identity) Delete() error {
	if i.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	i.state.DeletedAt = &now
	i.state.UpdatedAt = now
	return nil
}
