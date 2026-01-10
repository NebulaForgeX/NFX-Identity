package handler

import (
	"context"

	userEmailApp "nfxid/modules/directory/application/user_emails"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	useremailpb "nfxid/protos/gen/directory/user_email"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserEmailHandler struct {
	useremailpb.UnimplementedUserEmailServiceServer
	userEmailAppSvc *userEmailApp.Service
}

func NewUserEmailHandler(userEmailAppSvc *userEmailApp.Service) *UserEmailHandler {
	return &UserEmailHandler{
		userEmailAppSvc: userEmailAppSvc,
	}
}

// GetUserEmailByID 根据ID获取用户邮箱
func (h *UserEmailHandler) GetUserEmailByID(ctx context.Context, req *useremailpb.GetUserEmailByIDRequest) (*useremailpb.GetUserEmailByIDResponse, error) {
	userEmailID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_email_id: %v", err)
	}

	userEmailView, err := h.userEmailAppSvc.GetUserEmail(ctx, userEmailID)
	if err != nil {
		logx.S().Errorf("failed to get user email by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user email not found: %v", err)
	}

	userEmail := mapper.UserEmailROToProto(&userEmailView)
	return &useremailpb.GetUserEmailByIDResponse{UserEmail: userEmail}, nil
}

// GetUserEmailByEmail 根据邮箱地址获取用户邮箱
func (h *UserEmailHandler) GetUserEmailByEmail(ctx context.Context, req *useremailpb.GetUserEmailByEmailRequest) (*useremailpb.GetUserEmailByEmailResponse, error) {
	userEmailView, err := h.userEmailAppSvc.GetUserEmailByEmail(ctx, req.Email)
	if err != nil {
		logx.S().Errorf("failed to get user email by email: %v", err)
		return nil, status.Errorf(codes.NotFound, "user email not found: %v", err)
	}

	userEmail := mapper.UserEmailROToProto(&userEmailView)
	return &useremailpb.GetUserEmailByEmailResponse{UserEmail: userEmail}, nil
}

// GetUserEmailsByUserID 根据用户ID获取用户邮箱列表
func (h *UserEmailHandler) GetUserEmailsByUserID(ctx context.Context, req *useremailpb.GetUserEmailsByUserIDRequest) (*useremailpb.GetUserEmailsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userEmailViews, err := h.userEmailAppSvc.GetUserEmailsByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user emails by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user emails: %v", err)
	}

	userEmails := mapper.UserEmailListROToProto(userEmailViews)
	return &useremailpb.GetUserEmailsByUserIDResponse{UserEmails: userEmails}, nil
}
