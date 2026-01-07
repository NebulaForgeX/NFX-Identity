package errors

import "errors"

var (
	ErrUserPermissionNotFound            = errors.New("user permission not found")
	ErrUserPermissionUserIDRequired      = errors.New("user id is required")
	ErrUserPermissionPermissionIDRequired = errors.New("permission id is required")
	ErrUserPermissionAlreadyExists       = errors.New("user permission already exists")
)

