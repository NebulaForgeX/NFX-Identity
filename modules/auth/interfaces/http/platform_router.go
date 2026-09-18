package http

import (
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"path/filepath"
	"strconv"

	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Handler struct {
	svc *platform.Service
}

func NewHandler(svc *platform.Service) *Handler {
	return &Handler{svc: svc}
}

func RegisterRoutes(app fiber.Router, svc *platform.Service, verifier token.Verifier) {
	h := NewHandler(svc)
	auth := app.Group("/auth")
	auth.Post("/login/with-email", h.LoginWithEmail)
	auth.Post("/login/with-phone", h.LoginWithPhone)
	auth.Post("/login/github", h.LoginGitHub)
	auth.Get("/login/github/url", h.GitHubURL)
	auth.Post("/signup/send-code", h.SendCode)
	auth.Post("/signup/with-email", h.Signup)
	auth.Post("/refresh", h.Refresh)
	auth.Post("/logout", h.Logout)
	auth.Get("/health", func(c fiber.Ctx) error {
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]string{"service": "auth"}})
	})
	auth.Get("/locales/:lang", h.Locales)
	auth.Get("/messages/:lang", h.Messages)

	me := auth.Group("/me", middleware.TokenAuth(verifier))
	me.Post("/select-profile", h.SelectProfile)
	me.Get("/full-account-information-with-forger-profile", h.FullForger)
	me.Get("/full-account-information-with-authority-profile", h.FullAuthority)
	me.Patch("/forger-profile", h.patchForger)
	me.Patch("/authority-profile", h.patchAuthority)
	me.Patch("/forger-profile-settings", h.patchForgerSettings)
	me.Patch("/authority-profile-settings", h.patchAuthoritySettings)
	me.Put("/forger-profile/avatars", h.confirmForgerAvatar)
	me.Delete("/forger-profile/avatars", h.clearForgerAvatar)
	me.Put("/authority-profile/avatars", h.confirmAuthorityAvatar)
	me.Delete("/authority-profile/avatars", h.clearAuthorityAvatar)
	me.Put("/forger-profile/backgrounds", h.confirmForgerBackgrounds)
	me.Put("/authority-profile/backgrounds", h.confirmAuthorityBackgrounds)
	me.Put("/forger-profile/preference", h.prefForger)
	me.Put("/authority-profile/preference", h.prefAuthority)
	me.Post("/github", h.LinkGitHub)
	me.Delete("/github", h.UnlinkGitHub)
	me.Get("/emails", h.ListEmails)
	me.Post("/emails", h.CreateEmail)
	me.Post("/emails/:emailId/send-verification-code", h.SendEmailCode)
	me.Post("/emails/:emailId/verify", h.VerifyEmail)
	me.Patch("/emails/:emailId", h.UpdateEmail)
	me.Put("/emails/:emailId/primary", h.SetPrimaryEmail)
	me.Delete("/emails/:emailId", h.DeleteEmail)
	me.Get("/phones", h.ListPhones)
	me.Post("/phones", h.CreatePhone)
	me.Post("/phones/:phoneId/send-verification-code", h.SendPhoneCode)
	me.Post("/phones/:phoneId/verify", h.VerifyPhone)
	me.Patch("/phones/:phoneId", h.UpdatePhone)
	me.Put("/phones/:phoneId/primary", h.SetPrimaryPhone)
	me.Delete("/phones/:phoneId", h.DeletePhone)
	me.Get("/profiles", h.ListForgerProfiles)
	me.Post("/profiles", h.CreateForgerProfile)
	me.Post("/profiles/search", h.SearchForger)
	me.Delete("/profiles/:profileId", h.DeleteForger)
	me.Get("/profiles/:profileId/public-card", h.PublicCard)
	me.Get("/authority-profiles", h.ListAuthorityProfiles)
	me.Post("/authority-profiles", h.CreateAuthorityProfile)
	me.Post("/authority-profiles/search", h.SearchAuthority)
	me.Delete("/authority-profiles/:profileId", h.DeleteAuthority)
	me.Post("/password/send-verification-code", h.SendPasswordCode)
	me.Put("/password", h.ChangePassword)

	owner := auth.Group("/owner", middleware.TokenAuth(verifier))
	owner.Get("/forger-profiles", h.OwnerForger)
	owner.Get("/authority-profiles", h.OwnerAuthority)
	owner.Patch("/authority-profiles/:profileId/roles", h.UpdateAuthorityRoles)
}

