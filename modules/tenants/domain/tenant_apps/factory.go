package tenant_apps

import (
	"time"

	"github.com/google/uuid"
)

type NewTenantAppParams struct {
	TenantID  uuid.UUID
	AppID     uuid.UUID
	Status    TenantAppStatus
	CreatedBy *uuid.UUID
	Settings  map[string]interface{}
}

func NewTenantApp(p NewTenantAppParams) (*TenantApp, error) {
	if err := validateTenantAppParams(p); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = TenantAppStatusActive
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewTenantAppFromState(TenantAppState{
		ID:        id,
		TenantID:  p.TenantID,
		AppID:     p.AppID,
		Status:    status,
		CreatedAt: now,
		CreatedBy: p.CreatedBy,
		UpdatedAt: now,
		Settings:  p.Settings,
	}), nil
}

func NewTenantAppFromState(st TenantAppState) *TenantApp {
	return &TenantApp{state: st}
}

func validateTenantAppParams(p NewTenantAppParams) error {
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.AppID == uuid.Nil {
		return ErrAppIDRequired
	}
	if p.Status != "" {
		validStatuses := map[TenantAppStatus]struct{}{
			TenantAppStatusActive:    {},
			TenantAppStatusDisabled:  {},
			TenantAppStatusSuspended: {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidTenantAppStatus
		}
	}
	return nil
}
