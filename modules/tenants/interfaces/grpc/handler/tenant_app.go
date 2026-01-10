package handler

import (
	"context"

	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	tenantapppb "nfxid/protos/gen/tenants/tenant_app"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TenantAppHandler struct {
	tenantapppb.UnimplementedTenantAppServiceServer
	tenantAppAppSvc *tenantAppApp.Service
}

func NewTenantAppHandler(tenantAppAppSvc *tenantAppApp.Service) *TenantAppHandler {
	return &TenantAppHandler{
		tenantAppAppSvc: tenantAppAppSvc,
	}
}

// GetTenantAppByID 根据ID获取租户应用
func (h *TenantAppHandler) GetTenantAppByID(ctx context.Context, req *tenantapppb.GetTenantAppByIDRequest) (*tenantapppb.GetTenantAppByIDResponse, error) {
	tenantAppID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_app_id: %v", err)
	}

	tenantAppView, err := h.tenantAppAppSvc.GetTenantApp(ctx, tenantAppID)
	if err != nil {
		logx.S().Errorf("failed to get tenant app by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "tenant app not found: %v", err)
	}

	tenantApp := mapper.TenantAppROToProto(&tenantAppView)
	return &tenantapppb.GetTenantAppByIDResponse{TenantApp: tenantApp}, nil
}

// GetTenantAppsByTenantID 根据租户ID获取租户应用列表
func (h *TenantAppHandler) GetTenantAppsByTenantID(ctx context.Context, req *tenantapppb.GetTenantAppsByTenantIDRequest) (*tenantapppb.GetTenantAppsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	tenantAppViews, err := h.tenantAppAppSvc.GetTenantAppsByTenantID(ctx, tenantID)
	if err != nil {
		logx.S().Errorf("failed to get tenant apps by tenant_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get tenant apps: %v", err)
	}

	tenantApps := mapper.TenantAppListROToProto(tenantAppViews)
	return &tenantapppb.GetTenantAppsByTenantIDResponse{TenantApps: tenantApps}, nil
}

// GetTenantAppsByAppID 根据应用ID获取租户应用列表
func (h *TenantAppHandler) GetTenantAppsByAppID(ctx context.Context, req *tenantapppb.GetTenantAppsByAppIDRequest) (*tenantapppb.GetTenantAppsByAppIDResponse, error) {
	appID, err := uuid.Parse(req.AppId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid app_id: %v", err)
	}

	tenantAppViews, err := h.tenantAppAppSvc.GetTenantAppsByAppID(ctx, appID)
	if err != nil {
		logx.S().Errorf("failed to get tenant apps by app_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get tenant apps: %v", err)
	}

	tenantApps := mapper.TenantAppListROToProto(tenantAppViews)
	return &tenantapppb.GetTenantAppsByAppIDResponse{TenantApps: tenantApps}, nil
}
