package handler

import (
	arApp "nfxid/modules/access/application/action_requirements"
	arAppCommands "nfxid/modules/access/application/action_requirements/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ActionRequirementHandler struct {
	appSvc *arApp.Service
}

func NewActionRequirementHandler(appSvc *arApp.Service) *ActionRequirementHandler {
	return &ActionRequirementHandler{appSvc: appSvc}
}

// GetByID 根据 ID 获取 ActionRequirement
func (h *ActionRequirementHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ActionRequirementByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	result, err := h.appSvc.GetActionRequirement(c.Context(), req.ID)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Action requirement retrieved successfully", httpx.SuccessOptions{Data: respdto.ActionRequirementROToDTO(&result)})
}

// GetByPermissionID 根据 PermissionID 获取 ActionRequirement 列表
func (h *ActionRequirementHandler) GetByPermissionID(c fiber.Ctx) error {
	var req reqdto.ActionRequirementByPermissionIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	list, err := h.appSvc.GetByPermissionID(c.Context(), req.PermissionID)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Action requirements retrieved successfully", httpx.SuccessOptions{Data: respdto.ActionRequirementListROToDTO(list)})
}

// Create 创建 ActionRequirement
func (h *ActionRequirementHandler) Create(c fiber.Ctx) error {
	var req reqdto.ActionRequirementCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	cmd, err := req.ToCreateCmd()
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	arID, err := h.appSvc.CreateActionRequirement(c.Context(), cmd)
	if err != nil {
		return err
	}
	result, err := h.appSvc.GetActionRequirement(c.Context(), arID)
	if err != nil {
		return err
	}
	return fiberx.Created(c, "Action requirement created successfully", httpx.SuccessOptions{Data: respdto.ActionRequirementROToDTO(&result)})
}

// Delete 删除 ActionRequirement
func (h *ActionRequirementHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ActionRequirementDeleteRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	err := h.appSvc.DeleteActionRequirement(c.Context(), arAppCommands.DeleteActionRequirementCmd{ActionRequirementID: req.ID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Action requirement deleted successfully", httpx.SuccessOptions{})
}

// fiber:context-methods migrated
