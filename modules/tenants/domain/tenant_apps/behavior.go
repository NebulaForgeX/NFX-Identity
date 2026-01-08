package tenant_apps

import (
	"time"
)

func (ta *TenantApp) UpdateStatus(status TenantAppStatus) error {
	validStatuses := map[TenantAppStatus]struct{}{
		TenantAppStatusActive:    {},
		TenantAppStatusDisabled:  {},
		TenantAppStatusSuspended: {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidTenantAppStatus
	}

	ta.state.Status = status
	ta.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ta *TenantApp) UpdateSettings(settings map[string]interface{}) error {
	if settings == nil {
		return nil
	}
	ta.state.Settings = settings
	ta.state.UpdatedAt = time.Now().UTC()
	return nil
}
