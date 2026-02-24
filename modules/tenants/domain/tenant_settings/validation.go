package tenant_settings

import (
	tenantsErr "nfxid/errors/src/tenants"

	"github.com/google/uuid"
)

func (ts *TenantSetting) Validate() error {
	if ts.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	return nil
}
