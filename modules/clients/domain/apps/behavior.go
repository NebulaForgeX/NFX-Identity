package apps

import (
	"time"

	"github.com/google/uuid"
)

func (a *App) Update(name string, description *string, appType AppType, environment Environment, metadata map[string]interface{}, updatedBy *uuid.UUID) error {
	if a.DeletedAt() != nil {
		return ErrAppNotFound
	}
	if name == "" {
		return ErrNameRequired
	}
	if appType != "" {
		validTypes := map[AppType]struct{}{
			AppTypeServer:     {},
			AppTypeService:    {},
			AppTypeInternal:   {},
			AppTypePartner:    {},
			AppTypeThirdParty: {},
		}
		if _, ok := validTypes[appType]; !ok {
			return ErrInvalidAppType
		}
		a.state.Type = appType
	}
	if environment != "" {
		validEnvironments := map[Environment]struct{}{
			EnvironmentProduction:  {},
			EnvironmentStaging:     {},
			EnvironmentDevelopment: {},
			EnvironmentTest:        {},
		}
		if _, ok := validEnvironments[environment]; !ok {
			return ErrInvalidEnvironment
		}
		a.state.Environment = environment
	}

	a.state.Name = name
	a.state.Description = description
	if metadata != nil {
		a.state.Metadata = metadata
	}
	a.state.UpdatedBy = updatedBy
	a.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (a *App) UpdateStatus(status AppStatus, updatedBy *uuid.UUID) error {
	if a.DeletedAt() != nil {
		return ErrAppNotFound
	}
	validStatuses := map[AppStatus]struct{}{
		AppStatusActive:    {},
		AppStatusDisabled:  {},
		AppStatusSuspended: {},
		AppStatusPending:   {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidAppStatus
	}

	a.state.Status = status
	a.state.UpdatedBy = updatedBy
	a.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (a *App) Delete() error {
	if a.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	a.state.DeletedAt = &now
	a.state.UpdatedAt = now
	return nil
}
