package members

import "errors"

var (
	ErrMemberNotFound       = errors.New("member not found")
	ErrTenantIDRequired     = errors.New("tenant id is required")
	ErrUserIDRequired       = errors.New("user id is required")
	ErrMemberAlreadyExists  = errors.New("member already exists")
	ErrInvalidMemberStatus  = errors.New("invalid member status")
	ErrInvalidMemberSource  = errors.New("invalid member source")
)
