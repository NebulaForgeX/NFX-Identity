package handler

import (
	"strconv"

	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type profileHandler struct {
	svc  *platform.Service
	kind string
}

func (h *profileHandler) full(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	out, err := h.svc.FullAccountWithProfile(c.Context(), aid, pid, h.kind)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *profileHandler) patch(c fiber.Ctx) error {
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
	if err := h.svc.PatchProfile(c.Context(), aid, pid, h.kind, req); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *profileHandler) patchSettings(c fiber.Ctx) error {
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.PatchProfileSettings
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.PatchProfileSettings(c.Context(), pid, h.kind, req.LoginNotification); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *profileHandler) confirmAvatar(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.ConfirmAvatar
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.ConfirmAvatar(c.Context(), aid, pid, h.kind, req.ImageID); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *profileHandler) clearAvatar(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	if err := h.svc.ClearAvatar(c.Context(), aid, pid, h.kind); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *profileHandler) confirmBackgrounds(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.ConfirmBackgrounds
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
	if err := h.svc.ConfirmBackgrounds(c.Context(), aid, pid, h.kind, items); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *profileHandler) preference(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := profileID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.Preference
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdatePreference(c.Context(), aid, pid, h.kind, req.Preference); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *profileHandler) list(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	items, err := h.svc.ListAccountProfiles(c.Context(), aid, h.kind)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": len(items)}})
}

func (h *profileHandler) search(c fiber.Ctx) error {
	var req reqdto.SearchProfiles
	if err := c.Bind().Body(&req); err != nil {
		req.Query = c.Query("query")
		req.Limit, _ = strconv.Atoi(c.Query("limit", "20"))
		req.Offset, _ = strconv.Atoi(c.Query("offset", "0"))
	}
	items, total, err := h.svc.SearchProfiles(c.Context(), h.kind, req.Query, req.Limit, req.Offset)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"items": items, "total": total}})
}

func (h *profileHandler) deleteProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("profileId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidProfileID)
	}
	if err := h.svc.DeleteProfile(c.Context(), aid, pid, h.kind); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}
