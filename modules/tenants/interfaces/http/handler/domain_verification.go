package handler

import (
	domainVerificationApp "nfxid/modules/tenants/application/domain_verifications"
	domainVerificationAppCommands "nfxid/modules/tenants/application/domain_verifications/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type DomainVerificationHandler struct {
	appSvc *domainVerificationApp.Service
}

func NewDomainVerificationHandler(appSvc *domainVerificationApp.Service) *DomainVerificationHandler {
	return &DomainVerificationHandler{appSvc: appSvc}
}

func (h *DomainVerificationHandler) Create(c *fiber.Ctx) error {
	var req reqdto.DomainVerificationCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	domainVerificationID, err := h.appSvc.CreateDomainVerification(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create domain verification: "+err.Error())
	}

	// Get the created domain verification
	domainVerificationView, err := h.appSvc.GetDomainVerification(c.Context(), domainVerificationID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created domain verification: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Domain verification created successfully", httpresp.SuccessOptions{Data: respdto.DomainVerificationROToDTO(&domainVerificationView)})
}

func (h *DomainVerificationHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetDomainVerification(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Domain verification not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Domain verification retrieved successfully", httpresp.SuccessOptions{Data: respdto.DomainVerificationROToDTO(&result)})
}

func (h *DomainVerificationHandler) Update(c *fiber.Ctx) error {
	// Update can be Verify or Fail
	var verifyReq reqdto.DomainVerificationVerifyRequestDTO
	if err := c.ParamsParser(&verifyReq); err == nil {
		if err := c.BodyParser(&verifyReq); err == nil {
			cmd := verifyReq.ToVerifyCmd()
			if err := h.appSvc.VerifyDomain(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to verify domain: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Domain verified successfully")
		}
	}

	var failReq reqdto.DomainVerificationFailRequestDTO
	if err := c.ParamsParser(&failReq); err == nil {
		if err := c.BodyParser(&failReq); err == nil {
			cmd := failReq.ToFailCmd()
			if err := h.appSvc.FailDomainVerification(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to fail domain verification: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Domain verification failed successfully")
		}
	}

	return httpresp.Error(c, fiber.StatusBadRequest, "Invalid update request")
}

func (h *DomainVerificationHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := domainVerificationAppCommands.DeleteDomainVerificationCmd{DomainVerificationID: req.ID}
	if err := h.appSvc.DeleteDomainVerification(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete domain verification: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Domain verification deleted successfully")
}
