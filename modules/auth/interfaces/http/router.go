package http

import (
	"nfxid/modules/auth/interfaces/http/middleware"
	"nfxid/pkgs/tokenx"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app      fiber.Router
	tokenx   *tokenx.Tokenx
	handlers *Registry
}

func NewRouter(app fiber.Router, tokenx *tokenx.Tokenx, handlers *Registry) *Router {
	return &Router{
		app:      app,
		tokenx:   tokenx,
		handlers: handlers,
	}
}

func (r *Router) RegisterRoutes() {
	api := r.app.Group("/api/v1")

	// ========== 公开路由（不需要 token） ==========
	public := api.Group("/auth")
	{
		// 用户相关
		public.Post("/login", r.handlers.User.Login)
		public.Post("/refresh", r.handlers.User.RefreshToken)
		public.Post("/verification-code", r.handlers.User.SendVerificationCode) // 发送验证码
	}

	// ========== 需要认证的路由（需要 token） ==========
	auth := api.Group("/auth", middleware.AccessTokenMiddleware(r.tokenx))
	{
		// 用户相关
		auth.Post("/users", r.handlers.User.Create)
		auth.Get("/users", r.handlers.User.GetAll)
		auth.Get("/users/:id", r.handlers.User.GetByID)
		auth.Put("/users/:id", r.handlers.User.Update)
		auth.Delete("/users/:id", r.handlers.User.Delete)
		auth.Delete("/users/:id/account", r.handlers.User.DeleteAccount)

		// 资料相关
		auth.Post("/profiles", r.handlers.Profile.Create)
		auth.Get("/profiles", r.handlers.Profile.GetAll)
		auth.Get("/profiles/:id", r.handlers.Profile.GetByID)
		auth.Get("/profiles/user/:user_id", r.handlers.Profile.GetByUserID)
		auth.Put("/profiles/:id", r.handlers.Profile.Update)
		auth.Delete("/profiles/:id", r.handlers.Profile.Delete)

		// 角色相关
		auth.Post("/roles", r.handlers.Role.Create)
		auth.Get("/roles", r.handlers.Role.GetAll)
		auth.Get("/roles/:id", r.handlers.Role.GetByID)
		auth.Get("/roles/name/:name", r.handlers.Role.GetByName)
		auth.Put("/roles/:id", r.handlers.Role.Update)
		auth.Delete("/roles/:id", r.handlers.Role.Delete)

		// 徽章相关
		auth.Post("/badges", r.handlers.Badge.Create)
		auth.Get("/badges", r.handlers.Badge.GetAll)
		auth.Get("/badges/:id", r.handlers.Badge.GetByID)
		auth.Get("/badges/name/:name", r.handlers.Badge.GetByName)
		auth.Put("/badges/:id", r.handlers.Badge.Update)
		auth.Delete("/badges/:id", r.handlers.Badge.Delete)

		// 教育经历相关
		auth.Post("/educations", r.handlers.Education.Create)
		auth.Get("/educations", r.handlers.Education.GetAll)
		auth.Get("/educations/:id", r.handlers.Education.GetByID)
		auth.Get("/educations/profile/:profile_id", r.handlers.Education.GetByProfileID)
		auth.Put("/educations/:id", r.handlers.Education.Update)
		auth.Delete("/educations/:id", r.handlers.Education.Delete)

		// 职业信息相关
		auth.Post("/occupations", r.handlers.Occupation.Create)
		auth.Get("/occupations", r.handlers.Occupation.GetAll)
		auth.Get("/occupations/:id", r.handlers.Occupation.GetByID)
		auth.Get("/occupations/profile/:profile_id", r.handlers.Occupation.GetByProfileID)
		auth.Put("/occupations/:id", r.handlers.Occupation.Update)
		auth.Delete("/occupations/:id", r.handlers.Occupation.Delete)

		// 用户徽章关联相关
		auth.Post("/profile-badges", r.handlers.ProfileBadge.Create)
		auth.Get("/profile-badges/:id", r.handlers.ProfileBadge.GetByID)
		auth.Get("/profile-badges/profile/:profile_id", r.handlers.ProfileBadge.GetByProfileID)
		auth.Get("/profile-badges/badge/:badge_id", r.handlers.ProfileBadge.GetByBadgeID)
		auth.Put("/profile-badges/:id", r.handlers.ProfileBadge.Update)
		auth.Delete("/profile-badges/:id", r.handlers.ProfileBadge.Delete)
		auth.Delete("/profile-badges/profile/:profile_id/badge/:badge_id", r.handlers.ProfileBadge.DeleteByProfileAndBadge)
	}
}
