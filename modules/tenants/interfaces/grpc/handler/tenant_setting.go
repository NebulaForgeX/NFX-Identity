package handler

import (
	"context"

	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
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
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	tenantSettingView, err := h.tenantSettingAppSvc.GetTenantSetting(ctx, tenantSettingID)
	if err != nil {
		return nil, err
	}

	tenantSetting := mapper.TenantSettingROToProto(&tenantSettingView)
	return &tenantsettingpb.GetTenantSettingByIDResponse{TenantSetting: tenantSetting}, nil
}

// GetTenantSettingByTenantID 根据租户ID获取租户设置
func (h *TenantSettingHandler) GetTenantSettingByTenantID(ctx context.Context, req *tenantsettingpb.GetTenantSettingByTenantIDRequest) (*tenantsettingpb.GetTenantSettingByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	tenantSettingView, err := h.tenantSettingAppSvc.GetTenantSettingByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	tenantSetting := mapper.TenantSettingROToProto(&tenantSettingView)
	return &tenantsettingpb.GetTenantSettingByTenantIDResponse{TenantSetting: tenantSetting}, nil
}
