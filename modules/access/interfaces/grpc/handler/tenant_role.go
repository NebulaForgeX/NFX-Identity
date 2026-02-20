package handler

import (
	"context"

	tenantrolepb "nfxid/protos/gen/access/tenant_role"
	tenantrolesApp "nfxid/modules/access/application/tenant_roles"
	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TenantRoleHandler 实现 TenantRoleServiceServer
type TenantRoleHandler struct {
	tenantrolepb.UnimplementedTenantRoleServiceServer
	svc *tenantrolesApp.Service
}

// NewTenantRoleHandler 创建 handler
func NewTenantRoleHandler(svc *tenantrolesApp.Service) *TenantRoleHandler {
	return &TenantRoleHandler{svc: svc}
}

func (h *TenantRoleHandler) GetTenantRoleByID(ctx context.Context, req *tenantrolepb.GetTenantRoleByIDRequest) (*tenantrolepb.GetTenantRoleByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}
	r, err := h.svc.GetByID(ctx, id)
	if err != nil {
		if err == tenant_roles.ErrTenantRoleNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &tenantrolepb.GetTenantRoleByIDResponse{
		TenantRole: mapper.TenantRoleDomainToProto(r),
	}, nil
}

func (h *TenantRoleHandler) GetTenantRoleByTenantIDAndRoleKey(ctx context.Context, req *tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyRequest) (*tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid tenant_id")
	}
	r, err := h.svc.GetByTenantIDAndRoleKey(ctx, tenantID, req.RoleKey)
	if err != nil {
		if err == tenant_roles.ErrTenantRoleNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyResponse{
		TenantRole: mapper.TenantRoleDomainToProto(r),
	}, nil
}

func (h *TenantRoleHandler) ListTenantRolesByTenantID(ctx context.Context, req *tenantrolepb.ListTenantRolesByTenantIDRequest) (*tenantrolepb.ListTenantRolesByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid tenant_id")
	}
	list, err := h.svc.ListByTenantID(ctx, tenantID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	out := make([]*tenantrolepb.TenantRole, len(list))
	for i := range list {
		out[i] = mapper.TenantRoleDomainToProto(list[i])
	}
	return &tenantrolepb.ListTenantRolesByTenantIDResponse{TenantRoles: out}, nil
}

func (h *TenantRoleHandler) BatchGetTenantRoles(ctx context.Context, req *tenantrolepb.BatchGetTenantRolesRequest) (*tenantrolepb.BatchGetTenantRolesResponse, error) {
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
