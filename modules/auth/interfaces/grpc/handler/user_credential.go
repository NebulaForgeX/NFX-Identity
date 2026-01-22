package handler

import (
	"context"

	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	userCredentialAppCommands "nfxid/modules/auth/application/user_credentials/commands"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// CreateUserCredential 创建用户凭证
func (h *UserCredentialHandler) CreateUserCredential(ctx context.Context, req *usercredentialpb.CreateUserCredentialRequest) (*usercredentialpb.CreateUserCredentialResponse, error) {
	// 解析用户ID（id 直接引用 directory.users.id）
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}

	// 转换 protobuf 凭证类型到 domain 类型
	var credentialType userCredentialDomain.CredentialType
	switch req.CredentialType {
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSWORD:
		credentialType = userCredentialDomain.CredentialTypePassword
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSKEY:
		credentialType = userCredentialDomain.CredentialTypePasskey
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_OAUTH_LINK:
		credentialType = userCredentialDomain.CredentialTypeOauthLink
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_SAML:
		credentialType = userCredentialDomain.CredentialTypeSaml
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_LDAP:
		credentialType = userCredentialDomain.CredentialTypeLdap
	default:
		credentialType = userCredentialDomain.CredentialTypePassword
	}

	// 哈希密码（如果凭证类型是密码）
	var passwordHash *string
	var hashAlg *string
	hashParams := make(map[string]interface{})
	if credentialType == userCredentialDomain.CredentialTypePassword && req.Password != "" {
		// 使用 bcrypt 哈希密码
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			logx.S().Errorf("failed to hash password: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
		}
		hashedStr := string(hashedBytes)
		passwordHash = &hashedStr
		alg := "bcrypt"
		hashAlg = &alg
		hashParams["cost"] = bcrypt.DefaultCost
	}

	// 创建命令
	cmd := userCredentialAppCommands.CreateUserCredentialCmd{
		UserID:             userID,
		CredentialType:     credentialType,
		PasswordHash:       passwordHash,
		HashAlg:            hashAlg,
		HashParams:         hashParams,
		Status:             userCredentialDomain.CredentialStatusActive,
		MustChangePassword: req.MustChangePassword,
	}

	// 调用应用服务创建用户凭证
	userCredentialID, err := h.userCredentialAppSvc.CreateUserCredential(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to create user credential: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create user credential: %v", err)
	}

	// 获取创建的用户凭证
	userCredentialView, err := h.userCredentialAppSvc.GetUserCredential(ctx, userCredentialID)
	if err != nil {
		logx.S().Errorf("failed to get created user credential: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created user credential: %v", err)
	}

	// 转换为 protobuf 响应
	userCredential := mapper.UserCredentialROToProto(&userCredentialView)
	return &usercredentialpb.CreateUserCredentialResponse{UserCredential: userCredential}, nil
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
