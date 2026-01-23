package handler

import (
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *PasswordHistoryHandler) Create(c *fiber.Ctx) error {
	var req reqdto.PasswordHistoryCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	passwordHistoryID, err := h.appSvc.CreatePasswordHistory(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create password history: "+err.Error())
	}

	// Get the created password history
	passwordHistoryView, err := h.appSvc.GetPasswordHistory(c.Context(), passwordHistoryID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created password history: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Password history created successfully", httpresp.SuccessOptions{Data: respdto.PasswordHistoryROToDTO(&passwordHistoryView)})
}

// GetByID 根据 ID 获取密码历史
func (h *PasswordHistoryHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.PasswordHistoryByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetPasswordHistory(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Password history not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Password history retrieved successfully", httpresp.SuccessOptions{Data: respdto.PasswordHistoryROToDTO(&result)})
}
