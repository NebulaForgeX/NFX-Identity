package handler

import (
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ForgerHandler struct{ profileHandler }

func NewForgerHandler(svc *platform.Service) *ForgerHandler {
	return &ForgerHandler{profileHandler{svc: svc, kind: "forger"}}
}

func (h *ForgerHandler) GetFullAccountInformation(c fiber.Ctx) error { return h.full(c) }
func (h *ForgerHandler) PatchProfile(c fiber.Ctx) error              { return h.patch(c) }
func (h *ForgerHandler) PatchSettings(c fiber.Ctx) error             { return h.patchSettings(c) }
func (h *ForgerHandler) ConfirmAvatar(c fiber.Ctx) error             { return h.confirmAvatar(c) }
func (h *ForgerHandler) ClearAvatar(c fiber.Ctx) error               { return h.clearAvatar(c) }
func (h *ForgerHandler) ConfirmBackgrounds(c fiber.Ctx) error        { return h.confirmBackgrounds(c) }
func (h *ForgerHandler) UpdatePreference(c fiber.Ctx) error          { return h.preference(c) }
func (h *ForgerHandler) ListProfiles(c fiber.Ctx) error              { return h.list(c) }
func (h *ForgerHandler) SearchProfiles(c fiber.Ctx) error            { return h.search(c) }
func (h *ForgerHandler) DeleteProfile(c fiber.Ctx) error             { return h.deleteProfile(c) }

func (h *ForgerHandler) CreateProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.CreateProfile
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreateForgerProfile(c.Context(), aid, req.DisplayName, req.ProfileLanguage)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"profile_id": id}})
}

func (h *ForgerHandler) PublicCard(c fiber.Ctx) error {
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
