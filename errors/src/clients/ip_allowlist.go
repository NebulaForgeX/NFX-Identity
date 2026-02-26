package clients

import "nfxidentity/pkgs/errx"

const (
	CodeIPAllowlistNotFound       = "IP_ALLOWLIST_NOT_FOUND"
	CodeRuleIDRequired            = "RULE_ID_REQUIRED"
	CodeCIDRRequired              = "CIDR_REQUIRED"
	CodeRuleIDAlreadyExists       = "RULE_ID_ALREADY_EXISTS"
	CodeInvalidAllowlistStatus    = "INVALID_ALLOWLIST_STATUS"
	CodeIPAllowlistAlreadyRevoked = "IP_ALLOWLIST_ALREADY_REVOKED"
)

var (
	ErrIPAllowlistNotFound       = errx.NotFound(CodeIPAllowlistNotFound, "ip allowlist not found")
	ErrRuleIDRequired            = errx.InvalidArg(CodeRuleIDRequired, "rule id is required")
	ErrCIDRRequired              = errx.InvalidArg(CodeCIDRRequired, "cidr is required")
	ErrRuleIDAlreadyExists       = errx.Conflict(CodeRuleIDAlreadyExists, "rule id already exists")
	ErrInvalidAllowlistStatus    = errx.InvalidArg(CodeInvalidAllowlistStatus, "invalid allowlist status")
	ErrIPAllowlistAlreadyRevoked = errx.Conflict(CodeIPAllowlistAlreadyRevoked, "ip allowlist already revoked")
)

/*
!IP_ALLOWLIST_NOT_FOUND
*en<ip allowlist not found>
*zh<IP 允许列表不存在>
*fr<liste d'autorisation IP introuvable>

!RULE_ID_REQUIRED
*en<rule id required>
*zh<规则 ID 为必填>
*fr<id de règle requis>

!CIDR_REQUIRED
*en<cidr required>
*zh<CIDR 为必填>
*fr<CIDR requis>

!RULE_ID_ALREADY_EXISTS
*en<rule id already exists>
*zh<规则 ID 已存在>
*fr<id de règle existe déjà>

!INVALID_ALLOWLIST_STATUS
*en<invalid allowlist status>
*zh<无效的允许列表状态>
*fr<statut de liste d'autorisation invalide>

!IP_ALLOWLIST_ALREADY_REVOKED
*en<ip allowlist already revoked>
*zh<IP 允许列表已撤销>
*fr<liste d'autorisation IP déjà révoquée>

*/
