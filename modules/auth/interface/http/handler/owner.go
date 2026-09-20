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

type OwnerHandler struct {
	svc *platform.Service
}

func NewOwnerHandler(svc *platform.Service) *OwnerHandler {
	return &OwnerHandler{svc: svc}
}

func (h *OwnerHandler) ListForgerProfiles(c fiber.Ctx) error {
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

func (h *OwnerHandler) ListAuthorityProfiles(c fiber.Ctx) error {
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

func (h *OwnerHandler) UpdateAuthorityRoles(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("profileId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidProfileID)
	}
	var req reqdto.UpdateAuthorityRoles
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdateAuthorityRoles(c.Context(), aid, pid, req.AuthorityRoles); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}
