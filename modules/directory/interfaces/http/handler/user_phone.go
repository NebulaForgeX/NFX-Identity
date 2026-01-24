package handler

import (
	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPhoneAppCommands "nfxid/modules/directory/application/user_phones/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserPhoneHandler struct {
	appSvc *userPhoneApp.Service
}

func NewUserPhoneHandler(appSvc *userPhoneApp.Service) *UserPhoneHandler {
	return &UserPhoneHandler{appSvc: appSvc}
}

func (h *UserPhoneHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserPhoneCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userPhoneID, err := h.appSvc.CreateUserPhone(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user phone: "+err.Error())
	}

	// Get the created user phone
	userPhoneView, err := h.appSvc.GetUserPhone(c.Context(), userPhoneID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user phone: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User phone created successfully", httpresp.SuccessOptions{Data: respdto.UserPhoneROToDTO(&userPhoneView)})
}

func (h *UserPhoneHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserPhone(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User phone not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User phone retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserPhoneROToDTO(&result)})
}

func (h *UserPhoneHandler) Update(c *fiber.Ctx) error {
	// Update can be SetPrimary, Verify, or UpdateVerificationCode
	// Try SetPrimary first
	var primaryReq reqdto.UserPhoneSetPrimaryRequestDTO
	if err := c.ParamsParser(&primaryReq); err == nil {
		if err := c.BodyParser(&primaryReq); err == nil {
			cmd := primaryReq.ToSetPrimaryCmd()
			if err := h.appSvc.SetPrimaryPhone(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to set primary phone: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Primary phone set successfully")
		}
	}

	// Try Verify
	var verifyReq reqdto.UserPhoneVerifyRequestDTO
	if err := c.ParamsParser(&verifyReq); err == nil {
		if err := c.BodyParser(&verifyReq); err == nil {
			cmd := verifyReq.ToVerifyCmd()
			if err := h.appSvc.VerifyPhone(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to verify phone: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Phone verified successfully")
		}
	}

	// Try UpdateVerificationCode
	var updateCodeReq reqdto.UserPhoneUpdateVerificationCodeRequestDTO
	if err := c.ParamsParser(&updateCodeReq); err == nil {
		if err := c.BodyParser(&updateCodeReq); err == nil {
			cmd := updateCodeReq.ToUpdateVerificationCodeCmd()
			if err := h.appSvc.UpdateVerificationCode(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update verification code: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Verification code updated successfully")
		}
	}

	return httpresp.Error(c, fiber.StatusBadRequest, "Invalid update request")
}

// SetPrimary 设置主手机号
func (h *UserPhoneHandler) SetPrimary(c *fiber.Ctx) error {
	var req reqdto.UserPhoneSetPrimaryRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToSetPrimaryCmd()
	if err := h.appSvc.SetPrimaryPhone(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to set primary phone: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Primary phone set successfully")
}

// Verify 验证手机号
func (h *UserPhoneHandler) Verify(c *fiber.Ctx) error {
	var req reqdto.UserPhoneVerifyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToVerifyCmd()
	if err := h.appSvc.VerifyPhone(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to verify phone: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Phone verified successfully")
}

func (h *UserPhoneHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userPhoneAppCommands.DeleteUserPhoneCmd{UserPhoneID: req.ID}
	if err := h.appSvc.DeleteUserPhone(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user phone: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User phone deleted successfully")
}

// GetByUserID 根据用户ID获取用户电话列表
func (h *UserPhoneHandler) GetByUserID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetUserPhonesByUserID(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get user phones: "+err.Error())
	}

	dtos := respdto.UserPhoneListROToDTO(results)

	return httpresp.Success(c, fiber.StatusOK, "User phones retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}
