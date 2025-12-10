package handler

import (
	roleApp "nebulaid/modules/auth/application/role"
	"nebulaid/modules/auth/interfaces/http/dto/reqdto"
	"nebulaid/modules/auth/interfaces/http/dto/respdto"
	"nebulaid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var req reqdto.RoleCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreateRole(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create role: "+err.Error())
	}

	// Get the created role view
	roleView, err := h.appSvc.GetRole(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Role created successfully", httpresp.SuccessOptions{Data: respdto.RoleViewToDTO(&roleView)})
}

// GetByID 根据 ID 获取角色
func (h *RoleHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.RoleByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRole(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Role not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role retrieved successfully", httpresp.SuccessOptions{Data: respdto.RoleViewToDTO(&result)})
}

// GetByName 根据名称获取角色
func (h *RoleHandler) GetByName(c *fiber.Ctx) error {
	var req reqdto.RoleByNameRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRoleByName(c.Context(), req.Name)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Role not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role retrieved successfully", httpresp.SuccessOptions{Data: respdto.RoleViewToDTO(&result)})
}

// GetAll 获取角色列表
func (h *RoleHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.RoleQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetRoleList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get roles: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Roles retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.RoleListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新角色
func (h *RoleHandler) Update(c *fiber.Ctx) error {
	var req reqdto.RoleUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role updated successfully")
}

// Delete 删除角色
func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.RoleByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := roleApp.DeleteRoleCmd{RoleID: req.ID}
	if err := h.appSvc.DeleteRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role deleted successfully")
}
