package user_occupations

import "github.com/google/uuid"

func (uo *UserOccupation) Validate() error {
	if uo.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if uo.Company() == "" {
		return ErrCompanyRequired
	}
	if uo.Position() == "" {
		return ErrPositionRequired
	}
	return nil
}
