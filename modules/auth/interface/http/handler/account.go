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

type AccountHandler struct {
	svc *platform.Service
}

func NewAccountHandler(svc *platform.Service) *AccountHandler {
	return &AccountHandler{svc: svc}
}

func (h *AccountHandler) ListEmails(c fiber.Ctx) error {
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

func (h *AccountHandler) CreateEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.Email
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreateEmail(c.Context(), aid, req.Email)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"email_id": id}})
}

func (h *AccountHandler) SendEmailCode(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	var req reqdto.SendEmailCode
	_ = c.Bind().Body(&req)
	if err := h.svc.SendEmailVerificationCode(c.Context(), aid, eid, req.Lang); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *AccountHandler) VerifyEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	var req reqdto.VerifyCode
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.VerifyEmail(c.Context(), aid, eid, req.VerificationCode); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *AccountHandler) UpdateEmail(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	eid, err := uuid.Parse(c.Params("emailId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidEmailID)
	}
	var req reqdto.Email
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdateEmail(c.Context(), aid, eid, req.Email); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *AccountHandler) SetPrimaryEmail(c fiber.Ctx) error {
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

func (h *AccountHandler) DeleteEmail(c fiber.Ctx) error {
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

func (h *AccountHandler) ListPhones(c fiber.Ctx) error {
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

func (h *AccountHandler) CreatePhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.Phone
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	id, err := h.svc.CreatePhone(c.Context(), aid, req.Phone)
	if err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"phone_id": id}})
}

func (h *AccountHandler) SendPhoneCode(c fiber.Ctx) error {
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

func (h *AccountHandler) VerifyPhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	var req reqdto.VerifyCode
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.VerifyPhone(c.Context(), aid, pid, req.VerificationCode); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *AccountHandler) UpdatePhone(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	pid, err := uuid.Parse(c.Params("phoneId"))
	if err != nil {
		return fiberx.ErrorFromErrx(c, auth.ErrInvalidPhoneID)
	}
	var req reqdto.Phone
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.UpdatePhone(c.Context(), aid, pid, req.Phone); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *AccountHandler) SetPrimaryPhone(c fiber.Ctx) error {
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

func (h *AccountHandler) DeletePhone(c fiber.Ctx) error {
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

func (h *AccountHandler) SendPasswordCode(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.SendPasswordCode
	_ = c.Bind().Body(&req)
	if err := h.svc.SendPasswordCode(c.Context(), aid, req.Lang); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}

func (h *AccountHandler) ChangePassword(c fiber.Ctx) error {
	aid, err := accountID(c)
	if err != nil {
		return wrap(c, err)
	}
	var req reqdto.ChangePassword
	if err := c.Bind().Body(&req); err != nil {
		return fiberx.ErrorFromErrx(c, sys.ErrInvalidBody)
	}
	if err := h.svc.ChangePassword(c.Context(), aid, req.CurrentPassword, req.NewPassword, req.VerificationCode); err != nil {
		return wrap(c, err)
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
}
