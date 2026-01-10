package handler

import (
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	userCredentialAppCommands "nfxid/modules/auth/application/user_credentials/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *UserCredentialHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserCredentialCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userCredentialID, err := h.appSvc.CreateUserCredential(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user credential: "+err.Error())
	}

	// Get the created user credential
	userCredentialView, err := h.appSvc.GetUserCredential(c.Context(), userCredentialID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user credential: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User credential created successfully", httpresp.SuccessOptions{Data: respdto.UserCredentialROToDTO(&userCredentialView)})
}

// GetByID 根据 ID 获取用户凭证
func (h *UserCredentialHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.UserCredentialByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserCredential(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User credential not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User credential retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserCredentialROToDTO(&result)})
}

// Update 更新用户凭证
func (h *UserCredentialHandler) Update(c *fiber.Ctx) error {
	var req reqdto.UserCredentialUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserCredential(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user credential: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User credential updated successfully")
}

// Delete 删除用户凭证
func (h *UserCredentialHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.UserCredentialByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userCredentialAppCommands.DeleteUserCredentialCmd{UserCredentialID: req.ID}
	if err := h.appSvc.DeleteUserCredential(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user credential: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User credential deleted successfully")
}
