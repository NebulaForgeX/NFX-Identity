package handler

import (
	appApp "nfxid/modules/clients/application/apps"
	appAppCommands "nfxid/modules/clients/application/apps/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type AppHandler struct {
	appSvc *appApp.Service
}

func NewAppHandler(appSvc *appApp.Service) *AppHandler {
	return &AppHandler{
		appSvc: appSvc,
	}
}

// Create 创建应用
func (h *AppHandler) Create(c *fiber.Ctx) error {
	var req reqdto.AppCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	appID, err := h.appSvc.CreateApp(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create app: "+err.Error())
	}

	// Get the created app
	appView, err := h.appSvc.GetApp(c.Context(), appID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created app: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "App created successfully", httpresp.SuccessOptions{Data: respdto.AppROToDTO(&appView)})
}

// GetByID 根据 ID 获取应用
func (h *AppHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.AppByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetApp(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "App not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "App retrieved successfully", httpresp.SuccessOptions{Data: respdto.AppROToDTO(&result)})
}

// GetByAppID 根据 AppID 获取应用
func (h *AppHandler) GetByAppID(c *fiber.Ctx) error {
	var req reqdto.AppByAppIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetAppByAppID(c.Context(), req.AppID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "App not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "App retrieved successfully", httpresp.SuccessOptions{Data: respdto.AppROToDTO(&result)})
}

// Update 更新应用
func (h *AppHandler) Update(c *fiber.Ctx) error {
	var req reqdto.AppUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateApp(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update app: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "App updated successfully")
}

// Delete 删除应用
func (h *AppHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.AppByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := appAppCommands.DeleteAppCmd{AppID: req.ID}
	if err := h.appSvc.DeleteApp(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete app: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "App deleted successfully")
}
