package directory

import "nfxidentity/pkgs/errx"

const (
	CodeUserEducationNotFound = "USER_EDUCATION_NOT_FOUND"
	CodeSchoolRequired        = "SCHOOL_REQUIRED"
)

var (
	ErrUserEducationNotFound = errx.NotFound(CodeUserEducationNotFound, "user education not found")
	ErrSchoolRequired        = errx.InvalidArg(CodeSchoolRequired, "school is required")
)

/*
!USER_EDUCATION_NOT_FOUND
*en<user education not found>
*zh<用户教育经历不存在>
*fr<formation utilisateur introuvable>

!SCHOOL_REQUIRED
*en<school required>
*zh<学校为必填>
*fr<école requise>

*/
