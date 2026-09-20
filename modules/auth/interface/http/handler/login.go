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

type LoginHandler struct {
	svc *platform.Service
}

func NewLoginHandler(svc *platform.Service) *LoginHandler {
	return &LoginHandler{svc: svc}
}

func (h *LoginHandler) WithEmail(c fiber.Ctx) error {
	var req reqdto.LoginWithEmail
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.LoginWithEmail(c.Context(), req.Email, req.Password, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *LoginHandler) WithPhone(c fiber.Ctx) error {
	var req reqdto.LoginWithPhone
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.LoginWithPhone(c.Context(), req.Phone, req.Password, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *LoginHandler) Refresh(c fiber.Ctx) error {
	var req reqdto.Refresh
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.Refresh(c.Context(), req.RefreshToken, req.DeviceID)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *LoginHandler) Logout(c fiber.Ctx) error {
	var req reqdto.Logout
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.Logout(c.Context(), req.RefreshToken); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *LoginHandler) SelectProfile(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.SelectProfile
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
