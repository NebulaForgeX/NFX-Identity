package results

import (
	"time"

	"nfxid/modules/tenants/domain/member_app_roles"

	"github.com/google/uuid"
)

type MemberAppRoleRO struct {
	ID           uuid.UUID
	MemberID     uuid.UUID
	AppID        uuid.UUID
	RoleID       uuid.UUID
	AssignedAt   time.Time
	AssignedBy   *uuid.UUID
	ExpiresAt    *time.Time
	RevokedAt    *time.Time
	RevokedBy    *uuid.UUID
	RevokeReason *string
}

// MemberAppRoleMapper 将 Domain MemberAppRole 转换为 Application MemberAppRoleRO
func MemberAppRoleMapper(mar *member_app_roles.MemberAppRole) MemberAppRoleRO {
	if mar == nil {
		return MemberAppRoleRO{}
	}

	return MemberAppRoleRO{
		ID:           mar.ID(),
		MemberID:     mar.MemberID(),
		AppID:        mar.AppID(),
		RoleID:       mar.RoleID(),
		AssignedAt:   mar.AssignedAt(),
		AssignedBy:   mar.AssignedBy(),
		ExpiresAt:    mar.ExpiresAt(),
		RevokedAt:    mar.RevokedAt(),
		RevokedBy:    mar.RevokedBy(),
		RevokeReason: mar.RevokeReason(),
	}
}
