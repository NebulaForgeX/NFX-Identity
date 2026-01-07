package handler

import (
	userApp "nfxid/modules/auth/application/user"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/logx"
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
	result, err := h.appSvc.CreateUser(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user: "+err.Error())
	}

	// Get the created user view
	userView, err := h.appSvc.GetUser(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User created successfully", httpresp.SuccessOptions{Data: respdto.UserViewToDTO(&userView)})
}

// Login 登录（支持用户名、邮箱或手机号）
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req reqdto.UserLoginRequestDTO
	if err := c.BodyParser(&req); err != nil {
		logx.S().Errorf("❌ Failed to parse login request: %v", err)
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	logx.S().Infof("🔐 Login attempt for identifier: %s", req.Identifier)

	cmd := req.ToLoginCmd()
	result, err := h.appSvc.Login(c.Context(), cmd)
	if err != nil {
		logx.S().Errorf("❌ Login failed for identifier %s: %v", req.Identifier, err)
		return httpresp.Error(c, fiber.StatusUnauthorized, "Login failed: "+err.Error())
	}

	logx.S().Infof("✅ Login successful for identifier: %s", req.Identifier)
	return httpresp.Success(c, fiber.StatusOK, "Login successful", httpresp.SuccessOptions{Data: respdto.LoginResponseToDTO(result)})
}

// RefreshToken 刷新 Token
func (h *UserHandler) RefreshToken(c *fiber.Ctx) error {
	var req reqdto.UserRefreshTokenRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRefreshCmd()
	result, err := h.appSvc.RefreshToken(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusUnauthorized, "Failed to refresh token: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Token refreshed successfully", httpresp.SuccessOptions{Data: respdto.RefreshResponseToDTO(result)})
}

// SendVerificationCode 发送验证码（发送邮件，存储到 Redis）
func (h *UserHandler) SendVerificationCode(c *fiber.Ctx) error {
	var req reqdto.SendVerificationCodeRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	err := h.appSvc.SendVerificationCode(c.Context(), userApp.SendVerificationCodeCmd{
		Email:   req.Email,
		Purpose: "register",
	})
	if err != nil {
		logx.S().Errorf("failed to send verification code: %v", err)
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to send verification code: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Verification code sent successfully", httpresp.SuccessOptions{})
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

	return httpresp.Success(c, fiber.StatusOK, "User retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserViewToDTO(&result)})
}

// GetAll 获取用户列表
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.UserQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetUserList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get users: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Users retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.UserListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新用户
func (h *UserHandler) Update(c *fiber.Ctx) error {
	var req reqdto.UserUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUser(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User updated successfully")
}

// Delete 删除用户
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userApp.DeleteUserCmd{UserID: req.ID}
	if err := h.appSvc.DeleteUser(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User deleted successfully")
}

// DeleteAccount 删除账户（特殊逻辑）
func (h *UserHandler) DeleteAccount(c *fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	if err := h.appSvc.DeleteAccount(c.Context(), req.ID); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete account: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Account deleted successfully")
}
