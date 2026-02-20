package handler

import (
	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	userOccupationAppCommands "nfxid/modules/directory/application/user_occupations/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserOccupationHandler struct {
	appSvc *userOccupationApp.Service
}

func NewUserOccupationHandler(appSvc *userOccupationApp.Service) *UserOccupationHandler {
	return &UserOccupationHandler{appSvc: appSvc}
}

func (h *UserOccupationHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserOccupationCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userOccupationID, err := h.appSvc.CreateUserOccupation(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user occupation
	userOccupationView, err := h.appSvc.GetUserOccupation(c.Context(), userOccupationID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User occupation created successfully", httpx.SuccessOptions{Data: respdto.UserOccupationROToDTO(&userOccupationView)})
}

func (h *UserOccupationHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.UserOccupationByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserOccupation(c.Context(), req.UserOccupationID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User occupation retrieved successfully", httpx.SuccessOptions{Data: respdto.UserOccupationROToDTO(&result)})
}

func (h *UserOccupationHandler) Update(c fiber.Ctx) error {
	var req reqdto.UserOccupationUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserOccupation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User occupation updated successfully")
}

func (h *UserOccupationHandler) Delete(c fiber.Ctx) error {
	var req reqdto.UserOccupationByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userOccupationAppCommands.DeleteUserOccupationCmd{UserOccupationID: req.UserOccupationID}
	if err := h.appSvc.DeleteUserOccupation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User occupation deleted successfully")
}

// GetByUserID 根据用户ID获取用户职业列表
func (h *UserOccupationHandler) GetByUserID(c fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	results, err := h.appSvc.GetUserOccupationsByUserID(c.Context(), req.UserID, nil)
	if err != nil {
		return err
	}

	dtos := respdto.UserOccupationListROToDTO(results)

	return fiberx.OK(c, "User occupations retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

// fiber:context-methods migrated
