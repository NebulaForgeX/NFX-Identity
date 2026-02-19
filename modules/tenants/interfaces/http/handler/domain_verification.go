package handler

import (
	domainVerificationApp "nfxid/modules/tenants/application/domain_verifications"
	domainVerificationAppCommands "nfxid/modules/tenants/application/domain_verifications/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type DomainVerificationHandler struct {
	appSvc *domainVerificationApp.Service
}

func NewDomainVerificationHandler(appSvc *domainVerificationApp.Service) *DomainVerificationHandler {
	return &DomainVerificationHandler{appSvc: appSvc}
}

func (h *DomainVerificationHandler) Create(c fiber.Ctx) error {
	var req reqdto.DomainVerificationCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	domainVerificationID, err := h.appSvc.CreateDomainVerification(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created domain verification
	domainVerificationView, err := h.appSvc.GetDomainVerification(c.Context(), domainVerificationID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Domain verification created successfully", httpx.SuccessOptions{Data: respdto.DomainVerificationROToDTO(&domainVerificationView)})
}

func (h *DomainVerificationHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetDomainVerification(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Domain verification retrieved successfully", httpx.SuccessOptions{Data: respdto.DomainVerificationROToDTO(&result)})
}

func (h *DomainVerificationHandler) Update(c fiber.Ctx) error {
	// Update can be Verify or Fail
	var verifyReq reqdto.DomainVerificationVerifyRequestDTO
	if err := c.Bind().URI(&verifyReq); err == nil {
		if err := c.Bind().Body(&verifyReq); err == nil {
			cmd := verifyReq.ToVerifyCmd()
			if err := h.appSvc.VerifyDomain(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Domain verified successfully")
		}
	}

	var failReq reqdto.DomainVerificationFailRequestDTO
	if err := c.Bind().URI(&failReq); err == nil {
		if err := c.Bind().Body(&failReq); err == nil {
			cmd := failReq.ToFailCmd()
			if err := h.appSvc.FailDomainVerification(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Domain verification failed successfully")
		}
	}

	return errx.ErrInvalidParams
}

func (h *DomainVerificationHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := domainVerificationAppCommands.DeleteDomainVerificationCmd{DomainVerificationID: req.ID}
	if err := h.appSvc.DeleteDomainVerification(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Domain verification deleted successfully")
}

// fiber:context-methods migrated
