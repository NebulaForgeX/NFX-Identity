package audit

import "nfxidentity/pkgs/errx"

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

/*
!EVENT_ID_REQUIRED
*en<event id required>
*zh<事件 ID 为必填>
*fr<id d'événement requis>

!EVENT_ID_ALREADY_EXISTS
*en<event id already exists>
*zh<事件 ID 已存在>
*fr<id d'événement existe déjà>

!OCCURRED_AT_REQUIRED
*en<occurred at required>
*zh<发生时间为必填>
*fr<date d'occurrence requise>

!ACTOR_TYPE_REQUIRED
*en<actor type required>
*zh<操作者类型为必填>
*fr<type d'acteur requis>

!ACTOR_ID_REQUIRED
*en<actor id required>
*zh<操作者 ID 为必填>
*fr<id d'acteur requis>

!ACTION_REQUIRED
*en<action required>
*zh<需要操作>
*fr<action requise>

!RESULT_REQUIRED
*en<result required>
*zh<结果为必填>
*fr<résultat requis>

!INVALID_ACTOR_TYPE
*en<invalid actor type>
*zh<无效的操作者类型>
*fr<type d'acteur invalide>

!INVALID_RESULT_TYPE
*en<invalid result type>
*zh<无效的结果类型>
*fr<type de résultat invalide>

!INVALID_RISK_LEVEL
*en<invalid risk level>
*zh<无效的风险等级>
*fr<niveau de risque invalide>

!INVALID_DATA_CLASSIFICATION
*en<invalid data classification>
*zh<无效的数据分类>
*fr<classification des données invalide>

*/
