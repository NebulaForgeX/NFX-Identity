package handler

import (
	userApp "nfxid/modules/directory/application/users"
	userAppCommands "nfxid/modules/directory/application/users/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *UserHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userID, err := h.appSvc.CreateUser(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user
	userView, err := h.appSvc.GetUser(c.Context(), userID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User created successfully", httpx.SuccessOptions{Data: respdto.UserROToDTO(&userView)})
}

// GetByID 根据 ID 获取用户
func (h *UserHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUser(c.Context(), req.UserID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User retrieved successfully", httpx.SuccessOptions{Data: respdto.UserROToDTO(&result)})
}

// GetByUsername 根据 Username 获取用户
func (h *UserHandler) GetByUsername(c fiber.Ctx) error {
	var req reqdto.UserByUsernameRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserByUsername(c.Context(), req.Username)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User retrieved successfully", httpx.SuccessOptions{Data: respdto.UserROToDTO(&result)})
}

// UpdateStatus 更新用户状态
func (h *UserHandler) UpdateStatus(c fiber.Ctx) error {
	var req reqdto.UserUpdateStatusRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateStatusCmd()
	if err := h.appSvc.UpdateUserStatus(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User status updated successfully")
}

// UpdateUsername 更新用户名
func (h *UserHandler) UpdateUsername(c fiber.Ctx) error {
	var req reqdto.UserUpdateUsernameRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateUsernameCmd()
	if err := h.appSvc.UpdateUsername(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Username updated successfully")
}

// Verify 验证用户
func (h *UserHandler) Verify(c fiber.Ctx) error {
	var req reqdto.UserVerifyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := req.ToVerifyCmd()
	if err := h.appSvc.VerifyUser(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User verified successfully")
}

// Delete 删除用户（软删除）
func (h *UserHandler) Delete(c fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userAppCommands.DeleteUserCmd{UserID: req.UserID}
	if err := h.appSvc.DeleteUser(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User deleted successfully")
}

// fiber:context-methods migrated
