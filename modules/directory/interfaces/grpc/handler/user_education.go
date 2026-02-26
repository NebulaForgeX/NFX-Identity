package handler

import (
	"context"

	userEducationApp "nfxidentity/modules/directory/application/user_educations"
	"nfxidentity/modules/directory/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	usereducationpb "nfxidentity/protos/gen/directory/user_education"

	"github.com/google/uuid"
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
func (h *UserEducationHandler) GetUserEducationByID(
	ctx context.Context,
	req *usereducationpb.GetUserEducationByIDRequest,
) (*usereducationpb.GetUserEducationByIDResponse, error) {
	userEducationID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userEducationView, err := h.userEducationAppSvc.GetUserEducation(ctx, userEducationID)
	if err != nil {
		return nil, err
	}

	userEducation := mapper.UserEducationROToProto(&userEducationView)
	return &usereducationpb.GetUserEducationByIDResponse{UserEducation: userEducation}, nil
}

// GetUserEducationsByUserID 根据用户ID获取用户教育列表
func (h *UserEducationHandler) GetUserEducationsByUserID(
	ctx context.Context,
	req *usereducationpb.GetUserEducationsByUserIDRequest,
) (*usereducationpb.GetUserEducationsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userEducationViews, err := h.userEducationAppSvc.GetUserEducationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userEducations := mapper.UserEducationListROToProto(userEducationViews)
	return &usereducationpb.GetUserEducationsByUserIDResponse{UserEducations: userEducations}, nil
}
