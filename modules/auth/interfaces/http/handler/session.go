package handler

import (
	sessionApp "nfxid/modules/auth/application/sessions"
	sessionAppCommands "nfxid/modules/auth/application/sessions/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *SessionHandler) Create(c fiber.Ctx) error {
	var req reqdto.SessionCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	sessionID, err := h.appSvc.CreateSession(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created session
	sessionView, err := h.appSvc.GetSession(c.Context(), sessionID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Session created successfully", httpx.SuccessOptions{Data: respdto.SessionROToDTO(&sessionView)})
}

// GetByID 根据 ID 获取会话
func (h *SessionHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.SessionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetSession(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Session retrieved successfully", httpx.SuccessOptions{Data: respdto.SessionROToDTO(&result)})
}

// Revoke 撤销会话
func (h *SessionHandler) Revoke(c fiber.Ctx) error {
	var req reqdto.SessionRevokeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeSession(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Session revoked successfully")
}

// Delete 删除会话
func (h *SessionHandler) Delete(c fiber.Ctx) error {
	var req reqdto.SessionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := sessionAppCommands.DeleteSessionCmd{SessionID: req.ID}
	if err := h.appSvc.DeleteSession(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Session deleted successfully")
}

// fiber:context-methods migrated
