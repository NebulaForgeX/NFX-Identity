package tenant_roles

import (
	"time"

	"github.com/google/uuid"
)

func nowUTC() time.Time { return time.Now().UTC() }

// Validate 校验
func (r *TenantRole) Validate() error {
	if r.RoleKey() == "" {
		return ErrTenantRoleRoleKeyRequired
	}
	if r.TenantID() == uuid.Nil {
		return ErrTenantRoleTenantIDRequired
	}
	return nil
}
