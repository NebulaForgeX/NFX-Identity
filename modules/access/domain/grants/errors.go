package grants

import "errors"

var (
	ErrGrantNotFound       = errors.New("grant not found")
	ErrSubjectTypeRequired = errors.New("subject type is required")
	ErrSubjectIDRequired   = errors.New("subject id is required")
	ErrGrantTypeRequired   = errors.New("grant type is required")
	ErrGrantRefIDRequired  = errors.New("grant ref id is required")
	ErrInvalidSubjectType  = errors.New("invalid subject type")
	ErrInvalidGrantType    = errors.New("invalid grant type")
	ErrInvalidGrantEffect  = errors.New("invalid grant effect")
	ErrGrantAlreadyRevoked = errors.New("grant already revoked")
)
