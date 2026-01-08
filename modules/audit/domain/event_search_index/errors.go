package event_search_index

import "errors"

var (
	ErrEventSearchIndexNotFound      = errors.New("event search index not found")
	ErrEventIDRequired               = errors.New("event id is required")
	ErrEventIDAlreadyExists          = errors.New("event id already exists")
	ErrActorTypeRequired             = errors.New("actor type is required")
	ErrActorIDRequired               = errors.New("actor id is required")
	ErrActionRequired                = errors.New("action is required")
	ErrResultRequired                = errors.New("result is required")
	ErrOccurredAtRequired            = errors.New("occurred at is required")
	ErrInvalidActorType              = errors.New("invalid actor type")
	ErrInvalidResultType             = errors.New("invalid result type")
	ErrInvalidRiskLevel              = errors.New("invalid risk level")
	ErrInvalidDataClassification     = errors.New("invalid data classification")
)
