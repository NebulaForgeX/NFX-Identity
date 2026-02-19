package handler

import (
	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPhoneAppCommands "nfxid/modules/directory/application/user_phones/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserPhoneHandler struct {
	appSvc *userPhoneApp.Service
}

func NewUserPhoneHandler(appSvc *userPhoneApp.Service) *UserPhoneHandler {
	return &UserPhoneHandler{appSvc: appSvc}
}

func (h *UserPhoneHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserPhoneCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userPhoneID, err := h.appSvc.CreateUserPhone(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user phone
	userPhoneView, err := h.appSvc.GetUserPhone(c.Context(), userPhoneID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User phone created successfully", httpx.SuccessOptions{Data: respdto.UserPhoneROToDTO(&userPhoneView)})
}

func (h *UserPhoneHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserPhone(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User phone retrieved successfully", httpx.SuccessOptions{Data: respdto.UserPhoneROToDTO(&result)})
}

func (h *UserPhoneHandler) Update(c fiber.Ctx) error {
	// Update can be SetPrimary, Verify, or UpdateVerificationCode
	// Try SetPrimary first
	var primaryReq reqdto.UserPhoneSetPrimaryRequestDTO
	if err := c.Bind().URI(&primaryReq); err == nil {
		if err := c.Bind().Body(&primaryReq); err == nil {
			cmd := primaryReq.ToSetPrimaryCmd()
			if err := h.appSvc.SetPrimaryPhone(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Primary phone set successfully")
		}
	}

	// Try Verify
	var verifyReq reqdto.UserPhoneVerifyRequestDTO
	if err := c.Bind().URI(&verifyReq); err == nil {
		if err := c.Bind().Body(&verifyReq); err == nil {
			cmd := verifyReq.ToVerifyCmd()
			if err := h.appSvc.VerifyPhone(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Phone verified successfully")
		}
	}

	// Try UpdateVerificationCode
	var updateCodeReq reqdto.UserPhoneUpdateVerificationCodeRequestDTO
	if err := c.Bind().URI(&updateCodeReq); err == nil {
		if err := c.Bind().Body(&updateCodeReq); err == nil {
			cmd := updateCodeReq.ToUpdateVerificationCodeCmd()
			if err := h.appSvc.UpdateVerificationCode(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Verification code updated successfully")
		}
	}

	return errx.ErrInvalidParams
}

// SetPrimary 设置主手机号
func (h *UserPhoneHandler) SetPrimary(c fiber.Ctx) error {
	var req reqdto.UserPhoneSetPrimaryRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToSetPrimaryCmd()
	if err := h.appSvc.SetPrimaryPhone(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Primary phone set successfully")
}

// Verify 验证手机号
func (h *UserPhoneHandler) Verify(c fiber.Ctx) error {
	var req reqdto.UserPhoneVerifyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToVerifyCmd()
	if err := h.appSvc.VerifyPhone(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Phone verified successfully")
}

func (h *UserPhoneHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userPhoneAppCommands.DeleteUserPhoneCmd{UserPhoneID: req.ID}
	if err := h.appSvc.DeleteUserPhone(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User phone deleted successfully")
}

// GetByUserID 根据用户ID获取用户电话列表
func (h *UserPhoneHandler) GetByUserID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	results, err := h.appSvc.GetUserPhonesByUserID(c.Context(), req.ID)
	if err != nil {
		return err
	}

	dtos := respdto.UserPhoneListROToDTO(results)

	return fiberx.OK(c, "User phones retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

// fiber:context-methods migrated
