package users

func (u *User) Validate() error {
	if u.Username() == "" {
		return ErrUsernameRequired
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
