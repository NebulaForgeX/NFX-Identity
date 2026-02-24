package system

import "nfxid/pkgs/errx"

const (
	CodeSystemStateNotFound      = "SYSTEM_STATE_NOT_FOUND"
	CodeSystemAlreadyInitialized = "SYSTEM_ALREADY_INITIALIZED"
	CodeSystemNotInitialized     = "SYSTEM_NOT_INITIALIZED"
)

var (
	ErrSystemStateNotFound = errx.NotFound(CodeSystemStateNotFound, "system state not found")
	ErrAlreadyInitialized  = errx.Conflict(CodeSystemAlreadyInitialized, "system already initialized")
	ErrNotInitialized      = errx.FailedPrecond(CodeSystemNotInitialized, "system not initialized")
)
