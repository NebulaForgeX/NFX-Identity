package tenants

import (
	"time"

	"github.com/google/uuid"
)

type NewTenantParams struct {
	TenantID      string
	Name          string
	DisplayName   *string
	Status        TenantStatus
	PrimaryDomain *string
	Metadata      map[string]interface{}
}

func NewTenant(p NewTenantParams) (*Tenant, error) {
	if err := validateTenantParams(p); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = TenantStatusPending
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewTenantFromState(TenantState{
		ID:            id,
		TenantID:      p.TenantID,
		Name:          p.Name,
		DisplayName:   p.DisplayName,
		Status:        status,
		PrimaryDomain: p.PrimaryDomain,
		Metadata:      p.Metadata,
		CreatedAt:     now,
		UpdatedAt:     now,
	}), nil
}

func NewTenantFromState(st TenantState) *Tenant {
	return &Tenant{state: st}
}

func validateTenantParams(p NewTenantParams) error {
	if p.TenantID == "" {
		return ErrTenantIDRequired
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Status != "" {
		validStatuses := map[TenantStatus]struct{}{
			TenantStatusActive:    {},
			TenantStatusSuspended: {},
			TenantStatusClosed:    {},
			TenantStatusPending:   {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidTenantStatus
		}
	}
	return nil
}
