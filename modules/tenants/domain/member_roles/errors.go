package member_roles

import "errors"

var (
	ErrMemberRoleNotFound     = errors.New("member role not found")
	ErrTenantIDRequired       = errors.New("tenant id is required")
	ErrMemberIDRequired       = errors.New("member id is required")
	ErrRoleIDRequired         = errors.New("role id is required")
	ErrMemberRoleAlreadyExists = errors.New("member role already exists")
	ErrMemberRoleAlreadyRevoked = errors.New("member role already revoked")
	ErrMemberRoleExpired      = errors.New("member role expired")
)