func wrap(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	if e := errx.AsError(err); e != nil {
		return fiberx.ErrorFromErrx(c, e)
	}
	return fiberx.ErrorFromErrx(c, sys.ErrInternal.WithCause(err))
}

func accountID(c fiber.Ctx) (uuid.UUID, error) {
	id, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return uuid.UUID{}, sys.ErrInvalidToken
	}
	return id, nil
}

func profileID(c fiber.Ctx) (uuid.UUID, error) {
	id, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok {
		return uuid.UUID{}, sys.ErrInvalidToken
	}
	return id, nil
}

func (h *Handler) LoginWithEmail(c fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		DeviceID string `json:"device_id"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.LoginWithEmail(c.Context(), req.Email, req.Password, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) LoginWithPhone(c fiber.Ctx) error {
	var req struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
		DeviceID string `json:"device_id"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.LoginWithPhone(c.Context(), req.Phone, req.Password, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) Signup(c fiber.Ctx) error {
	var req struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		VerificationCode string `json:"verification_code"`
		Lang             string `json:"lang"`
		DeviceID         string `json:"device_id"`
		SignupPlatform   string `json:"signup_platform"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.SignupWithEmail(c.Context(), req.Email, req.Password, req.VerificationCode, req.Lang, req.DeviceID, req.SignupPlatform)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) SendCode(c fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
		Lang  string `json:"lang"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.SendSignupCode(c.Context(), req.Email, req.Lang); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"sent": true}})
}

func (h *Handler) GitHubURL(c fiber.Ctx) error {
	url, state, err := h.svc.GitHubAuthorizeURL(c.Context())
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"authorize_url": url, "state": state}})
}

