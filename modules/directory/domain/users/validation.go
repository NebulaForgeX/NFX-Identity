package users

import "github.com/google/uuid"

func (u *User) Validate() error {
	if u.Username() == "" {
		return ErrUsernameRequired
	}
	if u.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	validStatuses := map[UserStatus]struct{}{
		UserStatusPending:  {},
		UserStatusActive:   {},
		UserStatusDeactive: {},
	}
	if _, ok := validStatuses[u.Status()]; !ok {
		return ErrInvalidUserStatus
	}
	return nil
}
