package handler

import (
	"context"

	appApp "nfxid/modules/clients/application/apps"
	"nfxid/modules/clients/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	apppb "nfxid/protos/gen/clients/app"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppHandler struct {
	apppb.UnimplementedAppServiceServer
	appAppSvc *appApp.Service
}

func NewAppHandler(appAppSvc *appApp.Service) *AppHandler {
	return &AppHandler{
		appAppSvc: appAppSvc,
	}
}

// GetAppByID 根据ID获取应用
func (h *AppHandler) GetAppByID(ctx context.Context, req *apppb.GetAppByIDRequest) (*apppb.GetAppByIDResponse, error) {
	appID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid app_id: %v", err)
	}

	appView, err := h.appAppSvc.GetApp(ctx, appID)
	if err != nil {
		logx.S().Errorf("failed to get app by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "app not found: %v", err)
	}

	app := mapper.AppROToProto(&appView)
	return &apppb.GetAppByIDResponse{App: app}, nil
}

// GetAppByAppID 根据应用标识符获取应用
func (h *AppHandler) GetAppByAppID(ctx context.Context, req *apppb.GetAppByAppIDRequest) (*apppb.GetAppByAppIDResponse, error) {
	appView, err := h.appAppSvc.GetAppByAppID(ctx, req.AppId)
	if err != nil {
		logx.S().Errorf("failed to get app by app_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "app not found: %v", err)
	}

	app := mapper.AppROToProto(&appView)
	return &apppb.GetAppByAppIDResponse{App: app}, nil
}

// GetAppsByTenantID 根据租户ID获取应用列表
func (h *AppHandler) GetAppsByTenantID(ctx context.Context, req *apppb.GetAppsByTenantIDRequest) (*apppb.GetAppsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	appViews, err := h.appAppSvc.GetAppsByTenantID(ctx, tenantID)
	if err != nil {
		logx.S().Errorf("failed to get apps by tenant_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "apps not found: %v", err)
	}

	apps := mapper.AppListROToProto(appViews)
	return &apppb.GetAppsByTenantIDResponse{Apps: apps}, nil
}

// BatchGetApps 批量获取应用
func (h *AppHandler) BatchGetApps(ctx context.Context, req *apppb.BatchGetAppsRequest) (*apppb.BatchGetAppsResponse, error) {
	appIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		appIDs = append(appIDs, id)
	}

	apps := make([]*apppb.App, 0, len(appIDs))
	for _, appID := range appIDs {
		appView, err := h.appAppSvc.GetApp(ctx, appID)
		if err != nil {
			logx.S().Warnf("failed to get app %s: %v", appID, err)
			continue
		}
		app := mapper.AppROToProto(&appView)
		apps = append(apps, app)
	}

	return &apppb.BatchGetAppsResponse{Apps: apps}, nil
}
