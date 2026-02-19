package handler

import (
	actionApp "nfxid/modules/access/application/actions"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ActionHandler struct {
	appSvc *actionApp.Service
}

func NewActionHandler(appSvc *actionApp.Service) *ActionHandler {
	return &ActionHandler{appSvc: appSvc}
}

// GetByID 根据 ID 获取 Action
func (h *ActionHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ActionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	result, err := h.appSvc.GetAction(c.Context(), req.ID)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Action retrieved successfully", httpx.SuccessOptions{Data: respdto.ActionROToDTO(&result)})
}

// GetByKey 根据 Key 获取 Action
func (h *ActionHandler) GetByKey(c fiber.Ctx) error {
	var req reqdto.ActionByKeyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	result, err := h.appSvc.GetActionByKey(c.Context(), req.Key)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Action retrieved successfully", httpx.SuccessOptions{Data: respdto.ActionROToDTO(&result)})
}

// Create 创建 Action
func (h *ActionHandler) Create(c fiber.Ctx) error {
	var req reqdto.ActionCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	cmd := req.ToCreateCmd()
	actionID, err := h.appSvc.CreateAction(c.Context(), cmd)
	if err != nil {
		return err
	}
	result, err := h.appSvc.GetAction(c.Context(), actionID)
	if err != nil {
		return err
	}
	return fiberx.Created(c, "Action created successfully", httpx.SuccessOptions{Data: respdto.ActionROToDTO(&result)})
}

// fiber:context-methods migrated
