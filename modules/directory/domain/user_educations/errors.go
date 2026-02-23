package user_educations

import "nfxid/pkgs/errx"

var (
	ErrUserEducationNotFound = errx.NotFound("USER_EDUCATION_NOT_FOUND", "user education not found")
	ErrUserIDRequired         = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrSchoolRequired         = errx.InvalidArg("SCHOOL_REQUIRED", "school is required")
)
