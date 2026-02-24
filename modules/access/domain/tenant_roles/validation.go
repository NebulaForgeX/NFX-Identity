package tenant_roles

import (
	"time"

	accessErr "nfxid/errors/src/access"

	"github.com/google/uuid"
)

func nowUTC() time.Time { return time.Now().UTC() }

// Validate 校验
func (r *TenantRole) Validate() error {
	if r.RoleKey() == "" {
		return accessErr.ErrTenantRoleRoleKeyRequired
	}
	if r.TenantID() == uuid.Nil {
		return accessErr.ErrTenantRoleTenantIDRequired
	}
	return nil
}
