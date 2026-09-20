package http

import (
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	app      fiber.Router
	handlers *Registry
	verifier token.Verifier
}

func NewRouter(app fiber.Router, handlers *Registry, verifier token.Verifier) *Router {
	return &Router{app: app, handlers: handlers, verifier: verifier}
}

func (r *Router) RegisterRoutes() {
	auth := r.app.Group("/auth")
	r.RegisterLocalesGroup(auth)
	r.RegisterLoginGroup(auth)
	r.RegisterSignUpGroup(auth)
	r.RegisterSessionGroup(auth)
	r.RegisterMeGroup(auth)
	r.RegisterOwnerGroup(auth)
	auth.Get("/health", func(c fiber.Ctx) error {
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]string{"service": "auth"}})
	})
}

func (r *Router) RegisterLocalesGroup(auth fiber.Router) {
	locales := auth.Group("/locales")
	locales.Get("/:lang", r.handlers.Locales.GetErrorsByLang)
	messages := auth.Group("/messages")
	messages.Get("/:lang", r.handlers.Locales.GetMessagesByLang)
}

func (r *Router) RegisterLoginGroup(auth fiber.Router) {
	login := auth.Group("/login")
	login.Post("/with-email", r.handlers.Login.WithEmail)
	login.Post("/with-phone", r.handlers.Login.WithPhone)
}

func (r *Router) RegisterSignUpGroup(auth fiber.Router) {
	signup := auth.Group("/signup")
	signup.Post("/send-code", r.handlers.Signup.SendCode)
	signup.Post("/with-email", r.handlers.Signup.WithEmail)
}

func (r *Router) RegisterSessionGroup(auth fiber.Router) {
	auth.Post("/refresh", r.handlers.Login.Refresh)
	auth.Post("/logout", r.handlers.Login.Logout)
}

func (r *Router) RegisterMeGroup(auth fiber.Router) {
	me := auth.Group("/me", middleware.TokenAuth(r.verifier))
	me.Post("/select-profile", r.handlers.Login.SelectProfile)

	me.Get("/full-account-information-with-forger-profile", r.handlers.Forger.GetFullAccountInformation)
	me.Patch("/forger-profile", r.handlers.Forger.PatchProfile)
	me.Patch("/forger-profile-settings", r.handlers.Forger.PatchSettings)
	me.Put("/forger-profile/avatars", r.handlers.Forger.ConfirmAvatar)
	me.Delete("/forger-profile/avatars", r.handlers.Forger.ClearAvatar)
	me.Put("/forger-profile/backgrounds", r.handlers.Forger.ConfirmBackgrounds)
	me.Put("/forger-profile/preference", r.handlers.Forger.UpdatePreference)
	me.Get("/profiles", r.handlers.Forger.ListProfiles)
	me.Post("/profiles", r.handlers.Forger.CreateProfile)
	me.Post("/profiles/search", r.handlers.Forger.SearchProfiles)
	me.Delete("/profiles/:profileId", r.handlers.Forger.DeleteProfile)
	me.Get("/profiles/:profileId/public-card", r.handlers.Forger.PublicCard)

	me.Get("/full-account-information-with-authority-profile", r.handlers.Authority.GetFullAccountInformation)
	me.Patch("/authority-profile", r.handlers.Authority.PatchProfile)
	me.Patch("/authority-profile-settings", r.handlers.Authority.PatchSettings)
	me.Put("/authority-profile/avatars", r.handlers.Authority.ConfirmAvatar)
	me.Delete("/authority-profile/avatars", r.handlers.Authority.ClearAvatar)
	me.Put("/authority-profile/backgrounds", r.handlers.Authority.ConfirmBackgrounds)
	me.Put("/authority-profile/preference", r.handlers.Authority.UpdatePreference)
	me.Get("/authority-profiles", r.handlers.Authority.ListProfiles)
	me.Post("/authority-profiles", r.handlers.Authority.CreateProfile)
	me.Post("/authority-profiles/search", r.handlers.Authority.SearchProfiles)
	me.Delete("/authority-profiles/:profileId", r.handlers.Authority.DeleteProfile)

	me.Get("/emails", r.handlers.Account.ListEmails)
	me.Post("/emails", r.handlers.Account.CreateEmail)
	me.Post("/emails/:emailId/send-verification-code", r.handlers.Account.SendEmailCode)
	me.Post("/emails/:emailId/verify", r.handlers.Account.VerifyEmail)
	me.Patch("/emails/:emailId", r.handlers.Account.UpdateEmail)
	me.Put("/emails/:emailId/primary", r.handlers.Account.SetPrimaryEmail)
	me.Delete("/emails/:emailId", r.handlers.Account.DeleteEmail)
	me.Get("/phones", r.handlers.Account.ListPhones)
	me.Post("/phones", r.handlers.Account.CreatePhone)
	me.Post("/phones/:phoneId/send-verification-code", r.handlers.Account.SendPhoneCode)
	me.Post("/phones/:phoneId/verify", r.handlers.Account.VerifyPhone)
	me.Patch("/phones/:phoneId", r.handlers.Account.UpdatePhone)
	me.Put("/phones/:phoneId/primary", r.handlers.Account.SetPrimaryPhone)
	me.Delete("/phones/:phoneId", r.handlers.Account.DeletePhone)
	me.Post("/password/send-verification-code", r.handlers.Account.SendPasswordCode)
	me.Put("/password", r.handlers.Account.ChangePassword)
}

func (r *Router) RegisterOwnerGroup(auth fiber.Router) {
	owner := auth.Group("/owner", middleware.TokenAuth(r.verifier))
	owner.Get("/forger-profiles", r.handlers.Owner.ListForgerProfiles)
	owner.Get("/authority-profiles", r.handlers.Owner.ListAuthorityProfiles)
	owner.Patch("/authority-profiles/:profileId/roles", r.handlers.Owner.UpdateAuthorityRoles)
}
