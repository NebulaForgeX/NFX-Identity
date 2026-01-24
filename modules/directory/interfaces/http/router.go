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
	directory := r.app.Group("/directory")

	// 需要认证的路由（需要token）
	auth := directory.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 用户相关
		auth.Post("/users", r.handlers.User.Create)
		auth.Get("/users/username/:username", r.handlers.User.GetByUsername)
		auth.Get("/users/:id/user-emails", r.handlers.UserEmail.GetByUserID) // 根据用户ID获取用户邮箱列表（必须在 /users/:id 之前）
		auth.Get("/users/:id", r.handlers.User.GetByID)
		auth.Patch("/users/:id/status", r.handlers.User.UpdateStatus)
		auth.Patch("/users/:id/username", r.handlers.User.UpdateUsername)
		auth.Patch("/users/:id/verify", r.handlers.User.Verify)
		auth.Delete("/users/:id", r.handlers.User.Delete)

		// 徽章相关
		auth.Post("/badges", r.handlers.Badge.Create)
		auth.Get("/badges/:id", r.handlers.Badge.GetByID)
		auth.Get("/badges/name/:name", r.handlers.Badge.GetByName)
		auth.Put("/badges/:id", r.handlers.Badge.Update)
		auth.Delete("/badges/:id", r.handlers.Badge.Delete)

		// 用户徽章相关
		auth.Post("/user-badges", r.handlers.UserBadge.Create)
		auth.Get("/user-badges/:id", r.handlers.UserBadge.GetByID)
		auth.Delete("/user-badges/:id", r.handlers.UserBadge.Delete)

		// 用户教育相关
		auth.Post("/user-educations", r.handlers.UserEducation.Create)
		auth.Get("/user-educations/:id", r.handlers.UserEducation.GetByID)
		auth.Put("/user-educations/:id", r.handlers.UserEducation.Update)
		auth.Delete("/user-educations/:id", r.handlers.UserEducation.Delete)

		// 用户邮箱相关
		auth.Post("/user-emails", r.handlers.UserEmail.Create)
		auth.Get("/user-emails/:id", r.handlers.UserEmail.GetByID)
		auth.Put("/user-emails/:id", r.handlers.UserEmail.Update)
		auth.Patch("/user-emails/:id/set-primary", r.handlers.UserEmail.SetPrimary)
		auth.Patch("/user-emails/:id/verify", r.handlers.UserEmail.Verify)
		auth.Delete("/user-emails/:id", r.handlers.UserEmail.Delete)

		// 用户职业相关
		auth.Post("/user-occupations", r.handlers.UserOccupation.Create)
		auth.Get("/user-occupations/:id", r.handlers.UserOccupation.GetByID)
		auth.Put("/user-occupations/:id", r.handlers.UserOccupation.Update)
		auth.Delete("/user-occupations/:id", r.handlers.UserOccupation.Delete)

		// 用户电话相关
		auth.Post("/user-phones", r.handlers.UserPhone.Create)
		auth.Get("/user-phones/:id", r.handlers.UserPhone.GetByID)
		auth.Put("/user-phones/:id", r.handlers.UserPhone.Update)
		auth.Patch("/user-phones/:id/set-primary", r.handlers.UserPhone.SetPrimary)
		auth.Patch("/user-phones/:id/verify", r.handlers.UserPhone.Verify)
		auth.Delete("/user-phones/:id", r.handlers.UserPhone.Delete)

		// 用户偏好相关
		auth.Post("/user-preferences", r.handlers.UserPreference.Create)
		auth.Get("/user-preferences/:id", r.handlers.UserPreference.GetByID)
		auth.Put("/user-preferences/:id", r.handlers.UserPreference.Update)
		auth.Delete("/user-preferences/:id", r.handlers.UserPreference.Delete)

		// 用户资料相关
		auth.Post("/user-profiles", r.handlers.UserProfile.Create)
		auth.Get("/user-profiles/:id", r.handlers.UserProfile.GetByID)
		auth.Put("/user-profiles/:id", r.handlers.UserProfile.Update)
		auth.Delete("/user-profiles/:id", r.handlers.UserProfile.Delete)
	}
}