func (h *Handler) LoginGitHub(c fiber.Ctx) error {
	var req struct {
		Code           string `json:"code"`
		State          string `json:"state"`
		DeviceID       string `json:"device_id"`
		SignupPlatform string `json:"signup_platform"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.LoginWithGitHub(c.Context(), req.Code, req.State, req.DeviceID, req.SignupPlatform)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) LinkGitHub(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.LinkGitHub(c.Context(), aid, req.Code, req.State); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) UnlinkGitHub(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	if err := h.svc.UnlinkGitHub(c.Context(), aid); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) Refresh(c fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
		DeviceID     string `json:"device_id"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.Refresh(c.Context(), req.RefreshToken, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) Logout(c fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.Logout(c.Context(), req.RefreshToken); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) SelectProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		ProfileID string `json:"profile_id"`
		Kind      string `json:"kind"`
		DeviceID  string `json:"device_id"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	pid, err := uuid.Parse(req.ProfileID)
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidProfileID)
	}
	out, err := h.svc.SelectProfile(c.Context(), aid, pid, req.Kind, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) FullForger(c fiber.Ctx) error    { return h.full(c, "forger") }
func (h *Handler) FullAuthority(c fiber.Ctx) error { return h.full(c, "authority") }

func (h *Handler) full(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	out, err := h.svc.FullAccountWithProfile(c.Context(), aid, pid, kind)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *Handler) patchForger(c fiber.Ctx) error    { return h.patch(c, "forger") }
func (h *Handler) patchAuthority(c fiber.Ctx) error { return h.patch(c, "authority") }

func (h *Handler) patch(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req map[string]any
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.PatchProfile(c.Context(), aid, pid, kind, req); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) patchForgerSettings(c fiber.Ctx) error    { return h.patchSettings(c, "forger") }
func (h *Handler) patchAuthoritySettings(c fiber.Ctx) error { return h.patchSettings(c, "authority") }

func (h *Handler) patchSettings(c fiber.Ctx, kind string) error {
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		LoginNotification *bool `json:"login_notification"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.PatchProfileSettings(c.Context(), pid, kind, req.LoginNotification); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) confirmForgerAvatar(c fiber.Ctx) error {
	return h.confirmAvatar(c, "forger")
}
func (h *Handler) confirmAuthorityAvatar(c fiber.Ctx) error {
	return h.confirmAvatar(c, "authority")
}
func (h *Handler) clearForgerAvatar(c fiber.Ctx) error {
	return h.clearAvatar(c, "forger")
}
func (h *Handler) clearAuthorityAvatar(c fiber.Ctx) error {
	return h.clearAvatar(c, "authority")
}

func (h *Handler) confirmAvatar(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		ImageID string `json:"image_id"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.ConfirmAvatar(c.Context(), aid, pid, kind, req.ImageID); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) clearAvatar(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	if err := h.svc.ClearAvatar(c.Context(), aid, pid, kind); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) confirmForgerBackgrounds(c fiber.Ctx) error {
	return h.confirmBackgrounds(c, "forger")
}
func (h *Handler) confirmAuthorityBackgrounds(c fiber.Ctx) error {
	return h.confirmBackgrounds(c, "authority")
}

func (h *Handler) confirmBackgrounds(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		Images []struct {
			ImageID   string `json:"image_id"`
			SortOrder int    `json:"sort_order"`
		} `json:"images"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	items := make([]struct {
		ImageID   string
		SortOrder int
	}, 0, len(req.Images))
	for _, img := range req.Images {
		items = append(items, struct {
			ImageID   string
			SortOrder int
		}{ImageID: img.ImageID, SortOrder: img.SortOrder})
	}
	if err := h.svc.ConfirmBackgrounds(c.Context(), aid, pid, kind, items); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) prefForger(c fiber.Ctx) error    { return h.pref(c, "forger") }
func (h *Handler) prefAuthority(c fiber.Ctx) error { return h.pref(c, "authority") }

func (h *Handler) pref(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		Preference string `json:"preference"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdatePreference(c.Context(), aid, pid, kind, req.Preference); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) ListEmails(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	items, total, err := h.svc.ListEmails(c.Context(), aid)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": total}})
}

func (h *Handler) CreateEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		Email string `json:"email"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreateEmail(c.Context(), aid, req.Email)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"email_id": id}})
}

func (h *Handler) SendEmailCode(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	var req struct {
		Lang string `json:"lang"`
	}
	_ = c.Bind().Body(&req)
	if err := h.svc.SendEmailVerificationCode(c.Context(), aid, eid, req.Lang); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) VerifyEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	var req struct {
		VerificationCode string `json:"verification_code"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.VerifyEmail(c.Context(), aid, eid, req.VerificationCode); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) UpdateEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	var req struct {
		Email string `json:"email"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdateEmail(c.Context(), aid, eid, req.Email); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) SetPrimaryEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	if err := h.svc.SetPrimaryEmail(c.Context(), aid, eid); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) Locales(c fiber.Ctx) error {
	return fiberx.SendLangJSON(c, filepath.Join("errors", "langs"), c.Params("lang"))
}

func (h *Handler) Messages(c fiber.Ctx) error {
	return fiberx.SendLangJSON(c, filepath.Join("messages", "langs"), c.Params("lang"))
}

func (h *Handler) ListPhones(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	items, total, err := h.svc.ListPhones(c.Context(), aid)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": total}})
}

