package groups

import "nfxid/pkgs/errx"

var (
	ErrGroupNotFound       = errx.NotFound("GROUP_NOT_FOUND", "group not found")
	ErrGroupIDRequired     = errx.InvalidArg("GROUP_ID_REQUIRED", "group id is required")
	ErrNameRequired        = errx.InvalidArg("NAME_REQUIRED", "name is required")
	ErrTenantIDRequired     = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrGroupIDAlreadyExists = errx.Conflict("GROUP_ID_ALREADY_EXISTS", "group id already exists")
	ErrInvalidGroupType    = errx.InvalidArg("INVALID_GROUP_TYPE", "invalid group type")
)
