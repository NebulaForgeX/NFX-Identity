package handler

import (
	"context"

	profileApp "nebulaid/modules/auth/application/profile"
	roleApp "nebulaid/modules/auth/application/role"
	userApp "nebulaid/modules/auth/application/user"
	userDomain "nebulaid/modules/auth/domain/user"
	"nebulaid/modules/auth/interfaces/grpc/mapper"
	"nebulaid/pkgs/logx"
	authpb "nebulaid/protos/gen/auth/auth"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	userAppSvc    *userApp.Service
	profileAppSvc *profileApp.Service
	roleAppSvc    *roleApp.Service
	userRepo      userDomain.Repo
}

func NewAuthHandler(
	userAppSvc *userApp.Service,
	profileAppSvc *profileApp.Service,
	roleAppSvc *roleApp.Service,
	userRepo userDomain.Repo,
) *AuthHandler {
	return &AuthHandler{
		userAppSvc:    userAppSvc,
		profileAppSvc: profileAppSvc,
		roleAppSvc:    roleAppSvc,
		userRepo:      userRepo,
	}
}

// GetAuthByUserID 根据用户ID获取认证信息
func (h *AuthHandler) GetAuthByUserID(ctx context.Context, req *authpb.GetAuthByUserIDRequest) (*authpb.GetAuthByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userView, err := h.userAppSvc.GetUser(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get auth by user_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	auth := mapper.UserViewToProtoAuth(&userView)
	return &authpb.GetAuthByUserIDResponse{Auth: auth}, nil
}

// GetAuthByUsername 根据用户名获取认证信息
func (h *AuthHandler) GetAuthByUsername(ctx context.Context, req *authpb.GetAuthByUsernameRequest) (*authpb.GetAuthByUsernameResponse, error) {
	entity, err := h.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	userView, err := h.userAppSvc.GetUser(ctx, entity.ID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	auth := mapper.UserViewToProtoAuth(&userView)
	return &authpb.GetAuthByUsernameResponse{Auth: auth}, nil
}

// GetAuthByEmail 根据邮箱获取认证信息
func (h *AuthHandler) GetAuthByEmail(ctx context.Context, req *authpb.GetAuthByEmailRequest) (*authpb.GetAuthByEmailResponse, error) {
	entity, err := h.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	userView, err := h.userAppSvc.GetUser(ctx, entity.ID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	auth := mapper.UserViewToProtoAuth(&userView)
	return &authpb.GetAuthByEmailResponse{Auth: auth}, nil
}

// GetAuthByPhone 根据手机号获取认证信息
func (h *AuthHandler) GetAuthByPhone(ctx context.Context, req *authpb.GetAuthByPhoneRequest) (*authpb.GetAuthByPhoneResponse, error) {
	entity, err := h.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	userView, err := h.userAppSvc.GetUser(ctx, entity.ID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	auth := mapper.UserViewToProtoAuth(&userView)
	return &authpb.GetAuthByPhoneResponse{Auth: auth}, nil
}

// BatchGetAuth 批量获取认证信息
func (h *AuthHandler) BatchGetAuth(ctx context.Context, req *authpb.BatchGetAuthRequest) (*authpb.BatchGetAuthResponse, error) {
	userIDs := make([]uuid.UUID, 0, len(req.UserIds))
	for _, idStr := range req.UserIds {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		userIDs = append(userIDs, id)
	}

	auths := make([]*authpb.Auth, 0, len(userIDs))
	errorById := make(map[string]string)

	for _, userID := range userIDs {
		userView, err := h.userAppSvc.GetUser(ctx, userID)
		if err != nil {
			errorById[userID.String()] = err.Error()
			continue
		}
		auth := mapper.UserViewToProtoAuth(&userView)
		auths = append(auths, auth)
	}

	return &authpb.BatchGetAuthResponse{
		Auths:     auths,
		ErrorById: errorById,
	}, nil
}

// VerifyPassword 验证密码
func (h *AuthHandler) VerifyPassword(ctx context.Context, req *authpb.VerifyPasswordRequest) (*authpb.VerifyPasswordResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	entity, err := h.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(entity.Editable().Password), []byte(req.Password)); err != nil {
		return &authpb.VerifyPasswordResponse{Valid: false}, nil
	}

	return &authpb.VerifyPasswordResponse{Valid: true}, nil
}

// VerifyUserExists 验证用户是否存在
func (h *AuthHandler) VerifyUserExists(ctx context.Context, req *authpb.VerifyUserExistsRequest) (*authpb.VerifyUserExistsResponse, error) {
	if req.Username != nil && *req.Username != "" {
		exists, err := h.userRepo.ExistsByUsername(ctx, *req.Username)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
		}
		if exists {
			return &authpb.VerifyUserExistsResponse{Exists: true, Field: "username"}, nil
		}
	}

	if req.Email != nil && *req.Email != "" {
		exists, err := h.userRepo.ExistsByEmail(ctx, *req.Email)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
		}
		if exists {
			return &authpb.VerifyUserExistsResponse{Exists: true, Field: "email"}, nil
		}
	}

	if req.Phone != nil && *req.Phone != "" {
		exists, err := h.userRepo.ExistsByPhone(ctx, *req.Phone)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
		}
		if exists {
			return &authpb.VerifyUserExistsResponse{Exists: true, Field: "phone"}, nil
		}
	}

	return &authpb.VerifyUserExistsResponse{Exists: false, Field: ""}, nil
}
