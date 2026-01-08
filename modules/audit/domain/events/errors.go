package events

import "errors"

var (
	ErrEventNotFound        = errors.New("event not found")
	ErrEventIDRequired      = errors.New("event id is required")
	ErrEventIDAlreadyExists = errors.New("event id already exists")
	ErrOccurredAtRequired   = errors.New("occurred at is required")
	ErrActorTypeRequired    = errors.New("actor type is required")
	ErrActorIDRequired      = errors.New("actor id is required")
	ErrActionRequired       = errors.New("action is required")
	ErrResultRequired       = errors.New("result is required")
	ErrInvalidActorType     = errors.New("invalid actor type")
	ErrInvalidResultType    = errors.New("invalid result type")
	ErrInvalidRiskLevel     = errors.New("invalid risk level")
	ErrInvalidDataClassification = errors.New("invalid data classification")
)
