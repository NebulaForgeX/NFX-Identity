package apps

import "github.com/google/uuid"

func (a *App) Validate() error {
	if a.AppID() == "" {
		return ErrAppIDRequired
	}
	if a.Name() == "" {
		return ErrNameRequired
	}
	if a.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	validTypes := map[AppType]struct{}{
		AppTypeServer:     {},
		AppTypeService:    {},
		AppTypeInternal:   {},
		AppTypePartner:    {},
		AppTypeThirdParty: {},
	}
	if _, ok := validTypes[a.Type()]; !ok {
		return ErrInvalidAppType
	}
	validStatuses := map[AppStatus]struct{}{
		AppStatusActive:    {},
		AppStatusDisabled:  {},
		AppStatusSuspended: {},
		AppStatusPending:   {},
	}
	if _, ok := validStatuses[a.Status()]; !ok {
		return ErrInvalidAppStatus
	}
	validEnvironments := map[Environment]struct{}{
		EnvironmentProduction:  {},
		EnvironmentStaging:     {},
		EnvironmentDevelopment: {},
		EnvironmentTest:        {},
	}
	if _, ok := validEnvironments[a.Environment()]; !ok {
		return ErrInvalidEnvironment
	}
	return nil
}
