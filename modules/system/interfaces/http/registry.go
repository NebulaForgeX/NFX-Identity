package http

import (
	bootstrapApp "nfxid/modules/system/application/bootstrap"
	systemStateApp "nfxid/modules/system/application/system_state"
	"nfxid/modules/system/interfaces/http/handler"
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
