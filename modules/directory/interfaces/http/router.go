package http

import (
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
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
	auth := directory.Group("/auth", middleware.TokenAuth(r.tokenVerifier))
	{
		// 用户相关（Rex 规范：:user_id）
		auth.Post("/users", r.handlers.User.Create)
		auth.Get("/users/username/:username", r.handlers.User.GetByUsername)
		auth.Get("/users/:user_id/user-emails", r.handlers.UserEmail.GetByUserID)
		auth.Get("/users/:user_id/user-phones", r.handlers.UserPhone.GetByUserID)
		auth.Get("/users/:user_id/user-educations", r.handlers.UserEducation.GetByUserID)
		auth.Get("/users/:user_id/user-occupations", r.handlers.UserOccupation.GetByUserID)
		auth.Get("/users/:user_id", r.handlers.User.GetByID)
		auth.Patch("/users/:user_id/status", r.handlers.User.UpdateStatus)
		auth.Patch("/users/:user_id/username", r.handlers.User.UpdateUsername)
		auth.Patch("/users/:user_id/verify", r.handlers.User.Verify)
		auth.Delete("/users/:user_id", r.handlers.User.Delete)

		// 用户教育相关
		auth.Post("/user-educations", r.handlers.UserEducation.Create)
		auth.Get("/user-educations/:user_education_id", r.handlers.UserEducation.GetByID)
		auth.Put("/user-educations/:user_education_id", r.handlers.UserEducation.Update)
		auth.Delete("/user-educations/:user_education_id", r.handlers.UserEducation.Delete)

		// 用户邮箱相关
		auth.Post("/user-emails", r.handlers.UserEmail.Create)
		auth.Get("/user-emails/:user_email_id", r.handlers.UserEmail.GetByID)
		auth.Put("/user-emails/:user_email_id", r.handlers.UserEmail.Update)
		auth.Patch("/user-emails/:user_email_id/set-primary", r.handlers.UserEmail.SetPrimary)
		auth.Patch("/user-emails/:user_email_id/verify", r.handlers.UserEmail.Verify)
		auth.Delete("/user-emails/:user_email_id", r.handlers.UserEmail.Delete)

		// 用户职业相关
		auth.Post("/user-occupations", r.handlers.UserOccupation.Create)
		auth.Get("/user-occupations/:user_occupation_id", r.handlers.UserOccupation.GetByID)
		auth.Put("/user-occupations/:user_occupation_id", r.handlers.UserOccupation.Update)
		auth.Delete("/user-occupations/:user_occupation_id", r.handlers.UserOccupation.Delete)

		// 用户电话相关
		auth.Post("/user-phones", r.handlers.UserPhone.Create)
		auth.Get("/user-phones/:user_phone_id", r.handlers.UserPhone.GetByID)
		auth.Put("/user-phones/:user_phone_id", r.handlers.UserPhone.Update)
		auth.Patch("/user-phones/:user_phone_id/set-primary", r.handlers.UserPhone.SetPrimary)
		auth.Patch("/user-phones/:user_phone_id/verify", r.handlers.UserPhone.Verify)
		auth.Delete("/user-phones/:user_phone_id", r.handlers.UserPhone.Delete)

		// 用户偏好相关
		auth.Post("/user-preferences", r.handlers.UserPreference.Create)
		auth.Get("/user-preferences/:user_preference_id", r.handlers.UserPreference.GetByID)
		auth.Put("/user-preferences/:user_preference_id", r.handlers.UserPreference.Update)
		auth.Delete("/user-preferences/:user_preference_id", r.handlers.UserPreference.Delete)

		// 用户资料相关
		auth.Post("/user-profiles", r.handlers.UserProfile.Create)
		auth.Get("/user-profiles/:user_profile_id", r.handlers.UserProfile.GetByID)
		auth.Put("/user-profiles/:user_profile_id", r.handlers.UserProfile.Update)
		auth.Delete("/user-profiles/:user_profile_id", r.handlers.UserProfile.Delete)

		// 用户头像相关（1:1 与 user，均用 user_id）
		auth.Post("/user-avatars", r.handlers.UserAvatar.CreateOrUpdate)
		auth.Get("/user-avatars/user/:user_id", r.handlers.UserAvatar.GetByUserID)
		auth.Put("/user-avatars/:user_id", r.handlers.UserAvatar.Update)
		auth.Delete("/user-avatars/:user_id", r.handlers.UserAvatar.Delete)

		// 用户图片相关
		auth.Post("/user-images", r.handlers.UserImage.Create)
		auth.Get("/user-images/:user_image_id", r.handlers.UserImage.GetByID)
		auth.Get("/users/:user_id/user-images", r.handlers.UserImage.GetByUserID)
		auth.Get("/users/:user_id/user-images/current", r.handlers.UserImage.GetCurrent)
		auth.Put("/user-images/:user_image_id", r.handlers.UserImage.Update)
		auth.Patch("/user-images/:user_image_id/set-primary", r.handlers.UserImage.SetPrimary)
		auth.Patch("/user-images/:user_image_id/display-order", r.handlers.UserImage.UpdateDisplayOrder)
		auth.Patch("/users/:user_id/user-images/display-order", r.handlers.UserImage.UpdateDisplayOrderBatch)
		auth.Delete("/user-images/:user_image_id", r.handlers.UserImage.Delete)
	}
}
