package handler

import (
	"strconv"

	sysErr "nfxidentity/errors/src/sys"
	authmsg "nfxidentity/messages/src/auth"
	"nfxidentity/modules/auth/application/account"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type OwnerHandler struct {
	svc *account.Service
}

func NewOwnerHandler(svc *account.Service) *OwnerHandler {
	return &OwnerHandler{svc: svc}
}

func (h *OwnerHandler) ListForgerProfiles(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	out, err := h.svc.ListOwnerForgerProfiles(c.Context(), account.ListOwnerForgerProfilesInput{
		AccountID: accountID,
		Query:     c.Query("query"),
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *OwnerHandler) ListAuthorityProfiles(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	out, err := h.svc.ListOwnerAuthorityProfiles(c.Context(), account.ListOwnerAuthorityProfilesInput{
		AccountID: accountID,
		Query:     c.Query("query"),
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *OwnerHandler) UpdateAuthorityRoles(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.OwnerProfileURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if uri.ProfileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.UpdateAuthorityRoles
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	if err := h.svc.UpdateAuthorityRoles(c.Context(), account.UpdateAuthorityRolesInput{
		ActorAccountID: accountID,
		ProfileID:      uri.ProfileID,
		Roles:          req.AuthorityRoles,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.AUTHORITY_ROLES_UPDATED, httpx.SuccessOptions{Data: nil})
}
