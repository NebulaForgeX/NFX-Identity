package handler

import (
	"context"

	occupationApp "nebulaid/modules/auth/application/occupation"
	occupationAppQueries "nebulaid/modules/auth/application/occupation/queries"
	"nebulaid/modules/auth/interfaces/grpc/mapper"
	"nebulaid/pkgs/logx"
	occupationpb "nebulaid/protos/gen/auth/occupation"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OccupationHandler struct {
	occupationpb.UnimplementedOccupationServiceServer
	occupationAppSvc *occupationApp.Service
}

func NewOccupationHandler(occupationAppSvc *occupationApp.Service) *OccupationHandler {
	return &OccupationHandler{occupationAppSvc: occupationAppSvc}
}

// GetOccupationByID 根据ID获取职业信息
func (h *OccupationHandler) GetOccupationByID(ctx context.Context, req *occupationpb.GetOccupationByIDRequest) (*occupationpb.GetOccupationByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid occupation_id: %v", err)
	}

	occupationView, err := h.occupationAppSvc.GetOccupation(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get occupation by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "occupation not found: %v", err)
	}

	occupation := mapper.OccupationViewToProto(&occupationView)
	return &occupationpb.GetOccupationByIDResponse{Occupation: occupation}, nil
}

// GetOccupationsByProfileID 根据ProfileID获取职业信息列表
func (h *OccupationHandler) GetOccupationsByProfileID(ctx context.Context, req *occupationpb.GetOccupationsByProfileIDRequest) (*occupationpb.GetOccupationsByProfileIDResponse, error) {
	profileID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}

	occupationViews, err := h.occupationAppSvc.GetOccupationsByProfileID(ctx, profileID)
	if err != nil {
		logx.S().Errorf("failed to get occupations by profile_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get occupations: %v", err)
	}

	occupations := make([]*occupationpb.Occupation, len(occupationViews))
	for i, occupationView := range occupationViews {
		occupations[i] = mapper.OccupationViewToProto(&occupationView)
	}

	return &occupationpb.GetOccupationsByProfileIDResponse{Occupations: occupations}, nil
}

// GetAllOccupations 获取所有职业信息列表
func (h *OccupationHandler) GetAllOccupations(ctx context.Context, req *occupationpb.GetAllOccupationsRequest) (*occupationpb.GetAllOccupationsResponse, error) {
	listQuery := occupationAppQueries.OccupationListQuery{}

	// Set pagination (convert Page/PageSize to Offset/Limit)
	if req.Page > 0 && req.PageSize > 0 {
		listQuery.Offset = int((req.Page - 1) * req.PageSize)
		listQuery.Limit = int(req.PageSize)
	}

	if req.Search != nil {
		search := *req.Search
		listQuery.Search = &search
	}
	if req.ProfileId != nil {
		profileID, err := uuid.Parse(*req.ProfileId)
		if err == nil {
			listQuery.ProfileIDs = []uuid.UUID{profileID}
		}
	}

	result, err := h.occupationAppSvc.GetOccupationList(ctx, listQuery)
	if err != nil {
		logx.S().Errorf("failed to get all occupations: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get occupations: %v", err)
	}

	occupations := make([]*occupationpb.Occupation, len(result.Items))
	for i, occupationView := range result.Items {
		occupations[i] = mapper.OccupationViewToProto(&occupationView)
	}

	return &occupationpb.GetAllOccupationsResponse{
		Occupations: occupations,
		Total:       int32(result.Total),
	}, nil
}
