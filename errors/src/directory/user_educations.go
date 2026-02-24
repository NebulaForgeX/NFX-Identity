package directory

import "nfxid/pkgs/errx"

const (
	CodeUserEducationNotFound = "USER_EDUCATION_NOT_FOUND"
	CodeSchoolRequired        = "SCHOOL_REQUIRED"
)

var (
	ErrUserEducationNotFound = errx.NotFound(CodeUserEducationNotFound, "user education not found")
	ErrSchoolRequired        = errx.InvalidArg(CodeSchoolRequired, "school is required")
)
