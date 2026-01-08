package scope_permissions

import "errors"

var (
	ErrScopePermissionNotFound      = errors.New("scope permission not found")
	ErrScopePermissionAlreadyExists = errors.New("scope permission already exists")
	ErrScopeRequired                = errors.New("scope is required")
	ErrPermissionIDRequired         = errors.New("permission id is required")
)
