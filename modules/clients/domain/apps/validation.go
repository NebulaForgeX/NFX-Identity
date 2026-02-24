package apps

import (
	clientsErr "nfxid/errors/src/clients"

	"github.com/google/uuid"
)

func (a *App) Validate() error {
	if a.AppID() == "" {
		return clientsErr.ErrAppIDRequired
	}
	if a.Name() == "" {
		return clientsErr.ErrNameRequired
	}
	if a.TenantID() == uuid.Nil {
		return clientsErr.ErrTenantIDRequired
	}
	validTypes := map[AppType]struct{}{
		AppTypeServer:     {},
		AppTypeService:    {},
		AppTypeInternal:   {},
		AppTypePartner:    {},
		AppTypeThirdParty: {},
	}
	if _, ok := validTypes[a.Type()]; !ok {
		return clientsErr.ErrInvalidAppType
	}
	validStatuses := map[AppStatus]struct{}{
		AppStatusActive:    {},
		AppStatusDisabled:  {},
		AppStatusSuspended: {},
		AppStatusPending:   {},
	}
	if _, ok := validStatuses[a.Status()]; !ok {
		return clientsErr.ErrInvalidAppStatus
	}
	validEnvironments := map[Environment]struct{}{
		EnvironmentProduction:  {},
		EnvironmentStaging:     {},
		EnvironmentDevelopment: {},
		EnvironmentTest:        {},
	}
	if _, ok := validEnvironments[a.Environment()]; !ok {
		return clientsErr.ErrInvalidEnvironment
	}
	return nil
}
