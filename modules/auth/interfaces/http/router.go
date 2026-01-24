package http

import (
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/usertoken"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app           fiber.Router
	tokenVerifier token.Verifier
	handlers      *Registry
}

func NewRouter(app fiber.Router, tokenVerifier token.Verifier, handlers *Registry) *Router {
	return &Router{
		app:           app,
		tokenVerifier: tokenVerifier,
		handlers:      handlers,
	}
}

func (r *Router) RegisterRoutes() {
	auth := r.app.Group("/auth")

	// 公开路由（不需要认证）：登录、刷新 Token
	if r.handlers.Auth != nil {
		auth.Post("/login/email", r.handlers.Auth.LoginByEmail)
		auth.Post("/login/phone", r.handlers.Auth.LoginByPhone)
		auth.Post("/refresh", r.handlers.Auth.Refresh)
	}

	// 需要认证的路由（需要token）
	authGroup := auth.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 会话相关
		authGroup.Post("/sessions", r.handlers.Session.Create)
		authGroup.Get("/sessions/:id", r.handlers.Session.GetByID)
		authGroup.Put("/sessions/:session_id/revoke", r.handlers.Session.Revoke)
		authGroup.Delete("/sessions/:id", r.handlers.Session.Delete)

		// 用户凭证相关
		authGroup.Post("/user-credentials", r.handlers.UserCredential.Create)
		authGroup.Get("/user-credentials/:id", r.handlers.UserCredential.GetByID)
		authGroup.Put("/user-credentials/:id", r.handlers.UserCredential.Update)
		authGroup.Delete("/user-credentials/:id", r.handlers.UserCredential.Delete)

		// MFA 因子相关
		authGroup.Post("/mfa-factors", r.handlers.MFAFactor.Create)
		authGroup.Get("/mfa-factors/:id", r.handlers.MFAFactor.GetByID)
		authGroup.Put("/mfa-factors/:id", r.handlers.MFAFactor.Update)
		authGroup.Delete("/mfa-factors/:id", r.handlers.MFAFactor.Delete)

		// 刷新令牌相关
		authGroup.Post("/refresh-tokens", r.handlers.RefreshToken.Create)
		authGroup.Get("/refresh-tokens/:id", r.handlers.RefreshToken.GetByID)
		authGroup.Put("/refresh-tokens/:id", r.handlers.RefreshToken.Update)
		authGroup.Delete("/refresh-tokens/:id", r.handlers.RefreshToken.Delete)

		// 密码重置相关
		authGroup.Post("/password-resets", r.handlers.PasswordReset.Create)
		authGroup.Get("/password-resets/:id", r.handlers.PasswordReset.GetByID)
		authGroup.Put("/password-resets/:id", r.handlers.PasswordReset.Update)
		authGroup.Delete("/password-resets/:id", r.handlers.PasswordReset.Delete)

		// 密码历史相关
		authGroup.Post("/password-history", r.handlers.PasswordHistory.Create)
		authGroup.Get("/password-history/:id", r.handlers.PasswordHistory.GetByID)

		// 登录尝试相关
		authGroup.Post("/login-attempts", r.handlers.LoginAttempt.Create)
		authGroup.Get("/login-attempts/:id", r.handlers.LoginAttempt.GetByID)
		authGroup.Delete("/login-attempts/:id", r.handlers.LoginAttempt.Delete)

		// 账户锁定相关
		authGroup.Post("/account-lockouts", r.handlers.AccountLockout.Create)
		authGroup.Get("/account-lockouts/:user_id", r.handlers.AccountLockout.GetByUserID)
		authGroup.Put("/account-lockouts/unlock", r.handlers.AccountLockout.Unlock)
		authGroup.Delete("/account-lockouts/:user_id", r.handlers.AccountLockout.Delete)

		// 受信任设备相关
		authGroup.Post("/trusted-devices", r.handlers.TrustedDevice.Create)
		authGroup.Get("/trusted-devices/:id", r.handlers.TrustedDevice.GetByID)
		authGroup.Delete("/trusted-devices/:id", r.handlers.TrustedDevice.Delete)
	}
}
