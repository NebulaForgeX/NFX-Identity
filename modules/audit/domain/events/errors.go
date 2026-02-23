package events

import "nfxid/pkgs/errx"

var (
	ErrEventNotFound            = errx.NotFound("EVENT_NOT_FOUND", "event not found")
	ErrEventIDRequired          = errx.InvalidArg("EVENT_ID_REQUIRED", "event id is required")
	ErrEventIDAlreadyExists     = errx.Conflict("EVENT_ID_ALREADY_EXISTS", "event id already exists")
	ErrOccurredAtRequired       = errx.InvalidArg("OCCURRED_AT_REQUIRED", "occurred at is required")
	ErrActorTypeRequired        = errx.InvalidArg("ACTOR_TYPE_REQUIRED", "actor type is required")
	ErrActorIDRequired          = errx.InvalidArg("ACTOR_ID_REQUIRED", "actor id is required")
	ErrActionRequired           = errx.InvalidArg("ACTION_REQUIRED", "action is required")
	ErrResultRequired           = errx.InvalidArg("RESULT_REQUIRED", "result is required")
	ErrInvalidActorType         = errx.InvalidArg("INVALID_ACTOR_TYPE", "invalid actor type")
	ErrInvalidResultType        = errx.InvalidArg("INVALID_RESULT_TYPE", "invalid result type")
	ErrInvalidRiskLevel         = errx.InvalidArg("INVALID_RISK_LEVEL", "invalid risk level")
	ErrInvalidDataClassification = errx.InvalidArg("INVALID_DATA_CLASSIFICATION", "invalid data classification")
)
