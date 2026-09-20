package handler

import (
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type SignupHandler struct {
	svc *platform.Service
}

func NewSignupHandler(svc *platform.Service) *SignupHandler {
	return &SignupHandler{svc: svc}
}

func (h *SignupHandler) WithEmail(c fiber.Ctx) error {
	var req reqdto.SignupWithEmail
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	out, err := h.svc.SignupWithEmail(c.Context(), req.Email, req.Password, req.VerificationCode, req.Lang, req.DeviceID, req.SignupPlatform)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
}

func (h *SignupHandler) SendCode(c fiber.Ctx) error {
	var req reqdto.SendSignupCode
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.SendSignupCode(c.Context(), req.Email, req.Lang); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"sent": true}})
}
