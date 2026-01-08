package apps

import "errors"

var (
	ErrAppNotFound       = errors.New("app not found")
	ErrAppIDRequired     = errors.New("app id is required")
	ErrNameRequired      = errors.New("name is required")
	ErrTenantIDRequired  = errors.New("tenant id is required")
	ErrAppIDAlreadyExists = errors.New("app id already exists")
	ErrInvalidAppType    = errors.New("invalid app type")
	ErrInvalidAppStatus  = errors.New("invalid app status")
	ErrInvalidEnvironment = errors.New("invalid environment")
)
