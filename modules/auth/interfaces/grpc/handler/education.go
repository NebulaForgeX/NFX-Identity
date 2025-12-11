package handler

import (
	"context"

	educationApp "nfxid/modules/auth/application/profile_education"
	educationAppQueries "nfxid/modules/auth/application/profile_education/queries"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	profileeducationpb "nfxid/protos/gen/auth/profile_education"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EducationHandler struct {
	profileeducationpb.UnimplementedProfileEducationServiceServer
	educationAppSvc *educationApp.Service
}

func NewEducationHandler(educationAppSvc *educationApp.Service) *EducationHandler {
	return &EducationHandler{educationAppSvc: educationAppSvc}
}

// GetEducationByID 根据ID获取教育经历
func (h *EducationHandler) GetProfileEducationByID(ctx context.Context, req *profileeducationpb.GetProfileEducationByIDRequest) (*profileeducationpb.GetProfileEducationByIDResponse, error) {
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
	return &profileeducationpb.GetProfileEducationByIDResponse{ProfileEducation: education}, nil
}

// GetEducationsByProfileID 根据ProfileID获取教育经历列表
func (h *EducationHandler) GetProfileEducationsByProfileID(ctx context.Context, req *profileeducationpb.GetProfileEducationsByProfileIDRequest) (*profileeducationpb.GetProfileEducationsByProfileIDResponse, error) {
	profileID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}

	educationViews, err := h.educationAppSvc.GetEducationsByProfileID(ctx, profileID)
	if err != nil {
		logx.S().Errorf("failed to get educations by profile_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get educations: %v", err)
	}

	educations := make([]*profileeducationpb.ProfileEducation, len(educationViews))
	for i, educationView := range educationViews {
		educations[i] = mapper.EducationViewToProto(&educationView)
	}

	return &profileeducationpb.GetProfileEducationsByProfileIDResponse{ProfileEducations: educations}, nil
}

// GetAllEducations 获取所有教育经历列表
func (h *EducationHandler) GetAllProfileEducations(ctx context.Context, req *profileeducationpb.GetAllProfileEducationsRequest) (*profileeducationpb.GetAllProfileEducationsResponse, error) {
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

	educations := make([]*profileeducationpb.ProfileEducation, len(result.Items))
	for i, educationView := range result.Items {
		educations[i] = mapper.EducationViewToProto(&educationView)
	}

	return &profileeducationpb.GetAllProfileEducationsResponse{
		ProfileEducations: educations,
		Total:             int32(result.Total),
	}, nil
}
