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
	signup.Post("/send-code", r.handlers.Signup.BySendingCode)
	signup.Post("/with-email", r.handlers.Signup.WithEmail)
}

func (r *Router) RegisterSessionGroup(auth fiber.Router) {
	auth.Post("/refresh", r.handlers.Login.Refresh)
	auth.Post("/logout", r.handlers.Login.Logout)
}

func (r *Router) RegisterMeGroup(auth fiber.Router) {
	me := auth.Group("/me", middleware.TokenAuth(r.verifier))
	me.Post("/select-profile", r.handlers.Login.SelectProfile)

	me.Get("/full-account-information-with-forger-profile", r.handlers.Account.GetFullAccountInformationWithCommunityProfile)
	me.Patch("/forger-profile", r.handlers.Account.PatchCommunityProfile)
	me.Patch("/forger-profile-settings", r.handlers.Account.PatchCommunityProfileSettings)
	me.Put("/forger-profile/avatars", r.handlers.Account.ConfirmCommunityProfileAvatar)
	me.Delete("/forger-profile/avatars", r.handlers.Account.ClearCommunityProfileAvatar)
	me.Put("/forger-profile/backgrounds", r.handlers.Account.ConfirmCommunityProfileBackgrounds)
	me.Put("/forger-profile/preference", r.handlers.Account.UpdatePreference)
	me.Get("/profiles", r.handlers.Account.ListCommunityProfiles)
	me.Post("/profiles", r.handlers.Account.CreateCommunityProfile)
	me.Post("/profiles/search", r.handlers.Account.SearchCommunityProfiles)
	me.Delete("/profiles/:profileId", r.handlers.Account.DeleteCommunityProfile)
	me.Get("/profiles/:profileId/public-card", r.handlers.Account.GetPublicProfileCard)

	me.Get("/full-account-information-with-authority-profile", r.handlers.AuthorityAccount.GetFullAccountInformationWithAuthorityProfile)
	me.Patch("/authority-profile", r.handlers.AuthorityAccount.PatchAuthorityProfile)
	me.Patch("/authority-profile-settings", r.handlers.AuthorityAccount.PatchAuthorityProfileSettings)
	me.Put("/authority-profile/avatars", r.handlers.AuthorityAccount.ConfirmAuthorityProfileAvatar)
	me.Delete("/authority-profile/avatars", r.handlers.AuthorityAccount.ClearAuthorityProfileAvatar)
	me.Put("/authority-profile/backgrounds", r.handlers.AuthorityAccount.ConfirmAuthorityProfileBackgrounds)
	me.Put("/authority-profile/preference", r.handlers.AuthorityAccount.UpdateAuthorityPreference)
	me.Get("/authority-profiles", r.handlers.AuthorityAccount.ListAuthorityProfiles)
	me.Post("/authority-profiles", r.handlers.AuthorityAccount.CreateAuthorityProfile)
	me.Post("/authority-profiles/search", r.handlers.AuthorityAccount.SearchAuthorityProfiles)
	me.Delete("/authority-profiles/:profileId", r.handlers.AuthorityAccount.DeleteAuthorityProfile)

	me.Get("/emails", r.handlers.Account.ListEmails)
	me.Post("/emails", r.handlers.Account.CreateEmail)
	me.Post("/emails/:emailId/send-verification-code", r.handlers.Account.SendEmailVerificationCode)
	me.Post("/emails/:emailId/verify", r.handlers.Account.VerifyEmail)
	me.Patch("/emails/:emailId", r.handlers.Account.UpdateEmail)
	me.Put("/emails/:emailId/primary", r.handlers.Account.SetPrimaryEmail)
	me.Delete("/emails/:emailId", r.handlers.Account.DeleteEmail)
	me.Get("/phones", r.handlers.Account.ListPhones)
	me.Post("/phones", r.handlers.Account.CreatePhone)
	me.Post("/phones/:phoneId/send-verification-code", r.handlers.Account.SendPhoneVerificationCode)
	me.Post("/phones/:phoneId/verify", r.handlers.Account.VerifyPhone)
	me.Patch("/phones/:phoneId", r.handlers.Account.UpdatePhone)
	me.Put("/phones/:phoneId/primary", r.handlers.Account.SetPrimaryPhone)
	me.Delete("/phones/:phoneId", r.handlers.Account.DeletePhone)
	me.Post("/password/send-verification-code", r.handlers.Account.SendChangePasswordVerificationCode)
	me.Put("/password", r.handlers.Account.ChangePassword)
}

func (r *Router) RegisterOwnerGroup(auth fiber.Router) {
	owner := auth.Group("/owner", middleware.TokenAuth(r.verifier))
	owner.Get("/forger-profiles", r.handlers.Owner.ListForgerProfiles)
	owner.Get("/authority-profiles", r.handlers.Owner.ListAuthorityProfiles)
	owner.Patch("/authority-profiles/:profileId/roles", r.handlers.Owner.UpdateAuthorityRoles)
}
