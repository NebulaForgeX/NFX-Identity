package eventbus

import (
	"context"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"time"

	wmMiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type Deps interface {
	KafkaConfig() *kafkax.Config
	BusPublisher() *eventbus.BusPublisher
}

type Router struct {
	*eventbus.EventRouter
	registry *Registry
}

func NewServer(d Deps) (*Router, error) {
	// 创建订阅者
	sub, err := kafkax.NewSubscriber(d.KafkaConfig())
	if err != nil {
		return nil, err
	}

	router, err := eventbus.NewEventRouter(sub, eventbus.EventRouterConfig{
		CloseTimeout: 10 * time.Second,
		Logger:       logx.NewZapWatermillLogger(logx.L()),
	})
	if err != nil {
		return nil, err
	}

	// 添加中间件
	router.AddMiddleware(
		wmMiddleware.CorrelationID,
		wmMiddleware.Recoverer,
		wmMiddleware.Retry{
			MaxRetries:      3,
			InitialInterval: 200 * time.Millisecond,
			MaxInterval:     2 * time.Second,
			Multiplier:      2.0,
		}.Middleware,
		wmMiddleware.Timeout(10*time.Second),
	)

	// 创建 registry
	registry := &Registry{}

	return &Router{
		EventRouter: router,
		registry:    registry,
	}, nil
}

func (r *Router) RegisterRoutes() {
	// 注册事件处理器（如果需要的话）
	// eventbus.RegisterHandler(r.EventRouter, r.registry.XXX.OnXXX)
}

func (r *Router) Run(ctx context.Context) error {
	logx.S().Info("Starting eventbus router...")
	return r.Router.Run(ctx)
}

func (r *Router) Close() error {
	return r.Router.Close()
}

