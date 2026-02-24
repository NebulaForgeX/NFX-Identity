package tenants

import (
	tenantsErr "nfxid/errors/src/tenants"
)

func (t *Tenant) Validate() error {
	if t.TenantID() == "" {
		return tenantsErr.ErrTenantIDRequired
	}
	if t.Name() == "" {
		return tenantsErr.ErrNameRequired
	}
	validStatuses := map[TenantStatus]struct{}{
		TenantStatusActive:    {},
		TenantStatusSuspended: {},
		TenantStatusClosed:    {},
		TenantStatusPending:   {},
	}
	if _, ok := validStatuses[t.Status()]; !ok {
		return tenantsErr.ErrInvalidTenantStatus
	}
	return nil
}
