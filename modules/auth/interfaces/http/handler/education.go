package handler

import (
	educationApp "nebulaid/modules/auth/application/education"
	"nebulaid/modules/auth/interfaces/http/dto/reqdto"
	"nebulaid/modules/auth/interfaces/http/dto/respdto"
	"nebulaid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type EducationHandler struct {
	appSvc *educationApp.Service
}

func NewEducationHandler(appSvc *educationApp.Service) *EducationHandler {
	return &EducationHandler{
		appSvc: appSvc,
	}
}

// Create 创建教育经历
func (h *EducationHandler) Create(c *fiber.Ctx) error {
	var req reqdto.EducationCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreateEducation(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create education: "+err.Error())
	}

	// Get the created education view
	educationView, err := h.appSvc.GetEducation(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created education: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Education created successfully", httpresp.SuccessOptions{Data: respdto.EducationViewToDTO(&educationView)})
}

// GetByID 根据 ID 获取教育经历
func (h *EducationHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.EducationByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetEducation(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Education not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Education retrieved successfully", httpresp.SuccessOptions{Data: respdto.EducationViewToDTO(&result)})
}

// GetByProfileID 根据 ProfileID 获取教育经历列表
func (h *EducationHandler) GetByProfileID(c *fiber.Ctx) error {
	var req reqdto.EducationByProfileIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetEducationsByProfileID(c.Context(), req.ProfileID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get educations: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Educations retrieved successfully", httpresp.SuccessOptions{Data: respdto.EducationListViewToDTO(results)})
}

// GetAll 获取教育经历列表
func (h *EducationHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.EducationQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetEducationList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get educations: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Educations retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.EducationListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新教育经历
func (h *EducationHandler) Update(c *fiber.Ctx) error {
	var req reqdto.EducationUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateEducation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update education: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Education updated successfully")
}

// Delete 删除教育经历
func (h *EducationHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.EducationByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := educationApp.DeleteEducationCmd{EducationID: req.ID}
	if err := h.appSvc.DeleteEducation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete education: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Education deleted successfully")
}
