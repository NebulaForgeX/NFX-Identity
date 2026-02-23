package user_occupations

import "nfxid/pkgs/errx"

var (
	ErrUserOccupationNotFound = errx.NotFound("USER_OCCUPATION_NOT_FOUND", "user occupation not found")
	ErrUserIDRequired         = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrCompanyRequired        = errx.InvalidArg("COMPANY_REQUIRED", "company is required")
	ErrPositionRequired       = errx.InvalidArg("POSITION_REQUIRED", "position is required")
)
