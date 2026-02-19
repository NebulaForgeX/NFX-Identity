package handler

import (
	bootstrapApp "nfxid/modules/system/application/bootstrap"
	systemStateApp "nfxid/modules/system/application/system_state"
	systemStateCommands "nfxid/modules/system/application/system_state/commands"
	"nfxid/modules/system/interfaces/http/dto/reqdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *SystemStateHandler) GetLatest(c fiber.Ctx) error {
	result, err := h.appSvc.GetLatestSystemState(c.Context())
	if err != nil {
		return err
	}
	return fiberx.OK(c, "System state retrieved successfully", httpx.SuccessOptions{Data: result})
}

// GetByID 根据 ID 获取系统状态
func (h *SystemStateHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.SystemStateByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetSystemState(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "System state retrieved successfully", httpx.SuccessOptions{Data: result})
}

// Initialize 初始化系统
// 使用 BootstrapInit 服务进行完整的系统初始化
func (h *SystemStateHandler) Initialize(c fiber.Ctx) error {
	var req reqdto.SystemStateInitializeRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	// 使用 DTO 的转换方法
	bootstrapCmd := req.ToBootstrapInitCmd()

	if err := h.bootstrapSvc.BootstrapInit(c.Context(), bootstrapCmd); err != nil {
		return err
	}

	return fiberx.OK(c, "System initialized successfully")
}

// Reset 重置系统
func (h *SystemStateHandler) Reset(c fiber.Ctx) error {
	var req reqdto.SystemStateResetRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToResetSystemCmd()
	if err := h.appSvc.ResetSystem(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "System reset successfully")
}

// Delete 删除系统状态
func (h *SystemStateHandler) Delete(c fiber.Ctx) error {
	var req reqdto.SystemStateByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := systemStateCommands.DeleteSystemStateCmd{SystemStateID: req.ID}
	if err := h.appSvc.DeleteSystemState(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "System state deleted successfully")
}

// fiber:context-methods migrated
