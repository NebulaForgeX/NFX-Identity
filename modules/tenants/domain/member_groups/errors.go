package member_groups

import "nfxid/pkgs/errx"

var (
	ErrMemberGroupNotFound     = errx.NotFound("MEMBER_GROUP_NOT_FOUND", "member group not found")
	ErrMemberIDRequired        = errx.InvalidArg("MEMBER_ID_REQUIRED", "member id is required")
	ErrGroupIDRequired         = errx.InvalidArg("GROUP_ID_REQUIRED", "group id is required")
	ErrMemberGroupAlreadyExists = errx.Conflict("MEMBER_GROUP_ALREADY_EXISTS", "member group already exists")
	ErrMemberGroupAlreadyRevoked = errx.Conflict("MEMBER_GROUP_ALREADY_REVOKED", "member group already revoked")
)
