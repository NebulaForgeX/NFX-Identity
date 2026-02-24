package directory

import "nfxid/pkgs/errx"

const (
	CodeUserOccupationNotFound = "USER_OCCUPATION_NOT_FOUND"
	CodeCompanyRequired        = "COMPANY_REQUIRED"
	CodePositionRequired       = "POSITION_REQUIRED"
)

var (
	ErrUserOccupationNotFound = errx.NotFound(CodeUserOccupationNotFound, "user occupation not found")
	ErrCompanyRequired        = errx.InvalidArg(CodeCompanyRequired, "company is required")
	ErrPositionRequired       = errx.InvalidArg(CodePositionRequired, "position is required")
)

/*
!USER_OCCUPATION_NOT_FOUND
*en<user occupation not found>
*zh<用户职业经历不存在>
*fr<profession utilisateur introuvable>

!COMPANY_REQUIRED
*en<company required>
*zh<公司为必填>
*fr<entreprise requise>

!POSITION_REQUIRED
*en<position required>
*zh<职位为必填>
*fr<poste requis>

*/
