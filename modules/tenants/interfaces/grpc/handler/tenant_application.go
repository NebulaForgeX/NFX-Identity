package handler

import (
	"context"

	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	tenantapplicationpb "nfxid/protos/gen/tenants/tenant_application"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TenantApplicationHandler struct {
	tenantapplicationpb.UnimplementedTenantApplicationServiceServer
	tenantAppAppSvc *tenantAppApp.Service
}

func NewTenantApplicationHandler(tenantAppAppSvc *tenantAppApp.Service) *TenantApplicationHandler {
	return &TenantApplicationHandler{tenantAppAppSvc: tenantAppAppSvc}
}

func (h *TenantApplicationHandler) GetTenantApplicationByID(ctx context.Context, req *tenantapplicationpb.GetTenantApplicationByIDRequest) (*tenantapplicationpb.GetTenantApplicationByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_application_id: %v", err)
	}
	ro, err := h.tenantAppAppSvc.GetTenantApp(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get tenant application by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "tenant application not found: %v", err)
	}
	ta := mapper.TenantApplicationROToProto(&ro)
	return &tenantapplicationpb.GetTenantApplicationByIDResponse{TenantApplication: ta}, nil
}

func (h *TenantApplicationHandler) GetTenantApplicationsByTenantID(ctx context.Context, req *tenantapplicationpb.GetTenantApplicationsByTenantIDRequest) (*tenantapplicationpb.GetTenantApplicationsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}
	list, err := h.tenantAppAppSvc.GetTenantAppsByTenantID(ctx, tenantID)
	if err != nil {
		logx.S().Errorf("failed to get tenant applications by tenant_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get tenant applications: %v", err)
	}
	tas := mapper.TenantApplicationListROToProto(list)
	return &tenantapplicationpb.GetTenantApplicationsByTenantIDResponse{TenantApplications: tas}, nil
}

func (h *TenantApplicationHandler) GetTenantApplicationsByApplicationID(ctx context.Context, req *tenantapplicationpb.GetTenantApplicationsByApplicationIDRequest) (*tenantapplicationpb.GetTenantApplicationsByApplicationIDResponse, error) {
	applicationID, err := uuid.Parse(req.ApplicationId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid application_id: %v", err)
	}
	list, err := h.tenantAppAppSvc.GetTenantAppsByAppID(ctx, applicationID)
	if err != nil {
		logx.S().Errorf("failed to get tenant applications by application_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get tenant applications: %v", err)
	}
	tas := mapper.TenantApplicationListROToProto(list)
	return &tenantapplicationpb.GetTenantApplicationsByApplicationIDResponse{TenantApplications: tas}, nil
}
