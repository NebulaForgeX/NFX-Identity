package results

import (
	"time"

	"nfxid/modules/access/domain/permissions"

	"github.com/google/uuid"
)

type PermissionRO struct {
	ID          uuid.UUID
	Key         string
	Name        string
	Description *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// PermissionMapper 将 Domain Permission 转换为 Application PermissionRO
func PermissionMapper(p *permissions.Permission) PermissionRO {
	if p == nil {
		return PermissionRO{}
	}

	return PermissionRO{
		ID:          p.ID(),
		Key:         p.Key(),
		Name:        p.Name(),
		Description: p.Description(),
		IsSystem:    p.IsSystem(),
		CreatedAt:   p.CreatedAt(),
		UpdatedAt:   p.UpdatedAt(),
		DeletedAt:   p.DeletedAt(),
	}
}
