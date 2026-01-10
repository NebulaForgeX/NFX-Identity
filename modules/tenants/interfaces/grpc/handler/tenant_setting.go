package handler

import (
	"context"

	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TenantSettingHandler struct {
	tenantsettingpb.UnimplementedTenantSettingServiceServer
	tenantSettingAppSvc *tenantSettingApp.Service
}

func NewTenantSettingHandler(tenantSettingAppSvc *tenantSettingApp.Service) *TenantSettingHandler {
	return &TenantSettingHandler{
		tenantSettingAppSvc: tenantSettingAppSvc,
	}
}

// GetTenantSettingByID 根据ID获取租户设置
func (h *TenantSettingHandler) GetTenantSettingByID(ctx context.Context, req *tenantsettingpb.GetTenantSettingByIDRequest) (*tenantsettingpb.GetTenantSettingByIDResponse, error) {
	tenantSettingID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_setting_id: %v", err)
	}

	tenantSettingView, err := h.tenantSettingAppSvc.GetTenantSetting(ctx, tenantSettingID)
	if err != nil {
		logx.S().Errorf("failed to get tenant setting by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "tenant setting not found: %v", err)
	}

	tenantSetting := mapper.TenantSettingROToProto(&tenantSettingView)
	return &tenantsettingpb.GetTenantSettingByIDResponse{TenantSetting: tenantSetting}, nil
}

// GetTenantSettingByTenantID 根据租户ID获取租户设置
func (h *TenantSettingHandler) GetTenantSettingByTenantID(ctx context.Context, req *tenantsettingpb.GetTenantSettingByTenantIDRequest) (*tenantsettingpb.GetTenantSettingByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	tenantSettingView, err := h.tenantSettingAppSvc.GetTenantSettingByTenantID(ctx, tenantID)
	if err != nil {
		logx.S().Errorf("failed to get tenant setting by tenant_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "tenant setting not found: %v", err)
	}

	tenantSetting := mapper.TenantSettingROToProto(&tenantSettingView)
	return &tenantsettingpb.GetTenantSettingByTenantIDResponse{TenantSetting: tenantSetting}, nil
}
