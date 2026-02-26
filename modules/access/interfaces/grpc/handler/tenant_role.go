package handler

import (
	"context"

	tenantrolesApp "nfxidentity/modules/access/application/tenant_roles"
	"nfxidentity/modules/access/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	tenantrolepb "nfxidentity/protos/gen/access/tenant_role"

	"github.com/google/uuid"
)

type TenantRoleHandler struct {
	tenantrolepb.UnimplementedTenantRoleServiceServer
	svc *tenantrolesApp.Service
}

func NewTenantRoleHandler(svc *tenantrolesApp.Service) *TenantRoleHandler {
	return &TenantRoleHandler{svc: svc}
}

func (h *TenantRoleHandler) GetTenantRoleByID(
	ctx context.Context,
	req *tenantrolepb.GetTenantRoleByIDRequest,
) (*tenantrolepb.GetTenantRoleByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	r, err := h.svc.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &tenantrolepb.GetTenantRoleByIDResponse{
		TenantRole: mapper.TenantRoleDomainToProto(r),
	}, nil
}

func (h *TenantRoleHandler) GetTenantRoleByTenantIDAndRoleKey(
	ctx context.Context,
	req *tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyRequest,
) (*tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	r, err := h.svc.GetByTenantIDAndRoleKey(ctx, tenantID, req.RoleKey)
	if err != nil {
		return nil, err
	}
	return &tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyResponse{
		TenantRole: mapper.TenantRoleDomainToProto(r),
	}, nil
}

func (h *TenantRoleHandler) ListTenantRolesByTenantID(
	ctx context.Context,
	req *tenantrolepb.ListTenantRolesByTenantIDRequest,
) (*tenantrolepb.ListTenantRolesByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	list, err := h.svc.ListByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	out := make([]*tenantrolepb.TenantRole, len(list))
	for i := range list {
		out[i] = mapper.TenantRoleDomainToProto(list[i])
	}
	return &tenantrolepb.ListTenantRolesByTenantIDResponse{TenantRoles: out}, nil
}

func (h *TenantRoleHandler) BatchGetTenantRoles(
	ctx context.Context,
	req *tenantrolepb.BatchGetTenantRolesRequest,
) (*tenantrolepb.BatchGetTenantRolesResponse, error) {
	var out []*tenantrolepb.TenantRole
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		r, err := h.svc.GetByID(ctx, id)
		if err != nil {
			continue
		}
		out = append(out, mapper.TenantRoleDomainToProto(r))
	}
	return &tenantrolepb.BatchGetTenantRolesResponse{TenantRoles: out}, nil
}
