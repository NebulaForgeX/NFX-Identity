package audit

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple audit domains.
const (
	CodeEventIDRequired           = "EVENT_ID_REQUIRED"
	CodeEventIDAlreadyExists      = "EVENT_ID_ALREADY_EXISTS"
	CodeOccurredAtRequired        = "OCCURRED_AT_REQUIRED"
	CodeActorTypeRequired         = "ACTOR_TYPE_REQUIRED"
	CodeActorIDRequired           = "ACTOR_ID_REQUIRED"
	CodeActionRequired            = "ACTION_REQUIRED"
	CodeResultRequired            = "RESULT_REQUIRED"
	CodeInvalidActorType          = "INVALID_ACTOR_TYPE"
	CodeInvalidResultType         = "INVALID_RESULT_TYPE"
	CodeInvalidRiskLevel          = "INVALID_RISK_LEVEL"
	CodeInvalidDataClassification = "INVALID_DATA_CLASSIFICATION"
)

var (
	ErrEventIDRequired           = errx.InvalidArg(CodeEventIDRequired, "event id is required")
	ErrEventIDAlreadyExists      = errx.Conflict(CodeEventIDAlreadyExists, "event id already exists")
	ErrOccurredAtRequired        = errx.InvalidArg(CodeOccurredAtRequired, "occurred at is required")
	ErrActorTypeRequired         = errx.InvalidArg(CodeActorTypeRequired, "actor type is required")
	ErrActorIDRequired           = errx.InvalidArg(CodeActorIDRequired, "actor id is required")
	ErrActionRequired            = errx.InvalidArg(CodeActionRequired, "action is required")
	ErrResultRequired            = errx.InvalidArg(CodeResultRequired, "result is required")
	ErrInvalidActorType          = errx.InvalidArg(CodeInvalidActorType, "invalid actor type")
	ErrInvalidResultType         = errx.InvalidArg(CodeInvalidResultType, "invalid result type")
	ErrInvalidRiskLevel          = errx.InvalidArg(CodeInvalidRiskLevel, "invalid risk level")
	ErrInvalidDataClassification = errx.InvalidArg(CodeInvalidDataClassification, "invalid data classification")
)
