package handler

import (
	appApp "nfxid/modules/clients/application/apps"
	appAppCommands "nfxid/modules/clients/application/apps/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *AppHandler) Create(c fiber.Ctx) error {
	var req reqdto.AppCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	appID, err := h.appSvc.CreateApp(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created app
	appView, err := h.appSvc.GetApp(c.Context(), appID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "App created successfully", httpx.SuccessOptions{Data: respdto.AppROToDTO(&appView)})
}

// GetByID 根据 ID 获取应用
func (h *AppHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.AppByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetApp(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "App retrieved successfully", httpx.SuccessOptions{Data: respdto.AppROToDTO(&result)})
}

// GetByAppID 根据 AppID 获取应用
func (h *AppHandler) GetByAppID(c fiber.Ctx) error {
	var req reqdto.AppByAppIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetAppByAppID(c.Context(), req.AppID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "App retrieved successfully", httpx.SuccessOptions{Data: respdto.AppROToDTO(&result)})
}

// Update 更新应用
func (h *AppHandler) Update(c fiber.Ctx) error {
	var req reqdto.AppUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateApp(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "App updated successfully")
}

// Delete 删除应用
func (h *AppHandler) Delete(c fiber.Ctx) error {
	var req reqdto.AppByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := appAppCommands.DeleteAppCmd{AppID: req.ID}
	if err := h.appSvc.DeleteApp(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "App deleted successfully")
}

// fiber:context-methods migrated
