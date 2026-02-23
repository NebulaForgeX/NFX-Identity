package handler

import (
	"context"

	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	useroccupationpb "nfxid/protos/gen/directory/user_occupation"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
)

type UserOccupationHandler struct {
	useroccupationpb.UnimplementedUserOccupationServiceServer
	userOccupationAppSvc *userOccupationApp.Service
}

func NewUserOccupationHandler(userOccupationAppSvc *userOccupationApp.Service) *UserOccupationHandler {
	return &UserOccupationHandler{
		userOccupationAppSvc: userOccupationAppSvc,
	}
}

// GetUserOccupationByID 根据ID获取用户职业
func (h *UserOccupationHandler) GetUserOccupationByID(ctx context.Context, req *useroccupationpb.GetUserOccupationByIDRequest) (*useroccupationpb.GetUserOccupationByIDResponse, error) {
	userOccupationID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userOccupationView, err := h.userOccupationAppSvc.GetUserOccupation(ctx, userOccupationID)
	if err != nil {
		return nil, err
	}

	userOccupation := mapper.UserOccupationROToProto(&userOccupationView)
	return &useroccupationpb.GetUserOccupationByIDResponse{UserOccupation: userOccupation}, nil
}

// GetUserOccupationsByUserID 根据用户ID获取用户职业列表
func (h *UserOccupationHandler) GetUserOccupationsByUserID(ctx context.Context, req *useroccupationpb.GetUserOccupationsByUserIDRequest) (*useroccupationpb.GetUserOccupationsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	var isCurrent *bool
	if req.IsCurrent != nil {
		isCurrent = req.IsCurrent
	}

	userOccupationViews, err := h.userOccupationAppSvc.GetUserOccupationsByUserID(ctx, userID, isCurrent)
	if err != nil {
		return nil, err
	}

	userOccupations := mapper.UserOccupationListROToProto(userOccupationViews)
	return &useroccupationpb.GetUserOccupationsByUserIDResponse{UserOccupations: userOccupations}, nil
}
