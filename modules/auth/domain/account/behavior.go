package account

import (
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
)

func (u *Account) EnsureNotDeleted() error {
	if u.DeletedAt() != nil {
		return authErr.ErrAccountNotFound
	}
	return nil
}

func (u *Account) UpdateStatus(status enums.AuthAccountStatus) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := validateAccountStatus(status); err != nil {
		return err
	}
	u.state.AccountStatus = status
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *Account) Delete() error {
	if u.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	u.state.DeletedAt = &now
	u.state.UpdatedAt = now
	return nil
}
