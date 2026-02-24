package user_educations

import (
	dirErr "nfxid/errors/src/directory"

	"github.com/google/uuid"
)

func (ue *UserEducation) Validate() error {
	if ue.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if ue.School() == "" {
		return dirErr.ErrSchoolRequired
	}
	return nil
}
