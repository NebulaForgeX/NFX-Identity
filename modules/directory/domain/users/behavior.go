package users

import (
	"time"
)

func (u *User) UpdateStatus(status UserStatus) error {
	if u.DeletedAt() != nil {
		return ErrUserNotFound
	}
	validStatuses := map[UserStatus]struct{}{
		UserStatusPending:  {},
		UserStatusActive:   {},
		UserStatusDeactive: {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidUserStatus
	}

	u.state.Status = status
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdateLastLogin() error {
	if u.DeletedAt() != nil {
		return ErrUserNotFound
	}

	now := time.Now().UTC()
	u.state.LastLoginAt = &now
	u.state.UpdatedAt = now
	return nil
}

func (u *User) Verify() error {
	if u.DeletedAt() != nil {
		return ErrUserNotFound
	}

	u.state.IsVerified = true
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdateUsername(username string) error {
	if u.DeletedAt() != nil {
		return ErrUserNotFound
	}
	if username == "" {
		return ErrUsernameRequired
	}

	u.state.Username = username
	u.state.UpdatedAt = time.Now().UTC()
	return nil
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
