package http

import (
	"nfxid/pkgs/security/token"

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

	// 公开路由（仅注册页面用到的：LoginPage 登录/注册/验证码/刷新）
	if r.handlers.Auth != nil {
		auth.Post("/login/email", r.handlers.Auth.LoginByEmail)
		auth.Post("/login/phone", r.handlers.Auth.LoginByPhone)
		auth.Post("/refresh", r.handlers.Auth.Refresh)
		auth.Post("/send-verification-code", r.handlers.Auth.SendVerificationCode)
		auth.Post("/signup", r.handlers.Auth.Signup)
	}
}
