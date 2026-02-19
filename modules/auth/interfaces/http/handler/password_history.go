package handler

import (
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type PasswordHistoryHandler struct {
	appSvc *passwordHistoryApp.Service
}

func NewPasswordHistoryHandler(appSvc *passwordHistoryApp.Service) *PasswordHistoryHandler {
	return &PasswordHistoryHandler{
		appSvc: appSvc,
	}
}

// Create 创建密码历史
func (h *PasswordHistoryHandler) Create(c fiber.Ctx) error {
	var req reqdto.PasswordHistoryCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	passwordHistoryID, err := h.appSvc.CreatePasswordHistory(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created password history
	passwordHistoryView, err := h.appSvc.GetPasswordHistory(c.Context(), passwordHistoryID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Password history created successfully", httpx.SuccessOptions{Data: respdto.PasswordHistoryROToDTO(&passwordHistoryView)})
}

// GetByID 根据 ID 获取密码历史
func (h *PasswordHistoryHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.PasswordHistoryByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetPasswordHistory(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Password history retrieved successfully", httpx.SuccessOptions{Data: respdto.PasswordHistoryROToDTO(&result)})
}

// fiber:context-methods migrated
