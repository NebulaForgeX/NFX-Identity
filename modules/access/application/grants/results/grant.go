package results

import (
	"time"

	"nfxid/modules/access/domain/grants"

	"github.com/google/uuid"
)

type GrantRO struct {
	ID           uuid.UUID
	SubjectType  grants.SubjectType
	SubjectID    uuid.UUID
	GrantType    grants.GrantType
	GrantRefID   uuid.UUID
	TenantID     *uuid.UUID
	AppID        *uuid.UUID
	ResourceType *string
	ResourceID   *uuid.UUID
	Effect       grants.GrantEffect
	ExpiresAt    *time.Time
	CreatedAt    time.Time
	CreatedBy    *uuid.UUID
	RevokedAt    *time.Time
	RevokedBy    *uuid.UUID
	RevokeReason *string
}

// GrantMapper 将 Domain Grant 转换为 Application GrantRO
func GrantMapper(g *grants.Grant) GrantRO {
	if g == nil {
		return GrantRO{}
	}

	return GrantRO{
		ID:           g.ID(),
		SubjectType:  g.SubjectType(),
		SubjectID:    g.SubjectID(),
		GrantType:    g.GrantType(),
		GrantRefID:   g.GrantRefID(),
		TenantID:     g.TenantID(),
		AppID:        g.AppID(),
		ResourceType: g.ResourceType(),
		ResourceID:   g.ResourceID(),
		Effect:       g.Effect(),
		ExpiresAt:    g.ExpiresAt(),
		CreatedAt:    g.CreatedAt(),
		CreatedBy:    g.CreatedBy(),
		RevokedAt:    g.RevokedAt(),
		RevokedBy:    g.RevokedBy(),
		RevokeReason: g.RevokeReason(),
	}
}
