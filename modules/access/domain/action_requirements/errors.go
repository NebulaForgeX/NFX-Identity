package action_requirements

import "errors"

var (
	ErrActionRequirementNotFound      = errors.New("action requirement not found")
	ErrActionIDRequired               = errors.New("action id is required")
	ErrPermissionIDRequired           = errors.New("permission id is required")
	ErrActionRequirementAlreadyExists = errors.New("action requirement already exists for this action and permission")
)
