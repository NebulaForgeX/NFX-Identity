package handler

import (
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	mfaFactorAppCommands "nfxid/modules/auth/application/mfa_factors/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *MFAFactorHandler) Create(c fiber.Ctx) error {
	var req reqdto.MFAFactorCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	factorID, err := h.appSvc.CreateMFAFactor(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created MFA factor
	mfaFactorView, err := h.appSvc.GetMFAFactor(c.Context(), factorID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "MFA factor created successfully", httpx.SuccessOptions{Data: respdto.MFAFactorROToDTO(&mfaFactorView)})
}

// GetByID 根据 ID 获取 MFA 因子
func (h *MFAFactorHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.MFAFactorByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetMFAFactor(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "MFA factor retrieved successfully", httpx.SuccessOptions{Data: respdto.MFAFactorROToDTO(&result)})
}

// Update 更新 MFA 因子
func (h *MFAFactorHandler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "id is required"))
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get the MFA factor to get its factorID
	mfaFactor, err := h.appSvc.GetMFAFactor(c.Context(), id)
	if err != nil {
		return err
	}

	var req reqdto.MFAFactorUpdateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd(mfaFactor.FactorID)
	if err := h.appSvc.UpdateMFAFactor(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated MFA factor
	updatedMfaFactor, err := h.appSvc.GetMFAFactor(c.Context(), id)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "MFA factor updated successfully", httpx.SuccessOptions{Data: respdto.MFAFactorROToDTO(&updatedMfaFactor)})
}

// Delete 删除 MFA 因子
func (h *MFAFactorHandler) Delete(c fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "id is required"))
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get the MFA factor to get its factorID
	mfaFactor, err := h.appSvc.GetMFAFactor(c.Context(), id)
	if err != nil {
		return err
	}

	cmd := mfaFactorAppCommands.DeleteMFAFactorCmd{
		FactorID: mfaFactor.FactorID,
	}
	if err := h.appSvc.DeleteMFAFactor(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "MFA factor deleted successfully")
}

// fiber:context-methods migrated
