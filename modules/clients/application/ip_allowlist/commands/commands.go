package commands

import (
	"nfxid/modules/clients/domain/ip_allowlist"

	"github.com/google/uuid"
)

// CreateIPAllowlistCmd 创建IP白名单命令
type CreateIPAllowlistCmd struct {
	RuleID      string
	AppID       uuid.UUID
	CIDR        string
	Description *string
	Status      ip_allowlist.AllowlistStatus
	CreatedBy   *uuid.UUID
}

// UpdateIPAllowlistCmd 更新IP白名单命令
type UpdateIPAllowlistCmd struct {
	IPAllowlistID uuid.UUID
	CIDR          string
	Description   *string
	UpdatedBy     *uuid.UUID
}

// UpdateIPAllowlistStatusCmd 更新IP白名单状态命令
type UpdateIPAllowlistStatusCmd struct {
	IPAllowlistID uuid.UUID
	Status        ip_allowlist.AllowlistStatus
}

// RevokeIPAllowlistCmd 撤销IP白名单命令
type RevokeIPAllowlistCmd struct {
	RuleID       string
	RevokedBy    uuid.UUID
	RevokeReason *string
}

// DeleteIPAllowlistCmd 删除IP白名单命令
type DeleteIPAllowlistCmd struct {
	RuleID string
}
