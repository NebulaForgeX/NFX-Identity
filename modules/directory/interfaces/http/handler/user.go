package handler

import (
	userApp "nfxid/modules/directory/application/users"
	userAppCommands "nfxid/modules/directory/application/users/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	appSvc *userApp.Service
}

func NewUserHandler(appSvc *userApp.Service) *UserHandler {
	return &UserHandler{
		appSvc: appSvc,
	}
}

// Create 创建用户
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userID, err := h.appSvc.CreateUser(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user: "+err.Error())
	}

	// Get the created user
	userView, err := h.appSvc.GetUser(c.Context(), userID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User created successfully", httpresp.SuccessOptions{Data: respdto.UserROToDTO(&userView)})
}

// GetByID 根据 ID 获取用户
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUser(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserROToDTO(&result)})
}

// GetByUsername 根据 Username 获取用户
func (h *UserHandler) GetByUsername(c *fiber.Ctx) error {
	var req reqdto.UserByUsernameRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserByUsername(c.Context(), req.Username)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserROToDTO(&result)})
}

// UpdateStatus 更新用户状态
func (h *UserHandler) UpdateStatus(c *fiber.Ctx) error {
	var req reqdto.UserUpdateStatusRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateStatusCmd()
	if err := h.appSvc.UpdateUserStatus(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user status: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User status updated successfully")
}

// UpdateUsername 更新用户名
func (h *UserHandler) UpdateUsername(c *fiber.Ctx) error {
	var req reqdto.UserUpdateUsernameRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateUsernameCmd()
	if err := h.appSvc.UpdateUsername(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update username: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Username updated successfully")
}

// Verify 验证用户
func (h *UserHandler) Verify(c *fiber.Ctx) error {
	var req reqdto.UserVerifyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := req.ToVerifyCmd()
	if err := h.appSvc.VerifyUser(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to verify user: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User verified successfully")
}

// Delete 删除用户（软删除）
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userAppCommands.DeleteUserCmd{UserID: req.ID}
	if err := h.appSvc.DeleteUser(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User deleted successfully")
}
