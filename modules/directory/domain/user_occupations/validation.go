package user_occupations

import (
	dirErr "nfxid/errors/src/directory"

	"github.com/google/uuid"
)

func (uo *UserOccupation) Validate() error {
	if uo.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if uo.Company() == "" {
		return dirErr.ErrCompanyRequired
	}
	if uo.Position() == "" {
		return dirErr.ErrPositionRequired
	}
	return nil
}
