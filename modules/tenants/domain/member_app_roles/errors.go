package member_app_roles

import "errors"

var (
	ErrMemberAppRoleNotFound     = errors.New("member app role not found")
	ErrMemberIDRequired          = errors.New("member id is required")
	ErrAppIDRequired             = errors.New("app id is required")
	ErrRoleIDRequired            = errors.New("role id is required")
	ErrMemberAppRoleAlreadyExists = errors.New("member app role already exists")
	ErrMemberAppRoleAlreadyRevoked = errors.New("member app role already revoked")
	ErrMemberAppRoleExpired      = errors.New("member app role expired")
)
