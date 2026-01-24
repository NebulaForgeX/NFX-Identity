package auth

import (
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
	grpcClients "nfxid/modules/auth/infrastructure/grpc"
	"nfxid/pkgs/cache"
	emailPkg "nfxid/pkgs/email"
)

// Service 认证应用服务（登录、刷新 Token、注册）
type Service struct {
	credRepo             *userCredentialDomain.Repo
	loginAttemptRepo     *loginAttemptDomain.Repo
	accountLockoutRepo   *accountLockoutDomain.Repo
	refreshTokenRepo     *refreshTokenDomain.Repo
	grpcClients          *grpcClients.GRPCClients // gRPC 客户端（通过依赖注入）
	tokenIssuer          TokenIssuer
	expiresInSec         int64
	refreshTokenTTL      int64 // refresh token 有效期（秒）
	emailService         *emailPkg.EmailService
	cache                *cache.Connection
	userCredentialAppSvc *userCredentialApp.Service
}

// NewService 创建认证应用服务；注入 domain 仓库与 infra 实现的端口。
// expiresInSec、refreshTokenTTL 由 wiring 从配置解析，若配置缺失则使用 constants 默认值，此处不再做默认值兜底。
func NewService(
	credRepo *userCredentialDomain.Repo,
	loginAttemptRepo *loginAttemptDomain.Repo,
	accountLockoutRepo *accountLockoutDomain.Repo,
	refreshTokenRepo *refreshTokenDomain.Repo,
	grpcClients *grpcClients.GRPCClients,
	tokenIssuer TokenIssuer,
	expiresInSec int64,
	refreshTokenTTL int64,
	emailService *emailPkg.EmailService,
	cache *cache.Connection,
	userCredentialAppSvc *userCredentialApp.Service,
) *Service {
	return &Service{
		credRepo:             credRepo,
		loginAttemptRepo:     loginAttemptRepo,
		accountLockoutRepo:   accountLockoutRepo,
		refreshTokenRepo:     refreshTokenRepo,
		grpcClients:          grpcClients,
		tokenIssuer:          tokenIssuer,
		expiresInSec:         expiresInSec,
		refreshTokenTTL:      refreshTokenTTL,
		emailService:         emailService,
		cache:                cache,
		userCredentialAppSvc: userCredentialAppSvc,
	}
}
