package http

import (
	bootstrapApp "nfxid/modules/system/application/bootstrap"
	systemStateApp "nfxid/modules/system/application/system_state"
	"nfxid/modules/system/interfaces/http/handler"
)

type Registry struct {
	SystemState *handler.SystemStateHandler
}

func NewRegistry(
	systemStateAppSvc *systemStateApp.Service,
	bootstrapSvc *bootstrapApp.Service,
) *Registry {
	return &Registry{
		SystemState: handler.NewSystemStateHandler(systemStateAppSvc, bootstrapSvc),
	}
}
