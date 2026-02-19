package handler

import (
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	loginAttemptAppCommands "nfxid/modules/auth/application/login_attempts/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type LoginAttemptHandler struct {
	appSvc *loginAttemptApp.Service
}

func NewLoginAttemptHandler(appSvc *loginAttemptApp.Service) *LoginAttemptHandler {
	return &LoginAttemptHandler{
		appSvc: appSvc,
	}
}

// Create 创建登录尝试
func (h *LoginAttemptHandler) Create(c fiber.Ctx) error {
	var req reqdto.LoginAttemptCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	loginAttemptID, err := h.appSvc.CreateLoginAttempt(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created login attempt
	loginAttemptView, err := h.appSvc.GetLoginAttempt(c.Context(), loginAttemptID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Login attempt created successfully", httpx.SuccessOptions{Data: respdto.LoginAttemptROToDTO(&loginAttemptView)})
}

// GetByID 根据 ID 获取登录尝试
func (h *LoginAttemptHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.LoginAttemptByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetLoginAttempt(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Login attempt retrieved successfully", httpx.SuccessOptions{Data: respdto.LoginAttemptROToDTO(&result)})
}

// Delete 删除登录尝试
func (h *LoginAttemptHandler) Delete(c fiber.Ctx) error {
	var req reqdto.LoginAttemptByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := loginAttemptAppCommands.DeleteLoginAttemptCmd{LoginAttemptID: req.ID}
	if err := h.appSvc.DeleteLoginAttempt(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Login attempt deleted successfully")
}

// fiber:context-methods migrated
