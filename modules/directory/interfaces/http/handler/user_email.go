package handler

import (
	userEmailApp "nfxid/modules/directory/application/user_emails"
	userEmailAppCommands "nfxid/modules/directory/application/user_emails/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserEmailHandler struct {
	appSvc *userEmailApp.Service
}

func NewUserEmailHandler(appSvc *userEmailApp.Service) *UserEmailHandler {
	return &UserEmailHandler{appSvc: appSvc}
}

func (h *UserEmailHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserEmailCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userEmailID, err := h.appSvc.CreateUserEmail(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user email
	userEmailView, err := h.appSvc.GetUserEmail(c.Context(), userEmailID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User email created successfully", httpx.SuccessOptions{Data: respdto.UserEmailROToDTO(&userEmailView)})
}

func (h *UserEmailHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserEmail(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User email retrieved successfully", httpx.SuccessOptions{Data: respdto.UserEmailROToDTO(&result)})
}

func (h *UserEmailHandler) Update(c fiber.Ctx) error {
	// UserEmail has specific update methods: SetPrimary and Verify
	// This generic Update method can be used for SetPrimary
	var req reqdto.UserEmailSetPrimaryRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := req.ToSetPrimaryCmd()
	if err := h.appSvc.SetPrimaryEmail(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Primary email set successfully")
}

func (h *UserEmailHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userEmailAppCommands.DeleteUserEmailCmd{UserEmailID: req.ID}
	if err := h.appSvc.DeleteUserEmail(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User email deleted successfully")
}

// SetPrimary 设置主邮箱
func (h *UserEmailHandler) SetPrimary(c fiber.Ctx) error {
	var req reqdto.UserEmailSetPrimaryRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := req.ToSetPrimaryCmd()
	if err := h.appSvc.SetPrimaryEmail(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Primary email set successfully")
}

// Verify 验证邮箱
func (h *UserEmailHandler) Verify(c fiber.Ctx) error {
	var req reqdto.UserEmailVerifyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := req.ToVerifyCmd()
	if err := h.appSvc.VerifyEmail(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Email verified successfully")
}

// GetByUserID 根据用户ID获取用户邮箱列表
func (h *UserEmailHandler) GetByUserID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	results, err := h.appSvc.GetUserEmailsByUserID(c.Context(), req.ID)
	if err != nil {
		return err
	}

	dtos := respdto.UserEmailListROToDTO(results)

	return fiberx.OK(c, "User emails retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

// fiber:context-methods migrated
