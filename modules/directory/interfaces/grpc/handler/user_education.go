package handler

import (
	"context"

	userEducationApp "nfxid/modules/directory/application/user_educations"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	usereducationpb "nfxid/protos/gen/directory/user_education"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserEducationHandler struct {
	usereducationpb.UnimplementedUserEducationServiceServer
	userEducationAppSvc *userEducationApp.Service
}

func NewUserEducationHandler(userEducationAppSvc *userEducationApp.Service) *UserEducationHandler {
	return &UserEducationHandler{
		userEducationAppSvc: userEducationAppSvc,
	}
}

// GetUserEducationByID 根据ID获取用户教育
func (h *UserEducationHandler) GetUserEducationByID(ctx context.Context, req *usereducationpb.GetUserEducationByIDRequest) (*usereducationpb.GetUserEducationByIDResponse, error) {
	userEducationID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_education_id: %v", err)
	}

	userEducationView, err := h.userEducationAppSvc.GetUserEducation(ctx, userEducationID)
	if err != nil {
		logx.S().Errorf("failed to get user education by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user education not found: %v", err)
	}

	userEducation := mapper.UserEducationROToProto(&userEducationView)
	return &usereducationpb.GetUserEducationByIDResponse{UserEducation: userEducation}, nil
}

// GetUserEducationsByUserID 根据用户ID获取用户教育列表
func (h *UserEducationHandler) GetUserEducationsByUserID(ctx context.Context, req *usereducationpb.GetUserEducationsByUserIDRequest) (*usereducationpb.GetUserEducationsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userEducationViews, err := h.userEducationAppSvc.GetUserEducationsByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user educations by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user educations: %v", err)
	}

	userEducations := mapper.UserEducationListROToProto(userEducationViews)
	return &usereducationpb.GetUserEducationsByUserIDResponse{UserEducations: userEducations}, nil
}
