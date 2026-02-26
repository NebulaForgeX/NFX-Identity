package http

import (
	bootstrapApp "nfxidentity/modules/system/application/bootstrap"
	systemStateApp "nfxidentity/modules/system/application/system_state"
	"nfxidentity/modules/system/interfaces/http/handler"
)

type Registry struct {
	SystemState *handler.SystemStateHandler
	I18n        *handler.I18nHandler
}

func NewRegistry(
	systemStateAppSvc *systemStateApp.Service,
	bootstrapSvc *bootstrapApp.Service,
	errorsLangsPath string,
) *Registry {
	return &Registry{
		SystemState: handler.NewSystemStateHandler(systemStateAppSvc, bootstrapSvc),
		I18n:        handler.NewI18nHandler(errorsLangsPath),
	}
}
