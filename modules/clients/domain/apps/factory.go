package apps

import (
	"time"

	"github.com/google/uuid"
)

type NewAppParams struct {
	AppID       string
	TenantID    uuid.UUID
	Name        string
	Description *string
	Type        AppType
	Status      AppStatus
	Environment Environment
	CreatedBy   *uuid.UUID
	Metadata    map[string]interface{}
}

func NewApp(p NewAppParams) (*App, error) {
	if err := validateAppParams(p); err != nil {
		return nil, err
	}

	appType := p.Type
	if appType == "" {
		appType = AppTypeServer
	}

	status := p.Status
	if status == "" {
		status = AppStatusPending
	}

	environment := p.Environment
	if environment == "" {
		environment = EnvironmentDevelopment
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewAppFromState(AppState{
		ID:          id,
		AppID:       p.AppID,
		TenantID:    p.TenantID,
		Name:        p.Name,
		Description: p.Description,
		Type:        appType,
		Status:      status,
		Environment: environment,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   p.CreatedBy,
		Metadata:    p.Metadata,
	}), nil
}

func NewAppFromState(st AppState) *App {
	return &App{state: st}
}

func validateAppParams(p NewAppParams) error {
	if p.AppID == "" {
		return ErrAppIDRequired
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.Type != "" {
		validTypes := map[AppType]struct{}{
			AppTypeServer:     {},
			AppTypeService:    {},
			AppTypeInternal:   {},
			AppTypePartner:    {},
			AppTypeThirdParty: {},
		}
		if _, ok := validTypes[p.Type]; !ok {
			return ErrInvalidAppType
		}
	}
	if p.Status != "" {
		validStatuses := map[AppStatus]struct{}{
			AppStatusActive:    {},
			AppStatusDisabled:  {},
			AppStatusSuspended: {},
			AppStatusPending:   {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidAppStatus
		}
	}
	if p.Environment != "" {
		validEnvironments := map[Environment]struct{}{
			EnvironmentProduction:  {},
			EnvironmentStaging:     {},
			EnvironmentDevelopment: {},
			EnvironmentTest:        {},
		}
		if _, ok := validEnvironments[p.Environment]; !ok {
			return ErrInvalidEnvironment
		}
	}
	return nil
}
