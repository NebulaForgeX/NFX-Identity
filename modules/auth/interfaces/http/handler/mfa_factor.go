package handler

import (
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	mfaFactorAppCommands "nfxid/modules/auth/application/mfa_factors/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MFAFactorHandler struct {
	appSvc *mfaFactorApp.Service
}

func NewMFAFactorHandler(appSvc *mfaFactorApp.Service) *MFAFactorHandler {
	return &MFAFactorHandler{
		appSvc: appSvc,
	}
}

// Create 创建 MFA 因子
func (h *MFAFactorHandler) Create(c *fiber.Ctx) error {
	var req reqdto.MFAFactorCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	factorID, err := h.appSvc.CreateMFAFactor(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create MFA factor: "+err.Error())
	}

	// Get the created MFA factor
	mfaFactorView, err := h.appSvc.GetMFAFactor(c.Context(), factorID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created MFA factor: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "MFA factor created successfully", httpresp.SuccessOptions{Data: respdto.MFAFactorROToDTO(&mfaFactorView)})
}

// GetByID 根据 ID 获取 MFA 因子
func (h *MFAFactorHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.MFAFactorByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetMFAFactor(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "MFA factor not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "MFA factor retrieved successfully", httpresp.SuccessOptions{Data: respdto.MFAFactorROToDTO(&result)})
}

// Update 更新 MFA 因子
func (h *MFAFactorHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "id is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid id: "+err.Error())
	}

	// Get the MFA factor to get its factorID
	mfaFactor, err := h.appSvc.GetMFAFactor(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "MFA factor not found: "+err.Error())
	}

	var req reqdto.MFAFactorUpdateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd(mfaFactor.FactorID)
	if err := h.appSvc.UpdateMFAFactor(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update MFA factor: "+err.Error())
	}

	// Get the updated MFA factor
	updatedMfaFactor, err := h.appSvc.GetMFAFactor(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated MFA factor: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "MFA factor updated successfully", httpresp.SuccessOptions{Data: respdto.MFAFactorROToDTO(&updatedMfaFactor)})
}

// Delete 删除 MFA 因子
func (h *MFAFactorHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "id is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid id: "+err.Error())
	}

	// Get the MFA factor to get its factorID
	mfaFactor, err := h.appSvc.GetMFAFactor(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "MFA factor not found: "+err.Error())
	}

	cmd := mfaFactorAppCommands.DeleteMFAFactorCmd{
		FactorID: mfaFactor.FactorID,
	}
	if err := h.appSvc.DeleteMFAFactor(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete MFA factor: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "MFA factor deleted successfully")
}
