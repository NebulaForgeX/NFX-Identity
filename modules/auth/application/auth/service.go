package auth

import (
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
	grpcClients "nfxid/modules/auth/infrastructure/grpc"
)

// Service 认证应用服务（登录、刷新 Token）
type Service struct {
	credRepo           *userCredentialDomain.Repo
	loginAttemptRepo   *loginAttemptDomain.Repo
	accountLockoutRepo *accountLockoutDomain.Repo
	refreshTokenRepo   *refreshTokenDomain.Repo
	grpcClients        *grpcClients.GRPCClients // gRPC 客户端（通过依赖注入）
	tokenIssuer        TokenIssuer
	expiresInSec       int64
	refreshTokenTTL    int64 // refresh token 有效期（秒）
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
) *Service {
	return &Service{
		credRepo:           credRepo,
		loginAttemptRepo:   loginAttemptRepo,
		accountLockoutRepo: accountLockoutRepo,
		refreshTokenRepo:   refreshTokenRepo,
		grpcClients:        grpcClients,
		tokenIssuer:        tokenIssuer,
		expiresInSec:       expiresInSec,
		refreshTokenTTL:    refreshTokenTTL,
	}
}
