package commands

import (
	"nfxid/modules/tenants/domain/invitations"

	"github.com/google/uuid"
)

// CreateInvitationCmd 创建邀请命令
type CreateInvitationCmd struct {
	InviteID  string
	TenantID  uuid.UUID
	Email     string
	TokenHash string
	ExpiresAt string
	Status    invitations.InvitationStatus
	InvitedBy uuid.UUID
	RoleIDs   []uuid.UUID
	Metadata  map[string]interface{}
}

// AcceptInvitationCmd 接受邀请命令
type AcceptInvitationCmd struct {
	InviteID string
	UserID   uuid.UUID
}

// RevokeInvitationCmd 撤销邀请命令
type RevokeInvitationCmd struct {
	InviteID    string
	RevokedBy   uuid.UUID
	RevokeReason string
}

// DeleteInvitationCmd 删除邀请命令
type DeleteInvitationCmd struct {
	InviteID string
}
