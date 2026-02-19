package handler

import (
	roleApp "nfxid/modules/access/application/roles"
	roleAppCommands "nfxid/modules/access/application/roles/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type RoleHandler struct {
	appSvc *roleApp.Service
}

func NewRoleHandler(appSvc *roleApp.Service) *RoleHandler {
	return &RoleHandler{
		appSvc: appSvc,
	}
}

// Create 创建角色
func (h *RoleHandler) Create(c fiber.Ctx) error {
	var req reqdto.RoleCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	roleID, err := h.appSvc.CreateRole(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created role
	roleView, err := h.appSvc.GetRole(c.Context(), roleID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Role created successfully", httpx.SuccessOptions{Data: respdto.RoleROToDTO(&roleView)})
}

// GetByID 根据 ID 获取角色
func (h *RoleHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.RoleByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetRole(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Role retrieved successfully", httpx.SuccessOptions{Data: respdto.RoleROToDTO(&result)})
}

// GetByKey 根据 Key 获取角色
func (h *RoleHandler) GetByKey(c fiber.Ctx) error {
	var req reqdto.RoleByKeyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetRoleByKey(c.Context(), req.Key)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Role retrieved successfully", httpx.SuccessOptions{Data: respdto.RoleROToDTO(&result)})
}

// Update 更新角色
func (h *RoleHandler) Update(c fiber.Ctx) error {
	var req reqdto.RoleUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateRole(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Role updated successfully")
}

// Delete 删除角色（软删除）
func (h *RoleHandler) Delete(c fiber.Ctx) error {
	var req reqdto.RoleByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := roleAppCommands.DeleteRoleCmd{RoleID: req.ID}
	if err := h.appSvc.DeleteRole(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Role deleted successfully")
}

// fiber:context-methods migrated
