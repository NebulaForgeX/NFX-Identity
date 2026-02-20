package reqdto

import (
	"github.com/google/uuid"
)

type TenantRoleCreateRequestDTO struct {
	TenantID string  `json:"tenant_id" validate:"required,uuid"`
	RoleKey  string  `json:"role_key" validate:"required"`
	Name     *string `json:"name,omitempty"`
}

type TenantRoleByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type TenantRoleByTenantIDRequestDTO struct {
	TenantID uuid.UUID `uri:"tenant_id" validate:"required,uuid"`
}

type TenantRoleByTenantIDAndRoleKeyRequestDTO struct {
	TenantID uuid.UUID `uri:"tenant_id" validate:"required,uuid"`
	RoleKey  string    `uri:"role_key" validate:"required"`
}

type TenantRoleUpdateBodyRequestDTO struct {
	RoleKey string  `json:"role_key" validate:"required"`
	Name    *string `json:"name,omitempty"`
}

func (r *TenantRoleCreateRequestDTO) ToCreateParams() (tenantID uuid.UUID, roleKey string, name *string, err error) {
	tenantID, err = uuid.Parse(r.TenantID)
	if err != nil {
		return uuid.Nil, "", nil, err
	}
	return tenantID, r.RoleKey, r.Name, nil
}
