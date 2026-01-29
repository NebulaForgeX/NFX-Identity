package pipeline

import (
	imageApp "nfxid/modules/image/application/images"
	"nfxid/modules/image/interfaces/pipeline/handler"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/kafkax/eventbus"
	"nfxid/pkgs/logx"
	"time"

	wmMiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type Deps interface {
	KafkaConfig() *kafkax.Config
	BusPublisher() *eventbus.BusPublisher
	ImageAppSvc() *imageApp.Service
}

func NewServer(d Deps) (*Router, error) {
	// 创建订阅者
	sub, err := kafkax.NewSubscriber(d.KafkaConfig())
	if err != nil {
		return nil, err
	}

	registry := &Registry{
		ImageHandler: handler.NewImageHandler(d.ImageAppSvc()),
	}

	router, err := NewRouter(sub, registry, eventbus.EventRouterConfig{
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

	router.RegisterRoutes()

	return router, nil
}
