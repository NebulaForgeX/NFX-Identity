package errors

import "errors"

var (
	ErrPermissionNotFound            = errors.New("permission not found")
	ErrPermissionTagRequired         = errors.New("permission tag is required")
	ErrPermissionTagInvalid         = errors.New("permission tag is invalid")
	ErrPermissionNameRequired        = errors.New("permission name is required")
	ErrPermissionNameInvalid         = errors.New("permission name is invalid")
	ErrPermissionTagAlreadyExists    = errors.New("permission tag already exists")
	ErrPermissionSystemCannotModify  = errors.New("system permission cannot be modified")
	ErrPermissionSystemCannotDelete  = errors.New("system permission cannot be deleted")
)

