package tenant_roles

import (
	"github.com/google/uuid"
)

// NewTenantRoleParams 创建参数
type NewTenantRoleParams struct {
	TenantID uuid.UUID
	RoleKey  string
	Name     *string
}

// NewTenantRole 创建租户角色
func NewTenantRole(p NewTenantRoleParams) (*TenantRole, error) {
	if err := validateNewParams(p); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return NewTenantRoleFromState(TenantRoleState{
		ID:        id,
		TenantID:  p.TenantID,
		RoleKey:   p.RoleKey,
		Name:      p.Name,
		CreatedAt: nowUTC(),
	}), nil
}

// NewTenantRoleFromState 从状态还原
func NewTenantRoleFromState(st TenantRoleState) *TenantRole {
	return &TenantRole{state: st}
}

func validateNewParams(p NewTenantRoleParams) error {
	if p.TenantID == uuid.Nil {
		return ErrTenantRoleTenantIDRequired
	}
	if p.RoleKey == "" {
		return ErrTenantRoleRoleKeyRequired
	}
	return nil
}
