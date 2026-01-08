package event_retention_policies

import "errors"

var (
	ErrEventRetentionPolicyNotFound      = errors.New("event retention policy not found")
	ErrPolicyNameRequired                = errors.New("policy name is required")
	ErrRetentionDaysRequired             = errors.New("retention days is required")
	ErrRetentionActionRequired           = errors.New("retention action is required")
	ErrPolicyNameAlreadyExists           = errors.New("policy name already exists")
	ErrInvalidRetentionAction            = errors.New("invalid retention action")
	ErrInvalidStatus                     = errors.New("invalid status")
)
