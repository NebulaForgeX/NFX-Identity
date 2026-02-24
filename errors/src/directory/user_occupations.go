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
