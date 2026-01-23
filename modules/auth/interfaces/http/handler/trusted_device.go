package handler

import (
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	trustedDeviceAppCommands "nfxid/modules/auth/application/trusted_devices/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *TrustedDeviceHandler) Create(c *fiber.Ctx) error {
	var req reqdto.TrustedDeviceCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	trustedDeviceID, err := h.appSvc.CreateTrustedDevice(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create trusted device: "+err.Error())
	}

	// Get the created trusted device
	trustedDeviceView, err := h.appSvc.GetTrustedDevice(c.Context(), trustedDeviceID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created trusted device: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Trusted device created successfully", httpresp.SuccessOptions{Data: respdto.TrustedDeviceROToDTO(&trustedDeviceView)})
}

// GetByID 根据 ID 获取受信任设备
func (h *TrustedDeviceHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.TrustedDeviceByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetTrustedDevice(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Trusted device not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Trusted device retrieved successfully", httpresp.SuccessOptions{Data: respdto.TrustedDeviceROToDTO(&result)})
}

// Delete 删除受信任设备
func (h *TrustedDeviceHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "id is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid id: "+err.Error())
	}

	// Get the trusted device to get its deviceID
	trustedDevice, err := h.appSvc.GetTrustedDevice(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Trusted device not found: "+err.Error())
	}

	cmd := trustedDeviceAppCommands.DeleteTrustedDeviceCmd{
		DeviceID: trustedDevice.DeviceID,
	}
	if err := h.appSvc.DeleteTrustedDevice(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete trusted device: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Trusted device deleted successfully")
}
