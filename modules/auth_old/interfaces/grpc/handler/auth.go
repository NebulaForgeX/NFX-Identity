package handler

import (
	"context"

	profileApp "nfxid/modules/auth/application/profile"
	profileAppCommands "nfxid/modules/auth/application/profile/commands"
	roleApp "nfxid/modules/auth/application/role"
	userApp "nfxid/modules/auth/application/user"
	profileDomain "nfxid/modules/auth/domain/profile"
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	authpb "nfxid/protos/gen/auth/auth"

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
	userRepo      *userDomain.Repo
}

func NewAuthHandler(
	userAppSvc *userApp.Service,
	profileAppSvc *profileApp.Service,
	roleAppSvc *roleApp.Service,
	userRepo *userDomain.Repo,
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
	entity, err := h.userRepo.Get.ByUsername(ctx, req.Username)
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
	entity, err := h.userRepo.Get.ByEmail(ctx, req.Email)
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
	entity, err := h.userRepo.Get.ByPhone(ctx, req.Phone)
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

	entity, err := h.userRepo.Get.ByID(ctx, userID)
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
		exists, err := h.userRepo.Check.ByUsername(ctx, *req.Username)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
		}
		if exists {
			return &authpb.VerifyUserExistsResponse{Exists: true, Field: "username"}, nil
		}
	}

	if req.Email != nil && *req.Email != "" {
		exists, err := h.userRepo.Check.ByEmail(ctx, *req.Email)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
		}
		if exists {
			return &authpb.VerifyUserExistsResponse{Exists: true, Field: "email"}, nil
		}
	}

	if req.Phone != nil && *req.Phone != "" {
		exists, err := h.userRepo.Check.ByPhone(ctx, *req.Phone)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
		}
		if exists {
			return &authpb.VerifyUserExistsResponse{Exists: true, Field: "phone"}, nil
		}
	}

	return &authpb.VerifyUserExistsResponse{Exists: false, Field: ""}, nil
}

// SendVerificationCode 发送验证码（发送邮件，存储到 Redis）
func (h *AuthHandler) SendVerificationCode(ctx context.Context, req *authpb.SendVerificationCodeRequest) (*authpb.SendVerificationCodeResponse, error) {
	err := h.userAppSvc.SendVerificationCode(ctx, userApp.SendVerificationCodeCmd{
		Email:   req.Email,
		Purpose: "register", // 固定为注册用途
	})
	if err != nil {
		logx.S().Errorf("failed to send verification code: %v", err)
		return &authpb.SendVerificationCodeResponse{
			Success:      false,
			ErrorMessage: &[]string{err.Error()}[0],
		}, nil
	}

	return &authpb.SendVerificationCodeResponse{Success: true}, nil
}

// CheckUserAndVerificationCode 检查用户是否存在并验证验证码（用于注册流程）
func (h *AuthHandler) CheckUserAndVerificationCode(ctx context.Context, req *authpb.CheckUserAndVerificationCodeRequest) (*authpb.CheckUserAndVerificationCodeResponse, error) {
	// 检查用户是否已存在
	exists, err := h.userRepo.Check.ByEmail(ctx, req.Email)
	if err != nil {
		logx.S().Errorf("failed to check user exists: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to check user: %v", err)
	}
	if exists {
		return &authpb.CheckUserAndVerificationCodeResponse{
			UserExists:   true,
			CodeValid:    false,
			ErrorMessage: &[]string{"user_already_exists"}[0],
		}, nil
	}

	// 验证验证码
	err = h.userAppSvc.VerifyCode(ctx, userApp.VerifyCodeCmd{
		Email:   req.Email,
		Code:    req.VerificationCode,
		Purpose: "register",
	})
	if err != nil {
		return &authpb.CheckUserAndVerificationCodeResponse{
			UserExists:   false,
			CodeValid:    false,
			ErrorMessage: &[]string{"invalid_verification_code"}[0],
		}, nil
	}

	return &authpb.CheckUserAndVerificationCodeResponse{
		UserExists: false,
		CodeValid:  true,
	}, nil
}

// CreateUserWithProfile 创建用户和 Profile（用于注册流程）
func (h *AuthHandler) CreateUserWithProfile(ctx context.Context, req *authpb.CreateUserWithProfileRequest) (*authpb.CreateUserWithProfileResponse, error) {
	// 创建用户
	var phonePtr *string
	if req.Phone != nil && *req.Phone != "" {
		phonePtr = req.Phone
	}

	user, err := h.userAppSvc.CreateUser(ctx, userApp.CreateUserCmd{
		Editable: userDomain.UserEditable{
			Username: req.Username,
			Email:    req.Email,
			Phone:    phonePtr,
			Password: req.Password,
		},
		Status: "active", // 注册时用户状态为 active
	})
	if err != nil {
		logx.S().Errorf("failed to create user: %v", err)
		errorMsg := err.Error()
		return &authpb.CreateUserWithProfileResponse{
			ErrorMessage: &errorMsg,
		}, nil
	}

	// 创建 Profile（空 profile，只有 user_id）
	_, err = h.profileAppSvc.CreateProfile(ctx, profileAppCommands.CreateProfileCmd{
		UserID:   user.ID(),
		Editable: profileDomain.ProfileEditable{}, // 空 profile
	})
	if err != nil {
		logx.S().Errorf("failed to create profile: %v", err)
		// 即使 profile 创建失败，用户已创建，返回用户信息
		return &authpb.CreateUserWithProfileResponse{
			UserId:       user.ID().String(),
			Username:     user.Editable().Username,
			Email:        user.Editable().Email,
			ErrorMessage: &[]string{"user created but profile creation failed: " + err.Error()}[0],
		}, nil
	}

	return &authpb.CreateUserWithProfileResponse{
		UserId:   user.ID().String(),
		Username: user.Editable().Username,
		Email:    user.Editable().Email,
	}, nil
}
