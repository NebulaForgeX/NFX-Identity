package respdto

import (
	"time"

	domain "nfxid/modules/access/domain/tenant_roles"

	"github.com/google/uuid"
)

type TenantRoleDTO struct {
	ID        uuid.UUID  `json:"id"`
	TenantID  uuid.UUID  `json:"tenant_id"`
	RoleKey   string     `json:"role_key"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

func TenantRoleToDTO(r *domain.TenantRole) *TenantRoleDTO {
	if r == nil {
		return nil
	}
	return &TenantRoleDTO{
		ID:        r.ID(),
		TenantID:  r.TenantID(),
		RoleKey:   r.RoleKey(),
		Name:      r.Name(),
		CreatedAt: r.CreatedAt(),
	}
}

func TenantRoleListToDTO(list []*domain.TenantRole) []TenantRoleDTO {
	dtos := make([]TenantRoleDTO, len(list))
	for i, r := range list {
		if d := TenantRoleToDTO(r); d != nil {
			dtos[i] = *d
		}
	}
	return dtos
}
