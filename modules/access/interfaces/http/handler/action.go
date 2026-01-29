package handler

import (
	actionApp "nfxid/modules/access/application/actions"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ActionHandler struct {
	appSvc *actionApp.Service
}

func NewActionHandler(appSvc *actionApp.Service) *ActionHandler {
	return &ActionHandler{appSvc: appSvc}
}

// GetByID 根据 ID 获取 Action
func (h *ActionHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ActionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	result, err := h.appSvc.GetAction(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Action not found: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusOK, "Action retrieved successfully", httpresp.SuccessOptions{Data: respdto.ActionROToDTO(&result)})
}

// GetByKey 根据 Key 获取 Action
func (h *ActionHandler) GetByKey(c *fiber.Ctx) error {
	var req reqdto.ActionByKeyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	result, err := h.appSvc.GetActionByKey(c.Context(), req.Key)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Action not found: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusOK, "Action retrieved successfully", httpresp.SuccessOptions{Data: respdto.ActionROToDTO(&result)})
}

// Create 创建 Action
func (h *ActionHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ActionCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}
	cmd := req.ToCreateCmd()
	actionID, err := h.appSvc.CreateAction(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create action: "+err.Error())
	}
	result, err := h.appSvc.GetAction(c.Context(), actionID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created action: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusCreated, "Action created successfully", httpresp.SuccessOptions{Data: respdto.ActionROToDTO(&result)})
}
