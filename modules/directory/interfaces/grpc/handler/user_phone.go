package handler

import (
	"context"

	userPhoneApp "nfxid/modules/directory/application/user_phones"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userphonepb "nfxid/protos/gen/directory/user_phone"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserPhoneHandler struct {
	userphonepb.UnimplementedUserPhoneServiceServer
	userPhoneAppSvc *userPhoneApp.Service
}

func NewUserPhoneHandler(userPhoneAppSvc *userPhoneApp.Service) *UserPhoneHandler {
	return &UserPhoneHandler{
		userPhoneAppSvc: userPhoneAppSvc,
	}
}

// GetUserPhoneByID 根据ID获取用户手机
func (h *UserPhoneHandler) GetUserPhoneByID(ctx context.Context, req *userphonepb.GetUserPhoneByIDRequest) (*userphonepb.GetUserPhoneByIDResponse, error) {
	userPhoneID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_phone_id: %v", err)
	}

	userPhoneView, err := h.userPhoneAppSvc.GetUserPhone(ctx, userPhoneID)
	if err != nil {
		logx.S().Errorf("failed to get user phone by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user phone not found: %v", err)
	}

	userPhone := mapper.UserPhoneROToProto(&userPhoneView)
	return &userphonepb.GetUserPhoneByIDResponse{UserPhone: userPhone}, nil
}

// GetUserPhoneByPhone 根据手机号获取用户手机
func (h *UserPhoneHandler) GetUserPhoneByPhone(ctx context.Context, req *userphonepb.GetUserPhoneByPhoneRequest) (*userphonepb.GetUserPhoneByPhoneResponse, error) {
	userPhoneView, err := h.userPhoneAppSvc.GetUserPhoneByPhone(ctx, req.Phone)
	if err != nil {
		logx.S().Errorf("failed to get user phone by phone: %v", err)
		return nil, status.Errorf(codes.NotFound, "user phone not found: %v", err)
	}

	userPhone := mapper.UserPhoneROToProto(&userPhoneView)
	return &userphonepb.GetUserPhoneByPhoneResponse{UserPhone: userPhone}, nil
}

// GetUserPhonesByUserID 根据用户ID获取用户手机列表
func (h *UserPhoneHandler) GetUserPhonesByUserID(ctx context.Context, req *userphonepb.GetUserPhonesByUserIDRequest) (*userphonepb.GetUserPhonesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userPhoneViews, err := h.userPhoneAppSvc.GetUserPhonesByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user phones by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user phones: %v", err)
	}

	userPhones := mapper.UserPhoneListROToProto(userPhoneViews)
	return &userphonepb.GetUserPhonesByUserIDResponse{UserPhones: userPhones}, nil
}
