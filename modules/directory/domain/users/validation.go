package users

import (
	dirErr "nfxidentity/errors/src/directory"
)

func (u *User) Validate() error {
	if u.Username() == "" {
		return dirErr.ErrUsernameRequired
	}
	validStatuses := map[UserStatus]struct{}{
		UserStatusPending:  {},
		UserStatusActive:   {},
		UserStatusDeactive: {},
	}
	if _, ok := validStatuses[u.Status()]; !ok {
		return dirErr.ErrInvalidUserStatus
	}
	return nil
}
