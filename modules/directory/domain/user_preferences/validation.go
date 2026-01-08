package user_preferences

import "github.com/google/uuid"

func (up *UserPreference) Validate() error {
	if up.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	return nil
}
