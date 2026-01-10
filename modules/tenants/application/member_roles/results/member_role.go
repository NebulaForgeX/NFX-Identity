package results

import (
	"time"

	"nfxid/modules/tenants/domain/member_roles"

	"github.com/google/uuid"
)

type MemberRoleRO struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	MemberID    uuid.UUID
	RoleID      uuid.UUID
	AssignedAt  time.Time
	AssignedBy  *uuid.UUID
	ExpiresAt   *time.Time
	Scope       *string
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
}

// MemberRoleMapper 将 Domain MemberRole 转换为 Application MemberRoleRO
func MemberRoleMapper(mr *member_roles.MemberRole) MemberRoleRO {
	if mr == nil {
		return MemberRoleRO{}
	}

	return MemberRoleRO{
		ID:          mr.ID(),
		TenantID:    mr.TenantID(),
		MemberID:    mr.MemberID(),
		RoleID:      mr.RoleID(),
		AssignedAt:  mr.AssignedAt(),
		AssignedBy:  mr.AssignedBy(),
		ExpiresAt:   mr.ExpiresAt(),
		Scope:       mr.Scope(),
		RevokedAt:   mr.RevokedAt(),
		RevokedBy:   mr.RevokedBy(),
		RevokeReason: mr.RevokeReason(),
	}
}
