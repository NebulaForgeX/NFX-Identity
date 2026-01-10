package handler

import (
	"context"

	tenantApp "nfxid/modules/tenants/application/tenants"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	tenantpb "nfxid/protos/gen/tenants/tenant"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TenantHandler struct {
	tenantpb.UnimplementedTenantServiceServer
	tenantAppSvc *tenantApp.Service
}

func NewTenantHandler(tenantAppSvc *tenantApp.Service) *TenantHandler {
	return &TenantHandler{
		tenantAppSvc: tenantAppSvc,
	}
}

// GetTenantByID 根据ID获取租户
func (h *TenantHandler) GetTenantByID(ctx context.Context, req *tenantpb.GetTenantByIDRequest) (*tenantpb.GetTenantByIDResponse, error) {
	tenantID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	tenantView, err := h.tenantAppSvc.GetTenant(ctx, tenantID)
	if err != nil {
		logx.S().Errorf("failed to get tenant by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "tenant not found: %v", err)
	}

	tenant := mapper.TenantROToProto(&tenantView)
	return &tenantpb.GetTenantByIDResponse{Tenant: tenant}, nil
}

// GetTenantByTenantID 根据租户标识符获取租户
func (h *TenantHandler) GetTenantByTenantID(ctx context.Context, req *tenantpb.GetTenantByTenantIDRequest) (*tenantpb.GetTenantByTenantIDResponse, error) {
	tenantView, err := h.tenantAppSvc.GetTenantByTenantID(ctx, req.TenantId)
	if err != nil {
		logx.S().Errorf("failed to get tenant by tenant_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "tenant not found: %v", err)
	}

	tenant := mapper.TenantROToProto(&tenantView)
	return &tenantpb.GetTenantByTenantIDResponse{Tenant: tenant}, nil
}

// BatchGetTenants 批量获取租户
func (h *TenantHandler) BatchGetTenants(ctx context.Context, req *tenantpb.BatchGetTenantsRequest) (*tenantpb.BatchGetTenantsResponse, error) {
	tenantIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		tenantIDs = append(tenantIDs, id)
	}

	tenants := make([]*tenantpb.Tenant, 0, len(tenantIDs))
	for _, tenantID := range tenantIDs {
		tenantView, err := h.tenantAppSvc.GetTenant(ctx, tenantID)
		if err != nil {
			logx.S().Warnf("failed to get tenant %s: %v", tenantID, err)
			continue
		}
		tenant := mapper.TenantROToProto(&tenantView)
		tenants = append(tenants, tenant)
	}

	return &tenantpb.BatchGetTenantsResponse{Tenants: tenants}, nil
}
