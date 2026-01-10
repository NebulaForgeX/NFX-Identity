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
	eventbus.RegisterHandler(r.EventRouter, r.registry.AuditHandler.OnActorSnapshotsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.AuditHandler.OnEventRetentionPoliciesInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.AuditHandler.OnEventSearchIndexInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.AuditHandler.OnEventsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.AuditHandler.OnHashChainCheckpointsInvalidateCache)
}

func (r *Router) Run(ctx context.Context) error {
	logx.S().Info("Starting pipeline router...")
	return r.Router.Run(ctx)
}

func (r *Router) Close() error {
	return r.Router.Close()
}
