package audit

import "nfxid/pkgs/errx"

const (
	CodeEventRetentionPolicyNotFound = "EVENT_RETENTION_POLICY_NOT_FOUND"
	CodePolicyNameRequired           = "POLICY_NAME_REQUIRED"
	CodeRetentionDaysRequired        = "RETENTION_DAYS_REQUIRED"
	CodeRetentionActionRequired      = "RETENTION_ACTION_REQUIRED"
	CodePolicyNameAlreadyExists      = "POLICY_NAME_ALREADY_EXISTS"
	CodeInvalidRetentionAction       = "INVALID_RETENTION_ACTION"
	CodeInvalidStatus                = "INVALID_STATUS"
)

var (
	ErrEventRetentionPolicyNotFound = errx.NotFound(CodeEventRetentionPolicyNotFound, "event retention policy not found")
	ErrPolicyNameRequired           = errx.InvalidArg(CodePolicyNameRequired, "policy name is required")
	ErrRetentionDaysRequired        = errx.InvalidArg(CodeRetentionDaysRequired, "retention days is required")
	ErrRetentionActionRequired      = errx.InvalidArg(CodeRetentionActionRequired, "retention action is required")
	ErrPolicyNameAlreadyExists      = errx.Conflict(CodePolicyNameAlreadyExists, "policy name already exists")
	ErrInvalidRetentionAction       = errx.InvalidArg(CodeInvalidRetentionAction, "invalid retention action")
	ErrInvalidStatus                = errx.InvalidArg(CodeInvalidStatus, "invalid status")
)

/*
!EVENT_RETENTION_POLICY_NOT_FOUND
*en<event retention policy not found>
*zh<事件保留策略不存在>
*fr<politique de rétention d'événement introuvable>

!POLICY_NAME_REQUIRED
*en<policy name required>
*zh<策略名称为必填>
*fr<nom de politique requis>

!RETENTION_DAYS_REQUIRED
*en<retention days required>
*zh<保留天数为必填>
*fr<jours de rétention requis>

!RETENTION_ACTION_REQUIRED
*en<retention action required>
*zh<保留操作为必填>
*fr<action de rétention requise>

!POLICY_NAME_ALREADY_EXISTS
*en<policy name already exists>
*zh<策略名称已存在>
*fr<nom de politique existe déjà>

!INVALID_RETENTION_ACTION
*en<invalid retention action>
*zh<无效的保留操作>
*fr<action de rétention invalide>

!INVALID_STATUS
*en<invalid status>
*zh<无效的状态>
*fr<statut invalide>

*/
