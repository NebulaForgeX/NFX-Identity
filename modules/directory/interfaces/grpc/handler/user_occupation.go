package handler

import (
	"context"

	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	useroccupationpb "nfxid/protos/gen/directory/user_occupation"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_occupation_id: %v", err)
	}

	userOccupationView, err := h.userOccupationAppSvc.GetUserOccupation(ctx, userOccupationID)
	if err != nil {
		logx.S().Errorf("failed to get user occupation by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user occupation not found: %v", err)
	}

	userOccupation := mapper.UserOccupationROToProto(&userOccupationView)
	return &useroccupationpb.GetUserOccupationByIDResponse{UserOccupation: userOccupation}, nil
}

// GetUserOccupationsByUserID 根据用户ID获取用户职业列表
func (h *UserOccupationHandler) GetUserOccupationsByUserID(ctx context.Context, req *useroccupationpb.GetUserOccupationsByUserIDRequest) (*useroccupationpb.GetUserOccupationsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	var isCurrent *bool
	if req.IsCurrent != nil {
		isCurrent = req.IsCurrent
	}

	userOccupationViews, err := h.userOccupationAppSvc.GetUserOccupationsByUserID(ctx, userID, isCurrent)
	if err != nil {
		logx.S().Errorf("failed to get user occupations by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user occupations: %v", err)
	}

	userOccupations := mapper.UserOccupationListROToProto(userOccupationViews)
	return &useroccupationpb.GetUserOccupationsByUserIDResponse{UserOccupations: userOccupations}, nil
}
