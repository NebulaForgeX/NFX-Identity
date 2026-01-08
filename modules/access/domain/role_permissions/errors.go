package role_permissions

import "errors"

var (
	ErrRolePermissionNotFound        = errors.New("role permission not found")
	ErrRolePermissionAlreadyExists   = errors.New("role permission already exists")
	ErrRoleIDRequired                = errors.New("role id is required")
	ErrPermissionIDRequired          = errors.New("permission id is required")
)
