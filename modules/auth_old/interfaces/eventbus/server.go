package eventbus

import (
	"context"
	badgeApp "nfxid/modules/auth/application/badge"
	profileApp "nfxid/modules/auth/application/profile"
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	educationApp "nfxid/modules/auth/application/profile_education"
	occupationApp "nfxid/modules/auth/application/profile_occupation"
	roleApp "nfxid/modules/auth/application/role"
	userApp "nfxid/modules/auth/application/user"
	"nfxid/modules/auth/interfaces/eventbus/handler"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"time"

	wmMiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type Deps interface {
	KafkaConfig() *kafkax.Config
	BusPublisher() *eventbus.BusPublisher
	UserAppSvc() *userApp.Service
	ProfileAppSvc() *profileApp.Service
	RoleAppSvc() *roleApp.Service
	BadgeAppSvc() *badgeApp.Service
	EducationAppSvc() *educationApp.Service
	OccupationAppSvc() *occupationApp.Service
	ProfileBadgeAppSvc() *profileBadgeApp.Service
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
	registry := &Registry{
		Auth: handler.NewAuthHandler(
			d.UserAppSvc(),
			d.ProfileAppSvc(),
			d.RoleAppSvc(),
			d.BadgeAppSvc(),
			d.EducationAppSvc(),
			d.OccupationAppSvc(),
			d.ProfileBadgeAppSvc(),
		),
		Image: handler.NewImageHandler(),
	}

	return &Router{
		EventRouter: router,
		registry:    registry,
	}, nil
}

func (r *Router) RegisterRoutes() {
	// 注册 Auth 内部事件处理器（Auth -> Auth）
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_Success)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_Test)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_UserCreated)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_UserUpdated)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_UserDeleted)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_UserInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_ProfileInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_RoleInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_BadgeInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_EducationInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_OccupationInvalidateCache)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Auth.OnAuthToAuth_ProfileBadgeInvalidateCache)

	// 注册 Image -> Auth 事件处理器（Image 服务通知 Auth 服务）
	eventbus.RegisterHandler(r.EventRouter, r.registry.Image.OnImageToAuth_ImageDelete)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Image.OnImageToAuth_ImageSuccess)
	eventbus.RegisterHandler(r.EventRouter, r.registry.Image.OnImageToAuth_ImageTest)
}

func (r *Router) Run(ctx context.Context) error {
	logx.S().Info("Starting eventbus router...")
	return r.Router.Run(ctx)
}

func (r *Router) Close() error {
	return r.Router.Close()
}
