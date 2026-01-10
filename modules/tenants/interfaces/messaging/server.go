package messaging

import (
	"context"

	"nfxid/pkgs/logx"
	"nfxid/pkgs/rabbitmqx"
	"nfxid/pkgs/rabbitmqx/messaging"
	"time"

	wmMiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type Deps interface {
	RabbitMQConfig() *rabbitmqx.Config
}

type Router struct {
	*messaging.MessageRouter
}

func NewServer(d Deps) (*Router, error) {
	// 创建订阅者
	sub, err := rabbitmqx.NewSubscriber(d.RabbitMQConfig())
	if err != nil {
		return nil, err
	}

	router, err := messaging.NewMessageRouter(sub, messaging.MessageRouterConfig{
		CloseTimeout: 10 * time.Second,
		Logger:       logx.NewZapWatermillLogger(logx.L()),
	})
	if err != nil {
		return nil, err
	}

	// 添加中间件
	router.Router.AddMiddleware(
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

	return &Router{MessageRouter: router}, nil
}

func (r *Router) RegisterRoutes() {
	// 注册消息处理器
	// 注意：这里需要根据实际的消息类型注册处理器
	// 目前先留空，等待具体的消息处理器实现
}

func (r *Router) Run(ctx context.Context) error {
	logx.S().Info("Starting messaging router...")
	return r.Router.Run(ctx)
}

func (r *Router) Close() error {
	return r.Router.Close()
}
