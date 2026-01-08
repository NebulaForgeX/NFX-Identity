package tenants

import (
	"time"
)

func (t *Tenant) Update(name string, displayName, primaryDomain *string, metadata map[string]interface{}) error {
	if t.DeletedAt() != nil {
		return ErrTenantNotFound
	}
	if name == "" {
		return ErrNameRequired
	}

	t.state.Name = name
	if displayName != nil {
		t.state.DisplayName = displayName
	}
	if primaryDomain != nil {
		t.state.PrimaryDomain = primaryDomain
	}
	if metadata != nil {
		t.state.Metadata = metadata
	}

	t.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (t *Tenant) UpdateStatus(status TenantStatus) error {
	if t.DeletedAt() != nil {
		return ErrTenantNotFound
	}
	validStatuses := map[TenantStatus]struct{}{
		TenantStatusActive:    {},
		TenantStatusSuspended: {},
		TenantStatusClosed:    {},
		TenantStatusPending:   {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidTenantStatus
	}

	t.state.Status = status
	t.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (t *Tenant) Delete() error {
	if t.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	t.state.DeletedAt = &now
	t.state.UpdatedAt = now
	return nil
}
