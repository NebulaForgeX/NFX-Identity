package user_emails

import "github.com/google/uuid"

func (ue *UserEmail) Validate() error {
	if ue.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if ue.Email() == "" {
		return ErrEmailRequired
	}
	return nil
}
