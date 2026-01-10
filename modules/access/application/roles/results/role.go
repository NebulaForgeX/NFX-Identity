package results

import (
	"time"

	"nfxid/modules/access/domain/roles"

	"github.com/google/uuid"
)

type RoleRO struct {
	ID          uuid.UUID
	Key         string
	Name        string
	Description *string
	ScopeType   roles.ScopeType
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// RoleMapper 将 Domain Role 转换为 Application RoleRO
func RoleMapper(r *roles.Role) RoleRO {
	if r == nil {
		return RoleRO{}
	}

	return RoleRO{
		ID:          r.ID(),
		Key:         r.Key(),
		Name:        r.Name(),
		Description: r.Description(),
		ScopeType:   r.ScopeType(),
		IsSystem:    r.IsSystem(),
		CreatedAt:   r.CreatedAt(),
		UpdatedAt:   r.UpdatedAt(),
		DeletedAt:   r.DeletedAt(),
	}
}