func (h *Handler) CreatePhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		Phone string `json:"phone"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreatePhone(c.Context(), aid, req.Phone)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"phone_id": id}})
}

func (h *Handler) SendPhoneCode(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	if err := h.svc.SendPhoneVerificationCode(c.Context(), aid, pid); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) VerifyPhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	var req struct {
		VerificationCode string `json:"verification_code"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.VerifyPhone(c.Context(), aid, pid, req.VerificationCode); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) UpdatePhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	var req struct {
		Phone string `json:"phone"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdatePhone(c.Context(), aid, pid, req.Phone); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) SetPrimaryPhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	if err := h.svc.SetPrimaryPhone(c.Context(), aid, pid); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) DeletePhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	if err := h.svc.DeletePhone(c.Context(), aid, pid); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) DeleteEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	if err := h.svc.DeleteEmail(c.Context(), aid, eid); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) ListForgerProfiles(c fiber.Ctx) error {
	return h.listProfiles(c, "forger")
}
func (h *Handler) ListAuthorityProfiles(c fiber.Ctx) error {
	return h.listProfiles(c, "authority")
}

func (h *Handler) listProfiles(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	items, err := h.svc.ListAccountProfiles(c.Context(), aid, kind)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": len(items)}})
}

func (h *Handler) CreateForgerProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		DisplayName     string `json:"display_name"`
		ProfileLanguage string `json:"profile_language"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreateForgerProfile(c.Context(), aid, req.DisplayName, req.ProfileLanguage)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"profile_id": id}})
}

func (h *Handler) CreateAuthorityProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		DisplayName     string `json:"display_name"`
		ProfileLanguage string `json:"profile_language"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreateAuthorityProfile(c.Context(), aid, req.DisplayName, req.ProfileLanguage)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"profile_id": id}})
}

func (h *Handler) SearchForger(c fiber.Ctx) error    { return h.search(c, "forger") }
func (h *Handler) SearchAuthority(c fiber.Ctx) error { return h.search(c, "authority") }

func (h *Handler) search(c fiber.Ctx, kind string) error {
	var req struct {
		Query  string `json:"query"`
		Limit  int    `json:"limit"`
		Offset int    `json:"offset"`
	}
	if err := c.Bind().Body(&req); err != nil {
		req.Query = c.Query("query")
		req.Limit, _ = strconv.Atoi(c.Query("limit", "20"))
		req.Offset, _ = strconv.Atoi(c.Query("offset", "0"))
	}
	items, total, err := h.svc.SearchProfiles(c.Context(), kind, req.Query, req.Limit, req.Offset)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": total}})
}

func (h *Handler) DeleteForger(c fiber.Ctx) error    { return h.deleteProfile(c, "forger") }
func (h *Handler) DeleteAuthority(c fiber.Ctx) error { return h.deleteProfile(c, "authority") }

func (h *Handler) deleteProfile(c fiber.Ctx, kind string) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("profileId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidProfileID)
	}
	if err := h.svc.DeleteProfile(c.Context(), aid, pid, kind); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) PublicCard(c fiber.Ctx) error {
	pid, err := uuid.Parse(c.Params("profileId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidProfileID)
	}
	card, err := h.svc.PublicProfileCard(c.Context(), pid)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: card})
}

func (h *Handler) SendPasswordCode(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		Lang string `json:"lang"`
	}
	_ = c.Bind().Body(&req)
	if err := h.svc.SendPasswordCode(c.Context(), aid, req.Lang); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) ChangePassword(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req struct {
		CurrentPassword  string `json:"current_password"`
		NewPassword      string `json:"new_password"`
		VerificationCode string `json:"verification_code"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.ChangePassword(c.Context(), aid, req.CurrentPassword, req.NewPassword, req.VerificationCode); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *Handler) OwnerForger(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	items, total, err := h.svc.ListOwnerForgerProfiles(c.Context(), aid, c.Query("query"), limit, offset)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": total}})
}

func (h *Handler) OwnerAuthority(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	items, total, err := h.svc.ListOwnerAuthorityProfiles(c.Context(), aid, c.Query("query"), limit, offset)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": total}})
}

func (h *Handler) UpdateAuthorityRoles(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("profileId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidProfileID)
	}
	var req struct {
		AuthorityRoles []string `json:"authority_roles"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdateAuthorityRoles(c.Context(), aid, pid, req.AuthorityRoles); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}
