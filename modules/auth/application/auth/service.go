package auth

import (
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
)

// Service 认证应用服务（登录、刷新 Token）
type Service struct {
	credRepo          *userCredentialDomain.Repo
	loginAttemptRepo  *loginAttemptDomain.Repo
	accountLockoutRepo *accountLockoutDomain.Repo
	userResolver      UserResolver
	tokenIssuer       TokenIssuer
	expiresInSec      int64
}

// NewService 创建认证应用服务；注入 domain 仓库与 infra 实现的端口。expiresInSec 为 access token 有效期（秒），如 900。
func NewService(
	credRepo *userCredentialDomain.Repo,
	loginAttemptRepo *loginAttemptDomain.Repo,
	accountLockoutRepo *accountLockoutDomain.Repo,
	userResolver UserResolver,
	tokenIssuer TokenIssuer,
	expiresInSec int64,
) *Service {
	if expiresInSec <= 0 {
		expiresInSec = 900
	}
	return &Service{
		credRepo:          credRepo,
		loginAttemptRepo:  loginAttemptRepo,
		accountLockoutRepo: accountLockoutRepo,
		userResolver:      userResolver,
		tokenIssuer:       tokenIssuer,
		expiresInSec:      expiresInSec,
	}
}
