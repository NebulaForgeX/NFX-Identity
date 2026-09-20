package handler

import (
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type AuthorityHandler struct{ profileHandler }

func NewAuthorityHandler(svc *platform.Service) *AuthorityHandler {
	return &AuthorityHandler{profileHandler{svc: svc, kind: "authority"}}
}

func (h *AuthorityHandler) GetFullAccountInformation(c fiber.Ctx) error { return h.full(c) }
func (h *AuthorityHandler) PatchProfile(c fiber.Ctx) error              { return h.patch(c) }
func (h *AuthorityHandler) PatchSettings(c fiber.Ctx) error             { return h.patchSettings(c) }
func (h *AuthorityHandler) ConfirmAvatar(c fiber.Ctx) error             { return h.confirmAvatar(c) }
func (h *AuthorityHandler) ClearAvatar(c fiber.Ctx) error               { return h.clearAvatar(c) }
func (h *AuthorityHandler) ConfirmBackgrounds(c fiber.Ctx) error        { return h.confirmBackgrounds(c) }
func (h *AuthorityHandler) UpdatePreference(c fiber.Ctx) error          { return h.preference(c) }
func (h *AuthorityHandler) ListProfiles(c fiber.Ctx) error              { return h.list(c) }
func (h *AuthorityHandler) SearchProfiles(c fiber.Ctx) error            { return h.search(c) }
func (h *AuthorityHandler) DeleteProfile(c fiber.Ctx) error             { return h.deleteProfile(c) }

func (h *AuthorityHandler) CreateProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.CreateProfile
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreateAuthorityProfile(c.Context(), aid, req.DisplayName, req.ProfileLanguage)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"profile_id": id}})
}
