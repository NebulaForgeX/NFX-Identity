package handler

import (
	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	userOccupationAppCommands "nfxid/modules/directory/application/user_occupations/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserOccupationHandler struct {
	appSvc *userOccupationApp.Service
}

func NewUserOccupationHandler(appSvc *userOccupationApp.Service) *UserOccupationHandler {
	return &UserOccupationHandler{appSvc: appSvc}
}

func (h *UserOccupationHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserOccupationCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userOccupationID, err := h.appSvc.CreateUserOccupation(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user occupation: "+err.Error())
	}

	// Get the created user occupation
	userOccupationView, err := h.appSvc.GetUserOccupation(c.Context(), userOccupationID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user occupation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User occupation created successfully", httpresp.SuccessOptions{Data: respdto.UserOccupationROToDTO(&userOccupationView)})
}

func (h *UserOccupationHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserOccupation(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User occupation not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User occupation retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserOccupationROToDTO(&result)})
}

func (h *UserOccupationHandler) Update(c *fiber.Ctx) error {
	var req reqdto.UserOccupationUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserOccupation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user occupation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User occupation updated successfully")
}

func (h *UserOccupationHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userOccupationAppCommands.DeleteUserOccupationCmd{UserOccupationID: req.ID}
	if err := h.appSvc.DeleteUserOccupation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user occupation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User occupation deleted successfully")
}
