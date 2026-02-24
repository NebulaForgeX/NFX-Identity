package tenant_apps

import (
	tenantsErr "nfxid/errors/src/tenants"

	"github.com/google/uuid"
)

func (ta *TenantApp) Validate() error {
	if ta.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if ta.AppID() == uuid.Nil {
		return tenantsErr.ErrAppIDRequired
	}
	validStatuses := map[TenantAppStatus]struct{}{
		TenantAppStatusActive:    {},
		TenantAppStatusDisabled:  {},
		TenantAppStatusSuspended: {},
	}
	if _, ok := validStatuses[ta.Status()]; !ok {
		return tenantsErr.ErrInvalidTenantAppStatus
	}
	return nil
}
