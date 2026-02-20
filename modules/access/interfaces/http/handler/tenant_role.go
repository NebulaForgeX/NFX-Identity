package handler

import (
	tenantrolesApp "nfxid/modules/access/application/tenant_roles"
	domain "nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type TenantRoleHandler struct{ svc *tenantrolesApp.Service }

func NewTenantRoleHandler(svc *tenantrolesApp.Service) *TenantRoleHandler {
	return &TenantRoleHandler{svc: svc}
}

func (h *TenantRoleHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.TenantRoleByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	r, err := h.svc.GetByID(c.Context(), req.ID)
	if err != nil {
		if err == domain.ErrTenantRoleNotFound {
			return fiberx.ErrorFromErrx(c, errx.NotFound("NOT_FOUND", "tenant role not found").WithCause(err))
		}
		return err
	}
	return fiberx.OK(c, "Tenant role retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantRoleToDTO(r)})
}

func (h *TenantRoleHandler) GetByTenantIDAndRoleKey(c fiber.Ctx) error {
	var req reqdto.TenantRoleByTenantIDAndRoleKeyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	r, err := h.svc.GetByTenantIDAndRoleKey(c.Context(), req.TenantID, req.RoleKey)
	if err != nil {
		if err == domain.ErrTenantRoleNotFound {
			return fiberx.ErrorFromErrx(c, errx.NotFound("NOT_FOUND", "tenant role not found").WithCause(err))
		}
		return err
	}
	return fiberx.OK(c, "Tenant role retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantRoleToDTO(r)})
}

func (h *TenantRoleHandler) ListByTenantID(c fiber.Ctx) error {
	var req reqdto.TenantRoleByTenantIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	list, err := h.svc.ListByTenantID(c.Context(), req.TenantID)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Tenant roles retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantRoleListToDTO(list)})
}

func (h *TenantRoleHandler) Create(c fiber.Ctx) error {
	var req reqdto.TenantRoleCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	tenantID, roleKey, name, err := req.ToCreateParams()
	if err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	r, err := h.svc.Create(c.Context(), tenantrolesApp.CreateParams{TenantID: tenantID, RoleKey: roleKey, Name: name})
	if err != nil {
		if err == domain.ErrTenantRoleKeyExistsInTenant {
			return fiberx.ErrorFromErrx(c, errx.Conflict("CONFLICT", "tenant role key already exists in tenant").WithCause(err))
		}
		return err
	}
	return fiberx.Created(c, "Tenant role created successfully", httpx.SuccessOptions{Data: respdto.TenantRoleToDTO(r)})
}

func (h *TenantRoleHandler) Update(c fiber.Ctx) error {
	var uriReq reqdto.TenantRoleByIDRequestDTO
	if err := c.Bind().URI(&uriReq); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	var body reqdto.TenantRoleUpdateBodyRequestDTO
	if err := c.Bind().Body(&body); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.svc.Update(c.Context(), uriReq.ID, body.RoleKey, body.Name); err != nil {
		if err == domain.ErrTenantRoleNotFound {
			return fiberx.ErrorFromErrx(c, errx.NotFound("NOT_FOUND", "tenant role not found").WithCause(err))
		}
		return err
	}
	return fiberx.OK(c, "Tenant role updated successfully", httpx.SuccessOptions{})
}

func (h *TenantRoleHandler) DeleteByID(c fiber.Ctx) error {
	var req reqdto.TenantRoleByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := h.svc.DeleteByID(c.Context(), req.ID); err != nil {
		if err == domain.ErrTenantRoleNotFound {
			return fiberx.ErrorFromErrx(c, errx.NotFound("NOT_FOUND", "tenant role not found").WithCause(err))
		}
		return err
	}
	return fiberx.OK(c, "Tenant role deleted successfully", httpx.SuccessOptions{})
}
