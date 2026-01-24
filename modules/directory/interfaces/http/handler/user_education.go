package handler

import (
	userEducationApp "nfxid/modules/directory/application/user_educations"
	userEducationAppCommands "nfxid/modules/directory/application/user_educations/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserEducationHandler struct {
	appSvc *userEducationApp.Service
}

func NewUserEducationHandler(appSvc *userEducationApp.Service) *UserEducationHandler {
	return &UserEducationHandler{appSvc: appSvc}
}

func (h *UserEducationHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserEducationCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userEducationID, err := h.appSvc.CreateUserEducation(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user education: "+err.Error())
	}

	// Get the created user education
	userEducationView, err := h.appSvc.GetUserEducation(c.Context(), userEducationID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user education: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User education created successfully", httpresp.SuccessOptions{Data: respdto.UserEducationROToDTO(&userEducationView)})
}

func (h *UserEducationHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserEducation(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User education not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User education retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserEducationROToDTO(&result)})
}

func (h *UserEducationHandler) Update(c *fiber.Ctx) error {
	var req reqdto.UserEducationUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserEducation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user education: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User education updated successfully")
}

func (h *UserEducationHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userEducationAppCommands.DeleteUserEducationCmd{UserEducationID: req.ID}
	if err := h.appSvc.DeleteUserEducation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user education: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User education deleted successfully")
}

// GetByUserID 根据用户ID获取用户教育列表
func (h *UserEducationHandler) GetByUserID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetUserEducationsByUserID(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get user educations: "+err.Error())
	}

	dtos := respdto.UserEducationListROToDTO(results)

	return httpresp.Success(c, fiber.StatusOK, "User educations retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}
