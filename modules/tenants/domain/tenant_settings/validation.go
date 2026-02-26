package tenant_settings

import (
	tenantsErr "nfxidentity/errors/src/tenants"

	"github.com/google/uuid"
)

func (ts *TenantSetting) Validate() error {
	if ts.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	return nil
}
