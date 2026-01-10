package handler

import (
	userEmailApp "nfxid/modules/directory/application/user_emails"
	userEmailAppCommands "nfxid/modules/directory/application/user_emails/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserEmailHandler struct {
	appSvc *userEmailApp.Service
}

func NewUserEmailHandler(appSvc *userEmailApp.Service) *UserEmailHandler {
	return &UserEmailHandler{appSvc: appSvc}
}

func (h *UserEmailHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserEmailCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userEmailID, err := h.appSvc.CreateUserEmail(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user email: "+err.Error())
	}

	// Get the created user email
	userEmailView, err := h.appSvc.GetUserEmail(c.Context(), userEmailID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user email: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User email created successfully", httpresp.SuccessOptions{Data: respdto.UserEmailROToDTO(&userEmailView)})
}

func (h *UserEmailHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserEmail(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User email not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User email retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserEmailROToDTO(&result)})
}

func (h *UserEmailHandler) Update(c *fiber.Ctx) error {
	// UserEmail has specific update methods: SetPrimary and Verify
	// This generic Update method can be used for SetPrimary
	var req reqdto.UserEmailSetPrimaryRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := req.ToSetPrimaryCmd()
	if err := h.appSvc.SetPrimaryEmail(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to set primary email: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Primary email set successfully")
}

func (h *UserEmailHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userEmailAppCommands.DeleteUserEmailCmd{UserEmailID: req.ID}
	if err := h.appSvc.DeleteUserEmail(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user email: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User email deleted successfully")
}

// SetPrimary 设置主邮箱
func (h *UserEmailHandler) SetPrimary(c *fiber.Ctx) error {
	var req reqdto.UserEmailSetPrimaryRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := req.ToSetPrimaryCmd()
	if err := h.appSvc.SetPrimaryEmail(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to set primary email: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Primary email set successfully")
}

// Verify 验证邮箱
func (h *UserEmailHandler) Verify(c *fiber.Ctx) error {
	var req reqdto.UserEmailVerifyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := req.ToVerifyCmd()
	if err := h.appSvc.VerifyEmail(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to verify email: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Email verified successfully")
}
