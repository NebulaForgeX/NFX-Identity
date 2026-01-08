package user_profiles

import "github.com/google/uuid"

func (up *UserProfile) Validate() error {
	if up.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	return nil
}
