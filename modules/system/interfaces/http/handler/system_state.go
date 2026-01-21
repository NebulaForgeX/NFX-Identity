package handler

import (
	systemStateApp "nfxid/modules/system/application/system_state"
	systemStateCommands "nfxid/modules/system/application/system_state/commands"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SystemStateHandler struct {
	appSvc *systemStateApp.Service
}

func NewSystemStateHandler(appSvc *systemStateApp.Service) *SystemStateHandler {
	return &SystemStateHandler{appSvc: appSvc}
}

// GetLatest 获取最新的系统状态
func (h *SystemStateHandler) GetLatest(c *fiber.Ctx) error {
	result, err := h.appSvc.GetLatestSystemState(c.Context())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get system state: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System state retrieved successfully", httpresp.SuccessOptions{Data: result})
}

// GetByID 根据 ID 获取系统状态
func (h *SystemStateHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid system state ID: "+err.Error())
	}

	result, err := h.appSvc.GetSystemState(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "System state not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System state retrieved successfully", httpresp.SuccessOptions{Data: result})
}

// Initialize 初始化系统
func (h *SystemStateHandler) Initialize(c *fiber.Ctx) error {
	var req struct {
		Version string `json:"version"`
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := systemStateCommands.InitializeSystemCmd{Version: req.Version}
	if err := h.appSvc.InitializeSystem(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to initialize system: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System initialized successfully")
}

// Reset 重置系统
func (h *SystemStateHandler) Reset(c *fiber.Ctx) error {
	var req struct {
		ResetBy uuid.UUID `json:"reset_by"`
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := systemStateCommands.ResetSystemCmd{ResetBy: req.ResetBy}
	if err := h.appSvc.ResetSystem(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to reset system: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System reset successfully")
}

// Delete 删除系统状态
func (h *SystemStateHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid system state ID: "+err.Error())
	}

	cmd := systemStateCommands.DeleteSystemStateCmd{SystemStateID: id}
	if err := h.appSvc.DeleteSystemState(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete system state: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System state deleted successfully")
}
