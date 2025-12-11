package handler

import (
	"context"

	occupationApp "nfxid/modules/auth/application/profile_occupation"
	occupationAppQueries "nfxid/modules/auth/application/profile_occupation/queries"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	profileoccupationpb "nfxid/protos/gen/auth/profile_occupation"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OccupationHandler struct {
	profileoccupationpb.UnimplementedProfileOccupationServiceServer
	occupationAppSvc *occupationApp.Service
}

func NewOccupationHandler(occupationAppSvc *occupationApp.Service) *OccupationHandler {
	return &OccupationHandler{occupationAppSvc: occupationAppSvc}
}

// GetOccupationByID 根据ID获取职业信息
func (h *OccupationHandler) GetProfileOccupationByID(ctx context.Context, req *profileoccupationpb.GetProfileOccupationByIDRequest) (*profileoccupationpb.GetProfileOccupationByIDResponse, error) {
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
	return &profileoccupationpb.GetProfileOccupationByIDResponse{ProfileOccupation: occupation}, nil
}

// GetOccupationsByProfileID 根据ProfileID获取职业信息列表
func (h *OccupationHandler) GetProfileOccupationsByProfileID(ctx context.Context, req *profileoccupationpb.GetProfileOccupationsByProfileIDRequest) (*profileoccupationpb.GetProfileOccupationsByProfileIDResponse, error) {
	profileID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}

	occupationViews, err := h.occupationAppSvc.GetOccupationsByProfileID(ctx, profileID)
	if err != nil {
		logx.S().Errorf("failed to get occupations by profile_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get occupations: %v", err)
	}

	occupations := make([]*profileoccupationpb.ProfileOccupation, len(occupationViews))
	for i, occupationView := range occupationViews {
		occupations[i] = mapper.OccupationViewToProto(&occupationView)
	}

	return &profileoccupationpb.GetProfileOccupationsByProfileIDResponse{ProfileOccupations: occupations}, nil
}

// GetAllOccupations 获取所有职业信息列表
func (h *OccupationHandler) GetAllProfileOccupations(ctx context.Context, req *profileoccupationpb.GetAllProfileOccupationsRequest) (*profileoccupationpb.GetAllProfileOccupationsResponse, error) {
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

	occupations := make([]*profileoccupationpb.ProfileOccupation, len(result.Items))
	for i, occupationView := range result.Items {
		occupations[i] = mapper.OccupationViewToProto(&occupationView)
	}

	return &profileoccupationpb.GetAllProfileOccupationsResponse{
		ProfileOccupations: occupations,
		Total:              int32(result.Total),
	}, nil
}
