package results

import (
	"time"

	"nfxid/modules/tenants/domain/groups"

	"github.com/google/uuid"
)

type GroupRO struct {
	ID            uuid.UUID
	GroupID       string
	TenantID      uuid.UUID
	Name          string
	Type          groups.GroupType
	ParentGroupID *uuid.UUID
	Description   *string
	CreatedBy     *uuid.UUID
	Metadata      map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// GroupMapper 将 Domain Group 转换为 Application GroupRO
func GroupMapper(g *groups.Group) GroupRO {
	if g == nil {
		return GroupRO{}
	}

	return GroupRO{
		ID:            g.ID(),
		GroupID:       g.GroupID(),
		TenantID:      g.TenantID(),
		Name:          g.Name(),
		Type:          g.Type(),
		ParentGroupID: g.ParentGroupID(),
		Description:   g.Description(),
		CreatedBy:     g.CreatedBy(),
		Metadata:      g.Metadata(),
		CreatedAt:     g.CreatedAt(),
		UpdatedAt:     g.UpdatedAt(),
		DeletedAt:     g.DeletedAt(),
	}
}
