package system_state

import "nfxid/pkgs/errx"

var (
	ErrSystemStateNotFound = errx.NotFound("SYSTEM_STATE_NOT_FOUND", "system state not found")
	ErrAlreadyInitialized  = errx.Conflict("SYSTEM_ALREADY_INITIALIZED", "system already initialized")
	ErrNotInitialized      = errx.FailedPrecond("SYSTEM_NOT_INITIALIZED", "system not initialized")
)
