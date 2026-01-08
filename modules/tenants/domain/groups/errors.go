package groups

import "errors"

var (
	ErrGroupNotFound      = errors.New("group not found")
	ErrGroupIDRequired    = errors.New("group id is required")
	ErrNameRequired       = errors.New("name is required")
	ErrTenantIDRequired   = errors.New("tenant id is required")
	ErrGroupIDAlreadyExists = errors.New("group id already exists")
	ErrInvalidGroupType   = errors.New("invalid group type")
)
