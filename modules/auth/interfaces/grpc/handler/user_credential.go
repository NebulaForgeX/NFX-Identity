package handler

import (
	"context"

	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserCredentialHandler struct {
	usercredentialpb.UnimplementedUserCredentialServiceServer
	userCredentialAppSvc *userCredentialApp.Service
}

func NewUserCredentialHandler(userCredentialAppSvc *userCredentialApp.Service) *UserCredentialHandler {
	return &UserCredentialHandler{
		userCredentialAppSvc: userCredentialAppSvc,
	}
}

// GetUserCredentialByID 根据ID获取用户凭证
func (h *UserCredentialHandler) GetUserCredentialByID(ctx context.Context, req *usercredentialpb.GetUserCredentialByIDRequest) (*usercredentialpb.GetUserCredentialByIDResponse, error) {
	userCredentialID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_credential_id: %v", err)
	}

	userCredentialView, err := h.userCredentialAppSvc.GetUserCredential(ctx, userCredentialID)
	if err != nil {
		logx.S().Errorf("failed to get user credential by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user credential not found: %v", err)
	}

	userCredential := mapper.UserCredentialROToProto(&userCredentialView)
	return &usercredentialpb.GetUserCredentialByIDResponse{UserCredential: userCredential}, nil
}

// GetUserCredentialByUserID 根据UserID获取用户凭证
func (h *UserCredentialHandler) GetUserCredentialByUserID(ctx context.Context, req *usercredentialpb.GetUserCredentialByUserIDRequest) (*usercredentialpb.GetUserCredentialByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userCredentialView, err := h.userCredentialAppSvc.GetUserCredentialByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user credential by user_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user credential not found: %v", err)
	}

	userCredential := mapper.UserCredentialROToProto(&userCredentialView)
	return &usercredentialpb.GetUserCredentialByUserIDResponse{UserCredential: userCredential}, nil
}
