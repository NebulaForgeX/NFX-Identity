package tenant_apps

import "github.com/google/uuid"

func (ta *TenantApp) Validate() error {
	if ta.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if ta.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	validStatuses := map[TenantAppStatus]struct{}{
		TenantAppStatusActive:    {},
		TenantAppStatusDisabled:  {},
		TenantAppStatusSuspended: {},
	}
	if _, ok := validStatuses[ta.Status()]; !ok {
		return ErrInvalidTenantAppStatus
	}
	return nil
}
