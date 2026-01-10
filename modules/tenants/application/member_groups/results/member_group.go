package results

import (
	"time"

	"nfxid/modules/tenants/domain/member_groups"

	"github.com/google/uuid"
)

type MemberGroupRO struct {
	ID         uuid.UUID
	MemberID   uuid.UUID
	GroupID    uuid.UUID
	AssignedAt time.Time
	AssignedBy *uuid.UUID
	RevokedAt  *time.Time
	RevokedBy  *uuid.UUID
}

// MemberGroupMapper 将 Domain MemberGroup 转换为 Application MemberGroupRO
func MemberGroupMapper(mg *member_groups.MemberGroup) MemberGroupRO {
	if mg == nil {
		return MemberGroupRO{}
	}

	return MemberGroupRO{
		ID:         mg.ID(),
		MemberID:   mg.MemberID(),
		GroupID:    mg.GroupID(),
		AssignedAt: mg.AssignedAt(),
		AssignedBy: mg.AssignedBy(),
		RevokedAt:  mg.RevokedAt(),
		RevokedBy:  mg.RevokedBy(),
	}
}
