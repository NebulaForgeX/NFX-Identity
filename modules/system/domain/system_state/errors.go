package system_state

import "errors"

var (
	ErrSystemStateNotFound = errors.New("system state not found")
	ErrAlreadyInitialized  = errors.New("system already initialized")
	ErrNotInitialized      = errors.New("system not initialized")
)
