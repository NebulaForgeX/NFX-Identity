package handler

import (
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	userCredentialAppCommands "nfxid/modules/auth/application/user_credentials/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserCredentialHandler struct {
	appSvc *userCredentialApp.Service
}

func NewUserCredentialHandler(appSvc *userCredentialApp.Service) *UserCredentialHandler {
	return &UserCredentialHandler{
		appSvc: appSvc,
	}
}

// Create 创建用户凭证
func (h *UserCredentialHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserCredentialCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userCredentialID, err := h.appSvc.CreateUserCredential(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user credential
	userCredentialView, err := h.appSvc.GetUserCredential(c.Context(), userCredentialID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User credential created successfully", httpx.SuccessOptions{Data: respdto.UserCredentialROToDTO(&userCredentialView)})
}

// GetByID 根据 ID 获取用户凭证
func (h *UserCredentialHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.UserCredentialByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserCredential(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User credential retrieved successfully", httpx.SuccessOptions{Data: respdto.UserCredentialROToDTO(&result)})
}

// Update 更新用户凭证
func (h *UserCredentialHandler) Update(c fiber.Ctx) error {
	var req reqdto.UserCredentialUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserCredential(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User credential updated successfully")
}

// Delete 删除用户凭证
func (h *UserCredentialHandler) Delete(c fiber.Ctx) error {
	var req reqdto.UserCredentialByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userCredentialAppCommands.DeleteUserCredentialCmd{UserCredentialID: req.ID}
	if err := h.appSvc.DeleteUserCredential(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User credential deleted successfully")
}

// fiber:context-methods migrated
