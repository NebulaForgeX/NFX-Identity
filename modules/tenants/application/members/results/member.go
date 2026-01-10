package results

import (
	"time"

	"nfxid/modules/tenants/domain/members"

	"github.com/google/uuid"
)

type MemberRO struct {
	ID          uuid.UUID
	MemberID    uuid.UUID
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Status      members.MemberStatus
	Source      members.MemberSource
	JoinedAt    *time.Time
	LeftAt      *time.Time
	CreatedAt   time.Time
	CreatedBy   *uuid.UUID
	UpdatedAt   time.Time
	ExternalRef *string
	Metadata    map[string]interface{}
}

// MemberMapper 将 Domain Member 转换为 Application MemberRO
func MemberMapper(m *members.Member) MemberRO {
	if m == nil {
		return MemberRO{}
	}

	return MemberRO{
		ID:          m.ID(),
		MemberID:    m.MemberID(),
		TenantID:    m.TenantID(),
		UserID:      m.UserID(),
		Status:      m.Status(),
		Source:      m.Source(),
		JoinedAt:    m.JoinedAt(),
		LeftAt:      m.LeftAt(),
		CreatedAt:   m.CreatedAt(),
		CreatedBy:   m.CreatedBy(),
		UpdatedAt:   m.UpdatedAt(),
		ExternalRef: m.ExternalRef(),
		Metadata:    m.Metadata(),
	}
}
