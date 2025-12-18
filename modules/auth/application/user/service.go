package user

import (
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/pkgs/cache"
	"nfxid/pkgs/email"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/tokenx"
)

type Service struct {
	userRepo     *userDomain.Repo
	userQuery    *userDomain.Query
	busPublisher *eventbus.BusPublisher
	tokenx       *tokenx.Tokenx
	cache        cache.BaseCache
	emailService *email.EmailService
}

func NewService(
	userRepo *userDomain.Repo,
	userQuery *userDomain.Query,
	busPublisher *eventbus.BusPublisher,
	tokenx *tokenx.Tokenx,
	cacheConn *cache.Connection,
	emailService *email.EmailService,
) *Service {
	// 创建 BaseCache，使用默认 JSON codec（对于字符串验证码，会直接使用 Redis 原生类型）
	baseCache := cache.NewBaseCache(cacheConn.Client(), nil)
	return &Service{
		userRepo:     userRepo,
		userQuery:    userQuery,
		busPublisher: busPublisher,
		tokenx:       tokenx,
		cache:        baseCache,
		emailService: emailService,
	}
}
