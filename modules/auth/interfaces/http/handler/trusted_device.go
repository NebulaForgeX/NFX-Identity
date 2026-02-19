package handler

import (
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	trustedDeviceAppCommands "nfxid/modules/auth/application/trusted_devices/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TrustedDeviceHandler struct {
	appSvc *trustedDeviceApp.Service
}

func NewTrustedDeviceHandler(appSvc *trustedDeviceApp.Service) *TrustedDeviceHandler {
	return &TrustedDeviceHandler{
		appSvc: appSvc,
	}
}

// Create 创建受信任设备
func (h *TrustedDeviceHandler) Create(c fiber.Ctx) error {
	var req reqdto.TrustedDeviceCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	trustedDeviceID, err := h.appSvc.CreateTrustedDevice(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created trusted device
	trustedDeviceView, err := h.appSvc.GetTrustedDevice(c.Context(), trustedDeviceID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Trusted device created successfully", httpx.SuccessOptions{Data: respdto.TrustedDeviceROToDTO(&trustedDeviceView)})
}

// GetByID 根据 ID 获取受信任设备
func (h *TrustedDeviceHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.TrustedDeviceByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetTrustedDevice(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Trusted device retrieved successfully", httpx.SuccessOptions{Data: respdto.TrustedDeviceROToDTO(&result)})
}

// Delete 删除受信任设备
func (h *TrustedDeviceHandler) Delete(c fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "id is required"))
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get the trusted device to get its deviceID
	trustedDevice, err := h.appSvc.GetTrustedDevice(c.Context(), id)
	if err != nil {
		return err
	}

	cmd := trustedDeviceAppCommands.DeleteTrustedDeviceCmd{
		DeviceID: trustedDevice.DeviceID,
	}
	if err := h.appSvc.DeleteTrustedDevice(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Trusted device deleted successfully")
}

// fiber:context-methods migrated
