package members

import "nfxid/pkgs/errx"

var (
	ErrMemberNotFound      = errx.NotFound("MEMBER_NOT_FOUND", "member not found")
	ErrTenantIDRequired    = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrUserIDRequired      = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrMemberAlreadyExists = errx.Conflict("MEMBER_ALREADY_EXISTS", "member already exists")
	ErrInvalidMemberStatus = errx.InvalidArg("INVALID_MEMBER_STATUS", "invalid member status")
	ErrInvalidMemberSource = errx.InvalidArg("INVALID_MEMBER_SOURCE", "invalid member source")
)
