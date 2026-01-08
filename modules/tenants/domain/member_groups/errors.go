package member_groups

import "errors"

var (
	ErrMemberGroupNotFound     = errors.New("member group not found")
	ErrMemberIDRequired        = errors.New("member id is required")
	ErrGroupIDRequired         = errors.New("group id is required")
	ErrMemberGroupAlreadyExists = errors.New("member group already exists")
	ErrMemberGroupAlreadyRevoked = errors.New("member group already revoked")
)
