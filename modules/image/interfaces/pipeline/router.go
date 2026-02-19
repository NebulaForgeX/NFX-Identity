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
	// 注意：使用 eventbus.RegisterHandler，但文件夹名称是 pipeline
	eventbus.RegisterHandler(r.EventRouter, r.registry.ImageHandler.OnImagesInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.ImageHandler.OnImageTypesInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.ImageHandler.OnImageVariantsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.ImageHandler.OnImageTagsInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.ImageHandler.OnUserAvatarReplaced)
}

func (r *Router) Run(ctx context.Context) error {
	logx.S().Info("Starting pipeline router...")
	return r.Router.Run(ctx)
}

func (r *Router) Close() error {
	return r.Router.Close()
}
