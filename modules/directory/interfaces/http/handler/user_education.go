package handler

import (
	userEducationApp "nfxid/modules/directory/application/user_educations"
	userEducationAppCommands "nfxid/modules/directory/application/user_educations/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserEducationHandler struct {
	appSvc *userEducationApp.Service
}

func NewUserEducationHandler(appSvc *userEducationApp.Service) *UserEducationHandler {
	return &UserEducationHandler{appSvc: appSvc}
}

func (h *UserEducationHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserEducationCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userEducationID, err := h.appSvc.CreateUserEducation(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user education
	userEducationView, err := h.appSvc.GetUserEducation(c.Context(), userEducationID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User education created successfully", httpx.SuccessOptions{Data: respdto.UserEducationROToDTO(&userEducationView)})
}

func (h *UserEducationHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.UserEducationByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserEducation(c.Context(), req.UserEducationID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User education retrieved successfully", httpx.SuccessOptions{Data: respdto.UserEducationROToDTO(&result)})
}

func (h *UserEducationHandler) Update(c fiber.Ctx) error {
	var req reqdto.UserEducationUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserEducation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User education updated successfully")
}

func (h *UserEducationHandler) Delete(c fiber.Ctx) error {
	var req reqdto.UserEducationByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userEducationAppCommands.DeleteUserEducationCmd{UserEducationID: req.UserEducationID}
	if err := h.appSvc.DeleteUserEducation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User education deleted successfully")
}

// GetByUserID 根据用户ID获取用户教育列表
func (h *UserEducationHandler) GetByUserID(c fiber.Ctx) error {
	var req reqdto.UserByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	results, err := h.appSvc.GetUserEducationsByUserID(c.Context(), req.UserID)
	if err != nil {
		return err
	}

	dtos := respdto.UserEducationListROToDTO(results)

	return fiberx.OK(c, "User educations retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

// fiber:context-methods migrated
