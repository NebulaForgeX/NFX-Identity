package handler

import (
	"context"

	appApp "nfxidentity/modules/clients/application/apps"
	"nfxidentity/modules/clients/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	applicationpb "nfxidentity/protos/gen/clients/application"

	"github.com/google/uuid"
)

type ApplicationHandler struct {
	applicationpb.UnimplementedApplicationServiceServer
	appAppSvc *appApp.Service
}

func NewApplicationHandler(appAppSvc *appApp.Service) *ApplicationHandler {
	return &ApplicationHandler{appAppSvc: appAppSvc}
}

func (h *ApplicationHandler) GetApplicationByID(
	ctx context.Context,
	req *applicationpb.GetApplicationByIDRequest,
) (*applicationpb.GetApplicationByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	ro, err := h.appAppSvc.GetApp(ctx, id)
	if err != nil {
		return nil, err
	}
	app := mapper.ApplicationROToProto(&ro)
	return &applicationpb.GetApplicationByIDResponse{Application: app}, nil
}

func (h *ApplicationHandler) GetApplicationByApplicationID(
	ctx context.Context,
	req *applicationpb.GetApplicationByApplicationIDRequest,
) (*applicationpb.GetApplicationByApplicationIDResponse, error) {
	ro, err := h.appAppSvc.GetAppByAppID(ctx, req.ApplicationId)
	if err != nil {
		return nil, err
	}
	app := mapper.ApplicationROToProto(&ro)
	return &applicationpb.GetApplicationByApplicationIDResponse{Application: app}, nil
}

func (h *ApplicationHandler) GetApplicationsByTenantID(
	ctx context.Context,
	req *applicationpb.GetApplicationsByTenantIDRequest,
) (*applicationpb.GetApplicationsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	list, err := h.appAppSvc.GetAppsByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	apps := mapper.ApplicationListROToProto(list)
	return &applicationpb.GetApplicationsByTenantIDResponse{Applications: apps}, nil
}

func (h *ApplicationHandler) BatchGetApplications(
	ctx context.Context,
	req *applicationpb.BatchGetApplicationsRequest,
) (*applicationpb.BatchGetApplicationsResponse, error) {
	ids := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, errx.ErrInvalidParams.WithCause(err)
		}
		ids = append(ids, id)
	}
	list, err := h.appAppSvc.BatchGetApps(ctx, ids)
	if err != nil {
		return nil, err
	}
	apps := make([]*applicationpb.Application, 0, len(list))
	for i := range list {
		apps = append(apps, mapper.ApplicationROToProto(&list[i]))
	}
	return &applicationpb.BatchGetApplicationsResponse{Applications: apps}, nil
}
