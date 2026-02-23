package badges

import "nfxid/pkgs/errx"

var (
	ErrBadgeNotFound     = errx.NotFound("BADGE_NOT_FOUND", "badge not found")
	ErrNameRequired      = errx.InvalidArg("NAME_REQUIRED", "name is required")
	ErrNameAlreadyExists = errx.Conflict("NAME_ALREADY_EXISTS", "name already exists")
)
