package pipeline

import (
	"context"

	"nfxid/pkgs/kafkax/eventbus"
	"nfxid/pkgs/logx"
)

type Router struct {
	*eventbus.EventRouter
	registry *Registry
}

func NewRouter(sub *eventbus.BusSubscriber, registry *Registry, config eventbus.EventRouterConfig) (*Router, error) {
	router, err := eventbus.NewEventRouter(sub, config)
	if err != nil {
		return nil, err
	}
	return &Router{EventRouter: router, registry: registry}, nil
}

func (r *Router) RegisterRoutes() {
	// 注册事件处理器
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAccountLockoutsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnLoginAttemptsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnMFAFactorsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnPasswordHistoryInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnPasswordResetsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnRefreshTokensInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnSessionsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnTrustedDevicesInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnUserCredentialsInvalidateCache)
}

func (r *Router) Run(ctx context.Context) error {
	logx.S().Info("Starting pipeline router...")
	return r.Router.Run(ctx)
}

func (r *Router) Close() error {
	return r.Router.Close()
}
