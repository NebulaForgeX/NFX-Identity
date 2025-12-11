package handler

import (
	"context"

	educationApp "nfxid/modules/auth/application/education"
	educationAppQueries "nfxid/modules/auth/application/education/queries"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	educationpb "nfxid/protos/gen/auth/education"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EducationHandler struct {
	educationpb.UnimplementedEducationServiceServer
	educationAppSvc *educationApp.Service
}

func NewEducationHandler(educationAppSvc *educationApp.Service) *EducationHandler {
	return &EducationHandler{educationAppSvc: educationAppSvc}
}

// GetEducationByID 根据ID获取教育经历
func (h *EducationHandler) GetEducationByID(ctx context.Context, req *educationpb.GetEducationByIDRequest) (*educationpb.GetEducationByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid education_id: %v", err)
	}

	educationView, err := h.educationAppSvc.GetEducation(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get education by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "education not found: %v", err)
	}

	education := mapper.EducationViewToProto(&educationView)
	return &educationpb.GetEducationByIDResponse{Education: education}, nil
}

// GetEducationsByProfileID 根据ProfileID获取教育经历列表
func (h *EducationHandler) GetEducationsByProfileID(ctx context.Context, req *educationpb.GetEducationsByProfileIDRequest) (*educationpb.GetEducationsByProfileIDResponse, error) {
	profileID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}

	educationViews, err := h.educationAppSvc.GetEducationsByProfileID(ctx, profileID)
	if err != nil {
		logx.S().Errorf("failed to get educations by profile_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get educations: %v", err)
	}

	educations := make([]*educationpb.Education, len(educationViews))
	for i, educationView := range educationViews {
		educations[i] = mapper.EducationViewToProto(&educationView)
	}

	return &educationpb.GetEducationsByProfileIDResponse{Educations: educations}, nil
}

// GetAllEducations 获取所有教育经历列表
func (h *EducationHandler) GetAllEducations(ctx context.Context, req *educationpb.GetAllEducationsRequest) (*educationpb.GetAllEducationsResponse, error) {
	listQuery := educationAppQueries.EducationListQuery{}

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

	result, err := h.educationAppSvc.GetEducationList(ctx, listQuery)
	if err != nil {
		logx.S().Errorf("failed to get all educations: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get educations: %v", err)
	}

	educations := make([]*educationpb.Education, len(result.Items))
	for i, educationView := range result.Items {
		educations[i] = mapper.EducationViewToProto(&educationView)
	}

	return &educationpb.GetAllEducationsResponse{
		Educations: educations,
		Total:      int32(result.Total),
	}, nil
}
