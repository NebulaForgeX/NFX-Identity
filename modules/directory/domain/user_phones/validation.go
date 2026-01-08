package user_phones

import "github.com/google/uuid"

func (up *UserPhone) Validate() error {
	if up.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if up.Phone() == "" {
		return ErrPhoneRequired
	}
	return nil
}
