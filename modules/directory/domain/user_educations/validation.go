package user_educations

import "github.com/google/uuid"

func (ue *UserEducation) Validate() error {
	if ue.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if ue.School() == "" {
		return ErrSchoolRequired
	}
	return nil
}
