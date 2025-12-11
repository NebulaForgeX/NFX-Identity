package user

import (
	"nfxid/modules/auth/domain/user/errors"
	"time"
)

func (u *User) EnsureEditable(e UserEditable) error {
	if err := e.Validate(); err != nil {
		return err
	}
	if u.DeletedAt() != nil {
		return errors.ErrUserNotFound
	}
	return nil
}

func (u *User) Update(e UserEditable) error {
	if err := u.EnsureEditable(e); err != nil {
		return err
	}

	u.state.Editable = e
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdatePassword(hashedPassword string) error {
	if u.DeletedAt() != nil {
		return errors.ErrUserNotFound
	}
	if hashedPassword == "" {
		return errors.ErrUserPasswordRequired
	}

	u.state.Editable.Password = hashedPassword
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdateStatus(status string) error {
	if u.DeletedAt() != nil {
		return errors.ErrUserNotFound
	}
	validStatuses := map[string]struct{}{
		"pending":  {},
		"active":   {},
		"deactive": {},
	}
	if _, ok := validStatuses[status]; !ok {
		return errors.ErrUserInactive
	}

	u.state.Status = status
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) Verify() {
	u.state.IsVerified = true
	u.state.UpdatedAt = time.Now().UTC()
}

func (u *User) UpdateLastLogin() {
	now := time.Now().UTC()
	u.state.LastLoginAt = &now
	u.state.UpdatedAt = now
}

func (u *User) Delete() error {
	if u.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	u.state.DeletedAt = &now
	u.state.UpdatedAt = now
	return nil
}
