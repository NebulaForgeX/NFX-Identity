package handler

import (
	bootstrapApp "nfxid/modules/system/application/bootstrap"
	systemStateApp "nfxid/modules/system/application/system_state"
	systemStateCommands "nfxid/modules/system/application/system_state/commands"
	"nfxid/modules/system/interfaces/http/dto/reqdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type SystemStateHandler struct {
	appSvc       *systemStateApp.Service
	bootstrapSvc *bootstrapApp.Service
}

func NewSystemStateHandler(appSvc *systemStateApp.Service, bootstrapSvc *bootstrapApp.Service) *SystemStateHandler {
	return &SystemStateHandler{
		appSvc:       appSvc,
		bootstrapSvc: bootstrapSvc,
	}
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
	var req reqdto.SystemStateByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetSystemState(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "System state not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System state retrieved successfully", httpresp.SuccessOptions{Data: result})
}

// Initialize 初始化系统
// 使用 BootstrapInit 服务进行完整的系统初始化
func (h *SystemStateHandler) Initialize(c *fiber.Ctx) error {
	var req reqdto.SystemStateInitializeRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	// 使用 DTO 的转换方法
	bootstrapCmd := req.ToBootstrapInitCmd()

	if err := h.bootstrapSvc.BootstrapInit(c.Context(), bootstrapCmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to initialize system: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System initialized successfully")
}

// Reset 重置系统
func (h *SystemStateHandler) Reset(c *fiber.Ctx) error {
	var req reqdto.SystemStateResetRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToResetSystemCmd()
	if err := h.appSvc.ResetSystem(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to reset system: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System reset successfully")
}

// Delete 删除系统状态
func (h *SystemStateHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.SystemStateByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := systemStateCommands.DeleteSystemStateCmd{SystemStateID: req.ID}
	if err := h.appSvc.DeleteSystemState(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete system state: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "System state deleted successfully")
}
