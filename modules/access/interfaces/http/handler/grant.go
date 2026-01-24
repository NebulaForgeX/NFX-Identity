package handler

import (
	grantApp "nfxid/modules/access/application/grants"
	grantAppCommands "nfxid/modules/access/application/grants/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type GrantHandler struct {
	appSvc *grantApp.Service
}

func NewGrantHandler(appSvc *grantApp.Service) *GrantHandler {
	return &GrantHandler{
		appSvc: appSvc,
	}
}

// Create 创建授权
func (h *GrantHandler) Create(c *fiber.Ctx) error {
	var req reqdto.GrantCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	grantID, err := h.appSvc.CreateGrant(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create grant: "+err.Error())
	}

	// Get the created grant
	grantView, err := h.appSvc.GetGrant(c.Context(), grantID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created grant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Grant created successfully", httpresp.SuccessOptions{Data: respdto.GrantROToDTO(&grantView)})
}

// GetByID 根据 ID 获取授权
func (h *GrantHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.GrantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetGrant(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Grant not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Grant retrieved successfully", httpresp.SuccessOptions{Data: respdto.GrantROToDTO(&result)})
}

// Update 更新授权
func (h *GrantHandler) Update(c *fiber.Ctx) error {
	var req reqdto.GrantUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateGrant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update grant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Grant updated successfully")
}

// Revoke 撤销授权
func (h *GrantHandler) Revoke(c *fiber.Ctx) error {
	var req reqdto.GrantRevokeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeGrant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke grant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Grant revoked successfully")
}

// Delete 删除授权
func (h *GrantHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.GrantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := grantAppCommands.DeleteGrantCmd{GrantID: req.ID}
	if err := h.appSvc.DeleteGrant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete grant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Grant deleted successfully")
}

// GetBySubject 根据主体获取授权列表
func (h *GrantHandler) GetBySubject(c *fiber.Ctx) error {
	var req reqdto.GrantBySubjectRequestDTO
	if err := c.QueryParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request query: "+err.Error())
	}

	// 使用 application 层的方法，避免 handler 直接依赖 domain 层
	grants, err := h.appSvc.GetGrantsBySubjectString(c.Context(), req.SubjectType, req.SubjectID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Failed to get grants: "+err.Error())
	}

	dtos := make([]*respdto.GrantDTO, len(grants))
	for i, g := range grants {
		dtos[i] = respdto.GrantROToDTO(&g)
	}

	return httpresp.Success(c, fiber.StatusOK, "Grants retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}