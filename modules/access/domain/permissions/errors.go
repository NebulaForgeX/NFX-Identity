package permissions

import "errors"

var (
	ErrPermissionNotFound      = errors.New("permission not found")
	ErrPermissionKeyRequired   = errors.New("permission key is required")
	ErrPermissionNameRequired  = errors.New("permission name is required")
	ErrPermissionKeyExists     = errors.New("permission key already exists")
	ErrSystemPermissionDelete  = errors.New("cannot delete system permission")
)
