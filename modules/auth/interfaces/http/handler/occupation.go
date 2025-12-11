package handler

import (
	occupationApp "nfxid/modules/auth/application/occupation"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type OccupationHandler struct {
	appSvc *occupationApp.Service
}

func NewOccupationHandler(appSvc *occupationApp.Service) *OccupationHandler {
	return &OccupationHandler{
		appSvc: appSvc,
	}
}

// Create 创建职业信息
func (h *OccupationHandler) Create(c *fiber.Ctx) error {
	var req reqdto.OccupationCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreateOccupation(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create occupation: "+err.Error())
	}

	// Get the created occupation view
	occupationView, err := h.appSvc.GetOccupation(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created occupation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Occupation created successfully", httpresp.SuccessOptions{Data: respdto.OccupationViewToDTO(&occupationView)})
}

// GetByID 根据 ID 获取职业信息
func (h *OccupationHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.OccupationByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetOccupation(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Occupation not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Occupation retrieved successfully", httpresp.SuccessOptions{Data: respdto.OccupationViewToDTO(&result)})
}

// GetByProfileID 根据 ProfileID 获取职业信息列表
func (h *OccupationHandler) GetByProfileID(c *fiber.Ctx) error {
	var req reqdto.OccupationByProfileIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetOccupationsByProfileID(c.Context(), req.ProfileID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get occupations: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Occupations retrieved successfully", httpresp.SuccessOptions{Data: respdto.OccupationListViewToDTO(results)})
}

// GetAll 获取职业信息列表
func (h *OccupationHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.OccupationQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetOccupationList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get occupations: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Occupations retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.OccupationListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新职业信息
func (h *OccupationHandler) Update(c *fiber.Ctx) error {
	var req reqdto.OccupationUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateOccupation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update occupation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Occupation updated successfully")
}

// Delete 删除职业信息
func (h *OccupationHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.OccupationByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := occupationApp.DeleteOccupationCmd{OccupationID: req.ID}
	if err := h.appSvc.DeleteOccupation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete occupation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Occupation deleted successfully")
}
