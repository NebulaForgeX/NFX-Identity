package handler

import (
	sessionApp "nfxid/modules/auth/application/sessions"
	sessionAppCommands "nfxid/modules/auth/application/sessions/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type SessionHandler struct {
	appSvc *sessionApp.Service
}

func NewSessionHandler(appSvc *sessionApp.Service) *SessionHandler {
	return &SessionHandler{
		appSvc: appSvc,
	}
}

// Create 创建会话
func (h *SessionHandler) Create(c *fiber.Ctx) error {
	var req reqdto.SessionCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	sessionID, err := h.appSvc.CreateSession(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create session: "+err.Error())
	}

	// Get the created session
	sessionView, err := h.appSvc.GetSession(c.Context(), sessionID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created session: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Session created successfully", httpresp.SuccessOptions{Data: respdto.SessionROToDTO(&sessionView)})
}

// GetByID 根据 ID 获取会话
func (h *SessionHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.SessionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetSession(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Session not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Session retrieved successfully", httpresp.SuccessOptions{Data: respdto.SessionROToDTO(&result)})
}

// Revoke 撤销会话
func (h *SessionHandler) Revoke(c *fiber.Ctx) error {
	var req reqdto.SessionRevokeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeSession(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke session: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Session revoked successfully")
}

// Delete 删除会话
func (h *SessionHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.SessionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := sessionAppCommands.DeleteSessionCmd{SessionID: req.ID}
	if err := h.appSvc.DeleteSession(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete session: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Session deleted successfully")
}
