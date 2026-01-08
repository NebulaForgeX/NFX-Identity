package tenants

func (t *Tenant) Validate() error {
	if t.TenantID() == "" {
		return ErrTenantIDRequired
	}
	if t.Name() == "" {
		return ErrNameRequired
	}
	validStatuses := map[TenantStatus]struct{}{
		TenantStatusActive:    {},
		TenantStatusSuspended: {},
		TenantStatusClosed:    {},
		TenantStatusPending:   {},
	}
	if _, ok := validStatuses[t.Status()]; !ok {
		return ErrInvalidTenantStatus
	}
	return nil
}
